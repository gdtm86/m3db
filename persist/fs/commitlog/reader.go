// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package commitlog

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/m3db/m3db/persist/fs/msgpack"
	"github.com/m3db/m3db/persist/schema"
	"github.com/m3db/m3db/ts"
	"github.com/m3db/m3x/ident"
	"github.com/m3db/m3x/pool"
	xtime "github.com/m3db/m3x/time"
)

const decoderInBufChanSize = 1000
const decoderOutBufChanSize = 1000

var (
	emptyLogInfo schema.LogInfo

	errCommitLogReaderChunkSizeChecksumMismatch = errors.New("commit log reader encountered chunk size checksum mismatch")
	errCommitLogReaderIsNotReusable             = errors.New("commit log reader is not reusable")
	errCommitLogReaderMultipleReadloops         = errors.New("commit log reader tried to open multiple readLoops, do not call Read() concurrently")
	errCommitLogReaderMissingMetadata           = errors.New("commit log reader encountered a datapoint without corresponding metadata")
)

// ReadAllSeriesPredicate can be passed as the seriesPredicate for callers
// that want a convenient way to read all series in the commitlogs
func ReadAllSeriesPredicate() ReadSeriesPredicate {
	return func(id ident.ID, namespace ident.ID) bool { return true }
}

type seriesMetadata struct {
	Series
	passedPredicate bool
}

type commitLogReader interface {
	// Open opens the commit log for reading
	Open(filePath string) (time.Time, time.Duration, int, error)

	// Read returns the next id and data pair or error, will return io.EOF at end of volume
	Read() (Series, ts.Datapoint, xtime.Unit, ts.Annotation, uint64, error)

	// Close the reader
	Close() error
}

type readResponse struct {
	series      Series
	datapoint   ts.Datapoint
	unit        xtime.Unit
	annotation  ts.Annotation
	uniqueIndex uint64
	resultErr   error
}

type decoderArg struct {
	bytes                []byte
	err                  error
	decodeRemainingToken msgpack.DecodeLogEntryRemainingToken
	uniqueIndex          uint64
	offset               int
	bufPool              chan []byte
}

type readerMetadata struct {
	sync.RWMutex
	numBlockedOrFinishedDecoders int64
}

type reader struct {
	opts                 Options
	numConc              int64
	checkedBytesPool     pool.CheckedBytesPool
	chunkReader          *chunkReader
	infoDecoder          *msgpack.Decoder
	infoDecoderStream    msgpack.DecoderStream
	decoderQueues        []chan decoderArg
	decoderBufPools      []chan []byte
	outChan              chan readResponse
	cancelCtx            context.Context
	cancelFunc           context.CancelFunc
	shutdownCh           chan error
	metadata             readerMetadata
	nextIndex            int64
	hasBeenOpened        bool
	bgWorkersInitialized int64
	seriesPredicate      ReadSeriesPredicate
}

func newCommitLogReader(opts Options, seriesPredicate ReadSeriesPredicate) commitLogReader {
	decodingOpts := opts.FilesystemOptions().DecodingOptions()
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	numConc := opts.ReadConcurrency()
	decoderQueues := make([]chan decoderArg, 0, numConc)
	decoderBufs := make([]chan []byte, 0, numConc)
	for i := 0; i < numConc; i++ {
		decoderQueues = append(decoderQueues, make(chan decoderArg, decoderInBufChanSize))

		chanBufs := make(chan []byte, decoderInBufChanSize+1)
		for i := 0; i < decoderInBufChanSize+1; i++ {
			chanBufs <- make([]byte, opts.FlushSize())
		}
		decoderBufs = append(decoderBufs, chanBufs)
	}
	outBuf := make(chan readResponse, decoderOutBufChanSize*numConc)
	reader := &reader{
		opts:              opts,
		numConc:           int64(numConc),
		checkedBytesPool:  opts.BytesPool(),
		chunkReader:       newChunkReader(opts.FlushSize()),
		infoDecoder:       msgpack.NewDecoder(decodingOpts),
		infoDecoderStream: msgpack.NewDecoderStream(nil),
		decoderQueues:     decoderQueues,
		decoderBufPools:   decoderBufs,
		outChan:           outBuf,
		cancelCtx:         cancelCtx,
		cancelFunc:        cancelFunc,
		shutdownCh:        make(chan error),
		metadata:          readerMetadata{},
		nextIndex:         0,
		seriesPredicate:   seriesPredicate,
	}
	return reader
}

func (r *reader) Open(filePath string) (time.Time, time.Duration, int, error) {
	// Commitlog reader does not currently support being reused
	if r.hasBeenOpened {
		return timeZero, 0, 0, errCommitLogReaderIsNotReusable
	}
	r.hasBeenOpened = true

	fd, err := os.Open(filePath)
	if err != nil {
		return timeZero, 0, 0, err
	}

	r.chunkReader.reset(fd)
	info, err := r.readInfo()
	if err != nil {
		r.Close()
		return timeZero, 0, 0, err
	}
	start := time.Unix(0, info.Start)
	duration := time.Duration(info.Duration)
	index := int(info.Index)

	return start, duration, index, nil
}

// Read guarantees that the datapoints it returns will be in the same order as they are on disk
// for a given series, but they will not be in the same order they are on disk across series.
// I.E, if the commit log looked like this (letters are series and numbers are writes):
// A1, B1, B2, A2, C1, D1, D2, A3, B3, D2
// Then the caller is guaranteed to receive A1 before A2 and A2 before A3, and they are guaranteed
// to see B1 before B2, but they may see B1 before A1 and D2 before B3.
func (r *reader) Read() (
	series Series,
	datapoint ts.Datapoint,
	unit xtime.Unit,
	annotation ts.Annotation,
	uniqueIndex uint64,
	resultErr error,
) {
	if r.nextIndex == 0 {
		err := r.startBackgroundWorkers()
		if err != nil {
			return Series{}, ts.Datapoint{}, xtime.Unit(0), ts.Annotation(nil), 0, err
		}
	}
	rr, ok := <-r.outChan
	if !ok {
		return Series{}, ts.Datapoint{}, xtime.Unit(0), ts.Annotation(nil), 0, io.EOF
	}
	r.nextIndex++
	return rr.series, rr.datapoint, rr.unit, rr.annotation, rr.uniqueIndex, rr.resultErr
}

func (r *reader) startBackgroundWorkers() error {
	// Make sure background workers are never setup more than once
	set := atomic.CompareAndSwapInt64(&r.bgWorkersInitialized, 0, 1)
	if !set {
		return errCommitLogReaderMultipleReadloops
	}

	// Start background worker goroutines
	go r.readLoop()
	for _, decoderQueue := range r.decoderQueues {
		localDecoderQueue := decoderQueue
		go r.decoderLoop(localDecoderQueue, r.outChan)
	}

	return nil
}

func (r *reader) readLoop() {
	defer func() {
		for _, decoderQueue := range r.decoderQueues {
			close(decoderQueue)
		}
	}()

	decodingOpts := r.opts.FilesystemOptions().DecodingOptions()
	decoder := msgpack.NewDecoder(decodingOpts)
	decoderStream := msgpack.NewDecoderStream(nil)

	reusedBytes := make([]byte, 0, r.opts.FlushSize())

	for {
		select {
		case <-r.cancelCtx.Done():
			return
		default:
			data, err := r.readChunk(reusedBytes)
			if err != nil {
				if err == io.EOF {
					return
				}
				r.decoderQueues[0] <- decoderArg{
					bytes: data,
					err:   err,
				}
				continue
			}

			decoderStream.Reset(data)
			decoder.Reset(decoderStream)
			decodeRemainingToken, uniqueIndex, err := decoder.DecodeLogEntryUniqueIndex()

			fmt.Println("pulling")
			bufPool := r.decoderBufPools[uniqueIndex%uint64(r.numConc)]
			buf := <-bufPool
			fmt.Println("done")
			bufCap := cap(buf)
			dataLen := len(data)
			if bufCap < dataLen {
				diff := dataLen - bufCap
				for i := 0; i < diff; i++ {
					buf = append(buf, 0)
				}
			}
			buf = buf[:dataLen]
			copy(buf, data)

			// Distribute work by the uniqueIndex so that each decoder loop is receiving
			// all datapoints for a given series within relative order.
			r.decoderQueues[uniqueIndex%uint64(r.numConc)] <- decoderArg{
				bytes:                buf,
				err:                  err,
				decodeRemainingToken: decodeRemainingToken,
				uniqueIndex:          uniqueIndex,
				offset:               decoderStream.Offset(),
				bufPool:              bufPool,
			}
		}
	}
}

func (r *reader) decoderLoop(inBuf <-chan decoderArg, outBuf chan<- readResponse) {
	var (
		decodingOpts          = r.opts.FilesystemOptions().DecodingOptions()
		decoder               = msgpack.NewDecoder(decodingOpts)
		decoderStream         = msgpack.NewDecoderStream(nil)
		metadataDecoder       = msgpack.NewDecoder(decodingOpts)
		metadataDecoderStream = msgpack.NewDecoderStream(nil)
		metadataLookup        = make(map[uint64]seriesMetadata)
	)

	for arg := range inBuf {
		readResponse := readResponse{}
		// If there is a pre-existing error, just pipe it through
		if arg.err != nil {
			readResponse.resultErr = arg.err
			arg.bufPool <- arg.bytes
			outBuf <- readResponse
			continue
		}

		// Decode the log entry
		decoderStream.Reset(arg.bytes[arg.offset:])
		decoder.Reset(decoderStream)
		entry, err := decoder.DecodeLogEntryRemaining(arg.decodeRemainingToken, arg.uniqueIndex)
		if err != nil {
			readResponse.resultErr = err
			arg.bufPool <- arg.bytes
			outBuf <- readResponse
			continue
		}

		// If the log entry has associated metadata, decode that as well
		if len(entry.Metadata) != 0 {
			err := r.decodeAndHandleMetadata(metadataLookup, metadataDecoder, metadataDecoderStream, entry)
			if err != nil {
				readResponse.resultErr = err
				arg.bufPool <- arg.bytes
				outBuf <- readResponse
				continue
			}
		}

		metadata, hasMetadata := metadataLookup[entry.Index]
		if !hasMetadata {
			// Corrupt commit log
			readResponse.resultErr = errCommitLogReaderMissingMetadata
			arg.bufPool <- arg.bytes
			outBuf <- readResponse
			continue
		}

		if !metadata.passedPredicate {
			continue
		}

		readResponse.series = metadata.Series

		readResponse.datapoint = ts.Datapoint{
			Timestamp: time.Unix(0, entry.Timestamp),
			Value:     entry.Value,
		}
		readResponse.unit = xtime.Unit(byte(entry.Unit))
		readResponse.uniqueIndex = entry.Index
		// Copy annotation to prevent reference to pooled byte slice
		if len(entry.Annotation) > 0 {
			readResponse.annotation = append([]byte(nil), entry.Annotation...)
		}
		arg.bufPool <- arg.bytes
		outBuf <- readResponse
	}

	r.metadata.Lock()
	r.metadata.numBlockedOrFinishedDecoders++
	// If all of the decoders are either finished or blocked then we need to free
	// any pending waiters. This also guarantees that the last decoderLoop to
	// finish will free up any pending waiters (and by then any still-pending
	// metadata is definitely missing from the commitlog)
	if r.metadata.numBlockedOrFinishedDecoders >= r.numConc {
		close(outBuf)
	}
	r.metadata.Unlock()
}

func (r *reader) decodeAndHandleMetadata(
	metadataLookup map[uint64]seriesMetadata,
	metadataDecoder *msgpack.Decoder,
	metadataDecoderStream msgpack.DecoderStream,
	entry schema.LogEntry,
) error {
	metadataDecoderStream.Reset(entry.Metadata)
	metadataDecoder.Reset(metadataDecoderStream)
	decoded, err := metadataDecoder.DecodeLogMetadata()
	if err != nil {
		return err
	}

	id := r.checkedBytesPool.Get(len(decoded.ID))
	id.IncRef()
	id.AppendAll(decoded.ID)

	namespace := r.checkedBytesPool.Get(len(decoded.Namespace))
	namespace.IncRef()
	namespace.AppendAll(decoded.Namespace)

	_, ok := metadataLookup[entry.Index]
	// If the metadata already exists, we can skip this step
	if ok {
		id.DecRef()
		id.Finalize()
		namespace.DecRef()
		namespace.Finalize()
	} else {
		metadata := Series{
			UniqueIndex: entry.Index,
			ID:          ident.BinaryID(id),
			Namespace:   ident.BinaryID(namespace),
			Shard:       decoded.Shard,
		}
		metadataLookup[entry.Index] = seriesMetadata{
			Series:          metadata,
			passedPredicate: r.seriesPredicate(metadata.ID, metadata.Namespace),
		}

		namespace.DecRef()
		id.DecRef()
	}
	return nil
}

func (r *reader) readChunk(buf []byte) ([]byte, error) {
	// Read size of message
	size, err := binary.ReadUvarint(r.chunkReader)
	if err != nil {
		return nil, err
	}

	// Extend buffer as necessary
	if cap(buf) < int(size) {
		diff := int(size) - len(buf)
		for i := 0; i < diff; i++ {
			buf = append(buf, 0)
		}
	}

	// Size target buffer for reading
	buf = buf[:size]

	// Read message
	if _, err := r.chunkReader.Read(buf); err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *reader) readInfo() (schema.LogInfo, error) {
	data, err := r.readChunk([]byte{})
	if err != nil {
		return emptyLogInfo, err
	}
	r.infoDecoderStream.Reset(data)
	r.infoDecoder.Reset(r.infoDecoderStream)
	logInfo, err := r.infoDecoder.DecodeLogInfo()
	return logInfo, err
}

func (r *reader) Close() error {
	// Background goroutines were never started, safe to close immediately.
	if r.nextIndex == 0 {
		return r.close()
	}

	// Shutdown the readLoop goroutine which will shut down the decoderLoops
	// and close the fd
	r.cancelFunc()
	// Drain any unread data from the outBuffers to free any decoderLoops curently
	// in a blocking write
	for {
		_, ok := <-r.outChan
		r.nextIndex++
		if !ok {
			break
		}
	}

	return r.close()
}

func (r *reader) close() error {
	if r.chunkReader.fd == nil {
		return nil
	}
	return r.chunkReader.fd.Close()
}

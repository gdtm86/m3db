// Copyright (c) 2017 Uber Technologies, Inc.
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db/runtime/types.go

package runtime

import (
	gomock "github.com/golang/mock/gomock"
	ratelimit "github.com/m3db/m3db/ratelimit"
	close "github.com/m3db/m3x/close"
	time "time"
)

// Mock of Options interface
type MockOptions struct {
	ctrl     *gomock.Controller
	recorder *_MockOptionsRecorder
}

// Recorder for MockOptions (not exported)
type _MockOptionsRecorder struct {
	mock *MockOptions
}

func NewMockOptions(ctrl *gomock.Controller) *MockOptions {
	mock := &MockOptions{ctrl: ctrl}
	mock.recorder = &_MockOptionsRecorder{mock}
	return mock
}

func (_m *MockOptions) EXPECT() *_MockOptionsRecorder {
	return _m.recorder
}

func (_m *MockOptions) SetPersistRateLimitOptions(value ratelimit.Options) Options {
	ret := _m.ctrl.Call(_m, "SetPersistRateLimitOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetPersistRateLimitOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetPersistRateLimitOptions", arg0)
}

func (_m *MockOptions) PersistRateLimitOptions() ratelimit.Options {
	ret := _m.ctrl.Call(_m, "PersistRateLimitOptions")
	ret0, _ := ret[0].(ratelimit.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) PersistRateLimitOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PersistRateLimitOptions")
}

func (_m *MockOptions) SetWriteNewSeriesAsync(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetWriteNewSeriesAsync", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetWriteNewSeriesAsync(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetWriteNewSeriesAsync", arg0)
}

func (_m *MockOptions) WriteNewSeriesAsync() bool {
	ret := _m.ctrl.Call(_m, "WriteNewSeriesAsync")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockOptionsRecorder) WriteNewSeriesAsync() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteNewSeriesAsync")
}

func (_m *MockOptions) SetWriteNewSeriesBackoffDuration(value time.Duration) Options {
	ret := _m.ctrl.Call(_m, "SetWriteNewSeriesBackoffDuration", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetWriteNewSeriesBackoffDuration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetWriteNewSeriesBackoffDuration", arg0)
}

func (_m *MockOptions) WriteNewSeriesBackoffDuration() time.Duration {
	ret := _m.ctrl.Call(_m, "WriteNewSeriesBackoffDuration")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockOptionsRecorder) WriteNewSeriesBackoffDuration() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteNewSeriesBackoffDuration")
}

func (_m *MockOptions) SetWriteNewSeriesLimitPerShardPerSecond(value int) Options {
	ret := _m.ctrl.Call(_m, "SetWriteNewSeriesLimitPerShardPerSecond", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetWriteNewSeriesLimitPerShardPerSecond(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetWriteNewSeriesLimitPerShardPerSecond", arg0)
}

func (_m *MockOptions) WriteNewSeriesLimitPerShardPerSecond() int {
	ret := _m.ctrl.Call(_m, "WriteNewSeriesLimitPerShardPerSecond")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockOptionsRecorder) WriteNewSeriesLimitPerShardPerSecond() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteNewSeriesLimitPerShardPerSecond")
}

// Mock of OptionsManager interface
type MockOptionsManager struct {
	ctrl     *gomock.Controller
	recorder *_MockOptionsManagerRecorder
}

// Recorder for MockOptionsManager (not exported)
type _MockOptionsManagerRecorder struct {
	mock *MockOptionsManager
}

func NewMockOptionsManager(ctrl *gomock.Controller) *MockOptionsManager {
	mock := &MockOptionsManager{ctrl: ctrl}
	mock.recorder = &_MockOptionsManagerRecorder{mock}
	return mock
}

func (_m *MockOptionsManager) EXPECT() *_MockOptionsManagerRecorder {
	return _m.recorder
}

func (_m *MockOptionsManager) Update(value Options) {
	_m.ctrl.Call(_m, "Update", value)
}

func (_mr *_MockOptionsManagerRecorder) Update(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0)
}

func (_m *MockOptionsManager) Get() Options {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsManagerRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockOptionsManager) RegisterListener(l OptionsListener) close.SimpleCloser {
	ret := _m.ctrl.Call(_m, "RegisterListener", l)
	ret0, _ := ret[0].(close.SimpleCloser)
	return ret0
}

func (_mr *_MockOptionsManagerRecorder) RegisterListener(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterListener", arg0)
}

func (_m *MockOptionsManager) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockOptionsManagerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of OptionsListener interface
type MockOptionsListener struct {
	ctrl     *gomock.Controller
	recorder *_MockOptionsListenerRecorder
}

// Recorder for MockOptionsListener (not exported)
type _MockOptionsListenerRecorder struct {
	mock *MockOptionsListener
}

func NewMockOptionsListener(ctrl *gomock.Controller) *MockOptionsListener {
	mock := &MockOptionsListener{ctrl: ctrl}
	mock.recorder = &_MockOptionsListenerRecorder{mock}
	return mock
}

func (_m *MockOptionsListener) EXPECT() *_MockOptionsListenerRecorder {
	return _m.recorder
}

func (_m *MockOptionsListener) SetRuntimeOptions(value Options) {
	_m.ctrl.Call(_m, "SetRuntimeOptions", value)
}

func (_mr *_MockOptionsListenerRecorder) SetRuntimeOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRuntimeOptions", arg0)
}

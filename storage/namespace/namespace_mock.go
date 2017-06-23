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
// Source: github.com/m3db/m3db/storage/namespace/types.go

package namespace

import (
	gomock "github.com/golang/mock/gomock"
	retention "github.com/m3db/m3db/retention"
	ts "github.com/m3db/m3db/ts"
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

func (_m *MockOptions) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOptionsRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Validate")
}

func (_m *MockOptions) Equal(value Options) bool {
	ret := _m.ctrl.Call(_m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockOptionsRecorder) Equal(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Equal", arg0)
}

func (_m *MockOptions) SetNeedsBootstrap(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetNeedsBootstrap", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetNeedsBootstrap(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetNeedsBootstrap", arg0)
}

func (_m *MockOptions) NeedsBootstrap() bool {
	ret := _m.ctrl.Call(_m, "NeedsBootstrap")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockOptionsRecorder) NeedsBootstrap() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NeedsBootstrap")
}

func (_m *MockOptions) SetNeedsFlush(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetNeedsFlush", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetNeedsFlush(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetNeedsFlush", arg0)
}

func (_m *MockOptions) NeedsFlush() bool {
	ret := _m.ctrl.Call(_m, "NeedsFlush")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockOptionsRecorder) NeedsFlush() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NeedsFlush")
}

func (_m *MockOptions) SetWritesToCommitLog(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetWritesToCommitLog", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetWritesToCommitLog(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetWritesToCommitLog", arg0)
}

func (_m *MockOptions) WritesToCommitLog() bool {
	ret := _m.ctrl.Call(_m, "WritesToCommitLog")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockOptionsRecorder) WritesToCommitLog() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WritesToCommitLog")
}

func (_m *MockOptions) SetNeedsFilesetCleanup(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetNeedsFilesetCleanup", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetNeedsFilesetCleanup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetNeedsFilesetCleanup", arg0)
}

func (_m *MockOptions) NeedsFilesetCleanup() bool {
	ret := _m.ctrl.Call(_m, "NeedsFilesetCleanup")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockOptionsRecorder) NeedsFilesetCleanup() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NeedsFilesetCleanup")
}

func (_m *MockOptions) SetNeedsRepair(value bool) Options {
	ret := _m.ctrl.Call(_m, "SetNeedsRepair", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetNeedsRepair(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetNeedsRepair", arg0)
}

func (_m *MockOptions) NeedsRepair() bool {
	ret := _m.ctrl.Call(_m, "NeedsRepair")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockOptionsRecorder) NeedsRepair() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NeedsRepair")
}

func (_m *MockOptions) SetRetentionOptions(value retention.Options) Options {
	ret := _m.ctrl.Call(_m, "SetRetentionOptions", value)
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) SetRetentionOptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRetentionOptions", arg0)
}

func (_m *MockOptions) RetentionOptions() retention.Options {
	ret := _m.ctrl.Call(_m, "RetentionOptions")
	ret0, _ := ret[0].(retention.Options)
	return ret0
}

func (_mr *_MockOptionsRecorder) RetentionOptions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RetentionOptions")
}

// Mock of Metadata interface
type MockMetadata struct {
	ctrl     *gomock.Controller
	recorder *_MockMetadataRecorder
}

// Recorder for MockMetadata (not exported)
type _MockMetadataRecorder struct {
	mock *MockMetadata
}

func NewMockMetadata(ctrl *gomock.Controller) *MockMetadata {
	mock := &MockMetadata{ctrl: ctrl}
	mock.recorder = &_MockMetadataRecorder{mock}
	return mock
}

func (_m *MockMetadata) EXPECT() *_MockMetadataRecorder {
	return _m.recorder
}

func (_m *MockMetadata) Equal(value Metadata) bool {
	ret := _m.ctrl.Call(_m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockMetadataRecorder) Equal(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Equal", arg0)
}

func (_m *MockMetadata) ID() ts.ID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(ts.ID)
	return ret0
}

func (_mr *_MockMetadataRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockMetadata) Options() Options {
	ret := _m.ctrl.Call(_m, "Options")
	ret0, _ := ret[0].(Options)
	return ret0
}

func (_mr *_MockMetadataRecorder) Options() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Options")
}

// Mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *_MockRegistryRecorder
}

// Recorder for MockRegistry (not exported)
type _MockRegistryRecorder struct {
	mock *MockRegistry
}

func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &_MockRegistryRecorder{mock}
	return mock
}

func (_m *MockRegistry) EXPECT() *_MockRegistryRecorder {
	return _m.recorder
}

func (_m *MockRegistry) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRegistryRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Validate")
}

func (_m *MockRegistry) Equal(value Registry) bool {
	ret := _m.ctrl.Call(_m, "Equal", value)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockRegistryRecorder) Equal(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Equal", arg0)
}

func (_m *MockRegistry) Get(_param0 ts.ID) (Metadata, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRegistryRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockRegistry) IDs() []ts.ID {
	ret := _m.ctrl.Call(_m, "IDs")
	ret0, _ := ret[0].([]ts.ID)
	return ret0
}

func (_mr *_MockRegistryRecorder) IDs() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IDs")
}

func (_m *MockRegistry) Metadatas() []Metadata {
	ret := _m.ctrl.Call(_m, "Metadatas")
	ret0, _ := ret[0].([]Metadata)
	return ret0
}

func (_mr *_MockRegistryRecorder) Metadatas() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Metadatas")
}

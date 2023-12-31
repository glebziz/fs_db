// Code generated by MockGen. DO NOT EDIT.
// Source: reader.go
//
// Generated by this command:
//
//	mockgen -source reader.go -destination mocks/mocks.go -typed true
//
// Package mock_streamreader is a generated GoMock package.
package mock_streamreader

import (
	streamreader "github.com/glebziz/fs_db/internal/utils/grpc/streamreader"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRequest is a mock of Request interface.
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest.
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance.
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// GetChunk mocks base method.
func (m *MockRequest) GetChunk() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChunk")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetChunk indicates an expected call of GetChunk.
func (mr *MockRequestMockRecorder) GetChunk() *RequestGetChunkCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChunk", reflect.TypeOf((*MockRequest)(nil).GetChunk))
	return &RequestGetChunkCall{Call: call}
}

// RequestGetChunkCall wrap *gomock.Call
type RequestGetChunkCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RequestGetChunkCall) Return(arg0 []byte) *RequestGetChunkCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RequestGetChunkCall) Do(f func() []byte) *RequestGetChunkCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RequestGetChunkCall) DoAndReturn(f func() []byte) *RequestGetChunkCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockStream is a mock of Stream interface.
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream.
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance.
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// Recv mocks base method.
func (m *MockStream) Recv() (streamreader.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(streamreader.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockStreamMockRecorder) Recv() *StreamRecvCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockStream)(nil).Recv))
	return &StreamRecvCall{Call: call}
}

// StreamRecvCall wrap *gomock.Call
type StreamRecvCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StreamRecvCall) Return(arg0 streamreader.Request, arg1 error) *StreamRecvCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StreamRecvCall) Do(f func() (streamreader.Request, error)) *StreamRecvCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StreamRecvCall) DoAndReturn(f func() (streamreader.Request, error)) *StreamRecvCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/pkg/executor/stream (interfaces: LogsStream)

// Package stream is a generated GoMock package.
package stream

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogsStream is a mock of LogsStream interface.
type MockLogsStream struct {
	ctrl     *gomock.Controller
	recorder *MockLogsStreamMockRecorder
}

// MockLogsStreamMockRecorder is the mock recorder for MockLogsStream.
type MockLogsStreamMockRecorder struct {
	mock *MockLogsStream
}

// NewMockLogsStream creates a new mock instance.
func NewMockLogsStream(ctrl *gomock.Controller) *MockLogsStream {
	mock := &MockLogsStream{ctrl: ctrl}
	mock.recorder = &MockLogsStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogsStream) EXPECT() *MockLogsStreamMockRecorder {
	return m.recorder
}

// GetRange mocks base method.
func (m *MockLogsStream) GetRange(arg0 context.Context, arg1 string, arg2, arg3 int) (chan []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(chan []byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRange indicates an expected call of GetRange.
func (mr *MockLogsStreamMockRecorder) GetRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRange", reflect.TypeOf((*MockLogsStream)(nil).GetRange), arg0, arg1, arg2, arg3)
}

// Init mocks base method.
func (m *MockLogsStream) Init(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockLogsStreamMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockLogsStream)(nil).Init), arg0)
}

// Listen mocks base method.
func (m *MockLogsStream) Listen(arg0 context.Context, arg1 string) (chan []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listen", arg0, arg1)
	ret0, _ := ret[0].(chan []byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Listen indicates an expected call of Listen.
func (mr *MockLogsStreamMockRecorder) Listen(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockLogsStream)(nil).Listen), arg0, arg1)
}

// Publish mocks base method.
func (m *MockLogsStream) Publish(arg0 context.Context, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockLogsStreamMockRecorder) Publish(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockLogsStream)(nil).Publish), arg0, arg1, arg2)
}

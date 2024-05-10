// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/pkg/triggers (interfaces: LeaseBackend)

// Package triggers is a generated GoMock package.
package triggers

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLeaseBackend is a mock of LeaseBackend interface.
type MockLeaseBackend struct {
	ctrl     *gomock.Controller
	recorder *MockLeaseBackendMockRecorder
}

// MockLeaseBackendMockRecorder is the mock recorder for MockLeaseBackend.
type MockLeaseBackendMockRecorder struct {
	mock *MockLeaseBackend
}

// NewMockLeaseBackend creates a new mock instance.
func NewMockLeaseBackend(ctrl *gomock.Controller) *MockLeaseBackend {
	mock := &MockLeaseBackend{ctrl: ctrl}
	mock.recorder = &MockLeaseBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaseBackend) EXPECT() *MockLeaseBackendMockRecorder {
	return m.recorder
}

// TryAcquire mocks base method.
func (m *MockLeaseBackend) TryAcquire(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryAcquire", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TryAcquire indicates an expected call of TryAcquire.
func (mr *MockLeaseBackendMockRecorder) TryAcquire(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryAcquire", reflect.TypeOf((*MockLeaseBackend)(nil).TryAcquire), arg0, arg1, arg2)
}
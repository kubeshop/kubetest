// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/cmd/api-server/commons (interfaces: DeprecatedClients)

// Package commons is a generated GoMock package.
package commons

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	executors "github.com/kubeshop/testkube-operator/pkg/client/executors/v1"
	templates "github.com/kubeshop/testkube-operator/pkg/client/templates/v1"
	testexecutions "github.com/kubeshop/testkube-operator/pkg/client/testexecutions/v1"
	tests "github.com/kubeshop/testkube-operator/pkg/client/tests/v3"
	testsources "github.com/kubeshop/testkube-operator/pkg/client/testsources/v1"
	testsuiteexecutions "github.com/kubeshop/testkube-operator/pkg/client/testsuiteexecutions/v1"
	v3 "github.com/kubeshop/testkube-operator/pkg/client/testsuites/v3"
)

// MockDeprecatedClients is a mock of DeprecatedClients interface.
type MockDeprecatedClients struct {
	ctrl     *gomock.Controller
	recorder *MockDeprecatedClientsMockRecorder
}

// MockDeprecatedClientsMockRecorder is the mock recorder for MockDeprecatedClients.
type MockDeprecatedClientsMockRecorder struct {
	mock *MockDeprecatedClients
}

// NewMockDeprecatedClients creates a new mock instance.
func NewMockDeprecatedClients(ctrl *gomock.Controller) *MockDeprecatedClients {
	mock := &MockDeprecatedClients{ctrl: ctrl}
	mock.recorder = &MockDeprecatedClientsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeprecatedClients) EXPECT() *MockDeprecatedClientsMockRecorder {
	return m.recorder
}

// Executors mocks base method.
func (m *MockDeprecatedClients) Executors() executors.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Executors")
	ret0, _ := ret[0].(executors.Interface)
	return ret0
}

// Executors indicates an expected call of Executors.
func (mr *MockDeprecatedClientsMockRecorder) Executors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Executors", reflect.TypeOf((*MockDeprecatedClients)(nil).Executors))
}

// Templates mocks base method.
func (m *MockDeprecatedClients) Templates() templates.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Templates")
	ret0, _ := ret[0].(templates.Interface)
	return ret0
}

// Templates indicates an expected call of Templates.
func (mr *MockDeprecatedClientsMockRecorder) Templates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Templates", reflect.TypeOf((*MockDeprecatedClients)(nil).Templates))
}

// TestExecutions mocks base method.
func (m *MockDeprecatedClients) TestExecutions() testexecutions.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestExecutions")
	ret0, _ := ret[0].(testexecutions.Interface)
	return ret0
}

// TestExecutions indicates an expected call of TestExecutions.
func (mr *MockDeprecatedClientsMockRecorder) TestExecutions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestExecutions", reflect.TypeOf((*MockDeprecatedClients)(nil).TestExecutions))
}

// TestSources mocks base method.
func (m *MockDeprecatedClients) TestSources() testsources.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestSources")
	ret0, _ := ret[0].(testsources.Interface)
	return ret0
}

// TestSources indicates an expected call of TestSources.
func (mr *MockDeprecatedClientsMockRecorder) TestSources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestSources", reflect.TypeOf((*MockDeprecatedClients)(nil).TestSources))
}

// TestSuiteExecutions mocks base method.
func (m *MockDeprecatedClients) TestSuiteExecutions() testsuiteexecutions.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestSuiteExecutions")
	ret0, _ := ret[0].(testsuiteexecutions.Interface)
	return ret0
}

// TestSuiteExecutions indicates an expected call of TestSuiteExecutions.
func (mr *MockDeprecatedClientsMockRecorder) TestSuiteExecutions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestSuiteExecutions", reflect.TypeOf((*MockDeprecatedClients)(nil).TestSuiteExecutions))
}

// TestSuites mocks base method.
func (m *MockDeprecatedClients) TestSuites() v3.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestSuites")
	ret0, _ := ret[0].(v3.Interface)
	return ret0
}

// TestSuites indicates an expected call of TestSuites.
func (mr *MockDeprecatedClientsMockRecorder) TestSuites() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestSuites", reflect.TypeOf((*MockDeprecatedClients)(nil).TestSuites))
}

// Tests mocks base method.
func (m *MockDeprecatedClients) Tests() tests.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tests")
	ret0, _ := ret[0].(tests.Interface)
	return ret0
}

// Tests indicates an expected call of Tests.
func (mr *MockDeprecatedClientsMockRecorder) Tests() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tests", reflect.TypeOf((*MockDeprecatedClients)(nil).Tests))
}

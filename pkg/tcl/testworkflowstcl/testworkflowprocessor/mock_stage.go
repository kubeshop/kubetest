// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/pkg/tcl/testworkflowstcl/testworkflowprocessor (interfaces: Stage)

// Package testworkflowprocessor is a generated GoMock package.
package testworkflowprocessor

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"
	imageinspector "github.com/kubeshop/testkube/pkg/imageinspector"
	expressionstcl "github.com/kubeshop/testkube/pkg/tcl/expressionstcl"
)

// MockStage is a mock of Stage interface.
type MockStage struct {
	ctrl     *gomock.Controller
	recorder *MockStageMockRecorder
}

// MockStageMockRecorder is the mock recorder for MockStage.
type MockStageMockRecorder struct {
	mock *MockStage
}

// NewMockStage creates a new mock instance.
func NewMockStage(ctrl *gomock.Controller) *MockStage {
	mock := &MockStage{ctrl: ctrl}
	mock.recorder = &MockStageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStage) EXPECT() *MockStageMockRecorder {
	return m.recorder
}

// AppendConditions mocks base method.
func (m *MockStage) AppendConditions(arg0 ...string) StageLifecycle {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AppendConditions", varargs...)
	ret0, _ := ret[0].(StageLifecycle)
	return ret0
}

// AppendConditions indicates an expected call of AppendConditions.
func (mr *MockStageMockRecorder) AppendConditions(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendConditions", reflect.TypeOf((*MockStage)(nil).AppendConditions), arg0...)
}

// ApplyImages mocks base method.
func (m *MockStage) ApplyImages(arg0 map[string]*imageinspector.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyImages", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyImages indicates an expected call of ApplyImages.
func (mr *MockStageMockRecorder) ApplyImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyImages", reflect.TypeOf((*MockStage)(nil).ApplyImages), arg0)
}

// Condition mocks base method.
func (m *MockStage) Condition() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Condition")
	ret0, _ := ret[0].(string)
	return ret0
}

// Condition indicates an expected call of Condition.
func (mr *MockStageMockRecorder) Condition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Condition", reflect.TypeOf((*MockStage)(nil).Condition))
}

// ContainerStages mocks base method.
func (m *MockStage) ContainerStages() []ContainerStage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerStages")
	ret0, _ := ret[0].([]ContainerStage)
	return ret0
}

// ContainerStages indicates an expected call of ContainerStages.
func (mr *MockStageMockRecorder) ContainerStages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStages", reflect.TypeOf((*MockStage)(nil).ContainerStages))
}

// Flatten mocks base method.
func (m *MockStage) Flatten() []Stage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flatten")
	ret0, _ := ret[0].([]Stage)
	return ret0
}

// Flatten indicates an expected call of Flatten.
func (mr *MockStageMockRecorder) Flatten() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flatten", reflect.TypeOf((*MockStage)(nil).Flatten))
}

// GetImages mocks base method.
func (m *MockStage) GetImages() map[string]struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImages")
	ret0, _ := ret[0].(map[string]struct{})
	return ret0
}

// GetImages indicates an expected call of GetImages.
func (mr *MockStageMockRecorder) GetImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImages", reflect.TypeOf((*MockStage)(nil).GetImages))
}

// Len mocks base method.
func (m *MockStage) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockStageMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockStage)(nil).Len))
}

// Name mocks base method.
func (m *MockStage) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockStageMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockStage)(nil).Name))
}

// Negative mocks base method.
func (m *MockStage) Negative() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Negative")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Negative indicates an expected call of Negative.
func (mr *MockStageMockRecorder) Negative() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Negative", reflect.TypeOf((*MockStage)(nil).Negative))
}

// Optional mocks base method.
func (m *MockStage) Optional() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Optional")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Optional indicates an expected call of Optional.
func (mr *MockStageMockRecorder) Optional() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Optional", reflect.TypeOf((*MockStage)(nil).Optional))
}

// Ref mocks base method.
func (m *MockStage) Ref() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ref")
	ret0, _ := ret[0].(string)
	return ret0
}

// Ref indicates an expected call of Ref.
func (mr *MockStageMockRecorder) Ref() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ref", reflect.TypeOf((*MockStage)(nil).Ref))
}

// Resolve mocks base method.
func (m *MockStage) Resolve(arg0 ...expressionstcl.Machine) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Resolve", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resolve indicates an expected call of Resolve.
func (mr *MockStageMockRecorder) Resolve(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockStage)(nil).Resolve), arg0...)
}

// RetryPolicy mocks base method.
func (m *MockStage) RetryPolicy() v1.RetryPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetryPolicy")
	ret0, _ := ret[0].(v1.RetryPolicy)
	return ret0
}

// RetryPolicy indicates an expected call of RetryPolicy.
func (mr *MockStageMockRecorder) RetryPolicy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryPolicy", reflect.TypeOf((*MockStage)(nil).RetryPolicy))
}

// SetCondition mocks base method.
func (m *MockStage) SetCondition(arg0 string) StageLifecycle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCondition", arg0)
	ret0, _ := ret[0].(StageLifecycle)
	return ret0
}

// SetCondition indicates an expected call of SetCondition.
func (mr *MockStageMockRecorder) SetCondition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCondition", reflect.TypeOf((*MockStage)(nil).SetCondition), arg0)
}

// SetName mocks base method.
func (m *MockStage) SetName(arg0 string, arg1 ...string) StageMetadata {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetName", varargs...)
	ret0, _ := ret[0].(StageMetadata)
	return ret0
}

// SetName indicates an expected call of SetName.
func (mr *MockStageMockRecorder) SetName(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockStage)(nil).SetName), varargs...)
}

// SetNegative mocks base method.
func (m *MockStage) SetNegative(arg0 bool) StageLifecycle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNegative", arg0)
	ret0, _ := ret[0].(StageLifecycle)
	return ret0
}

// SetNegative indicates an expected call of SetNegative.
func (mr *MockStageMockRecorder) SetNegative(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNegative", reflect.TypeOf((*MockStage)(nil).SetNegative), arg0)
}

// SetOptional mocks base method.
func (m *MockStage) SetOptional(arg0 bool) StageLifecycle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOptional", arg0)
	ret0, _ := ret[0].(StageLifecycle)
	return ret0
}

// SetOptional indicates an expected call of SetOptional.
func (mr *MockStageMockRecorder) SetOptional(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOptional", reflect.TypeOf((*MockStage)(nil).SetOptional), arg0)
}

// SetRetryPolicy mocks base method.
func (m *MockStage) SetRetryPolicy(arg0 v1.RetryPolicy) StageLifecycle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRetryPolicy", arg0)
	ret0, _ := ret[0].(StageLifecycle)
	return ret0
}

// SetRetryPolicy indicates an expected call of SetRetryPolicy.
func (mr *MockStageMockRecorder) SetRetryPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRetryPolicy", reflect.TypeOf((*MockStage)(nil).SetRetryPolicy), arg0)
}

// SetTimeout mocks base method.
func (m *MockStage) SetTimeout(arg0 string) StageLifecycle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTimeout", arg0)
	ret0, _ := ret[0].(StageLifecycle)
	return ret0
}

// SetTimeout indicates an expected call of SetTimeout.
func (mr *MockStageMockRecorder) SetTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeout", reflect.TypeOf((*MockStage)(nil).SetTimeout), arg0)
}

// Signature mocks base method.
func (m *MockStage) Signature() Signature {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signature")
	ret0, _ := ret[0].(Signature)
	return ret0
}

// Signature indicates an expected call of Signature.
func (mr *MockStageMockRecorder) Signature() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signature", reflect.TypeOf((*MockStage)(nil).Signature))
}

// Timeout mocks base method.
func (m *MockStage) Timeout() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(string)
	return ret0
}

// Timeout indicates an expected call of Timeout.
func (mr *MockStageMockRecorder) Timeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MockStage)(nil).Timeout))
}

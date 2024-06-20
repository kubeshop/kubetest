// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/pkg/expressions (interfaces: StaticValue)

// Package expressions is a generated GoMock package.
package expressions

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStaticValue is a mock of StaticValue interface.
type MockStaticValue struct {
	ctrl     *gomock.Controller
	recorder *MockStaticValueMockRecorder
}

// MockStaticValueMockRecorder is the mock recorder for MockStaticValue.
type MockStaticValueMockRecorder struct {
	mock *MockStaticValue
}

// NewMockStaticValue creates a new mock instance.
func NewMockStaticValue(ctrl *gomock.Controller) *MockStaticValue {
	mock := &MockStaticValue{ctrl: ctrl}
	mock.recorder = &MockStaticValueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStaticValue) EXPECT() *MockStaticValueMockRecorder {
	return m.recorder
}

// Accessors mocks base method.
func (m *MockStaticValue) Accessors() map[string]struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accessors")
	ret0, _ := ret[0].(map[string]struct{})
	return ret0
}

// Accessors indicates an expected call of Accessors.
func (mr *MockStaticValueMockRecorder) Accessors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accessors", reflect.TypeOf((*MockStaticValue)(nil).Accessors))
}

// BoolValue mocks base method.
func (m *MockStaticValue) BoolValue() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BoolValue")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BoolValue indicates an expected call of BoolValue.
func (mr *MockStaticValueMockRecorder) BoolValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BoolValue", reflect.TypeOf((*MockStaticValue)(nil).BoolValue))
}

// FloatValue mocks base method.
func (m *MockStaticValue) FloatValue() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FloatValue")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FloatValue indicates an expected call of FloatValue.
func (mr *MockStaticValueMockRecorder) FloatValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FloatValue", reflect.TypeOf((*MockStaticValue)(nil).FloatValue))
}

// Functions mocks base method.
func (m *MockStaticValue) Functions() map[string]struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Functions")
	ret0, _ := ret[0].(map[string]struct{})
	return ret0
}

// Functions indicates an expected call of Functions.
func (mr *MockStaticValueMockRecorder) Functions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Functions", reflect.TypeOf((*MockStaticValue)(nil).Functions))
}

// IntValue mocks base method.
func (m *MockStaticValue) IntValue() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntValue")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IntValue indicates an expected call of IntValue.
func (mr *MockStaticValueMockRecorder) IntValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntValue", reflect.TypeOf((*MockStaticValue)(nil).IntValue))
}

// IsBool mocks base method.
func (m *MockStaticValue) IsBool() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBool")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBool indicates an expected call of IsBool.
func (mr *MockStaticValueMockRecorder) IsBool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBool", reflect.TypeOf((*MockStaticValue)(nil).IsBool))
}

// IsInt mocks base method.
func (m *MockStaticValue) IsInt() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInt")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInt indicates an expected call of IsInt.
func (mr *MockStaticValueMockRecorder) IsInt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInt", reflect.TypeOf((*MockStaticValue)(nil).IsInt))
}

// IsMap mocks base method.
func (m *MockStaticValue) IsMap() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMap")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMap indicates an expected call of IsMap.
func (mr *MockStaticValueMockRecorder) IsMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMap", reflect.TypeOf((*MockStaticValue)(nil).IsMap))
}

// IsNone mocks base method.
func (m *MockStaticValue) IsNone() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNone")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNone indicates an expected call of IsNone.
func (mr *MockStaticValueMockRecorder) IsNone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNone", reflect.TypeOf((*MockStaticValue)(nil).IsNone))
}

// IsNumber mocks base method.
func (m *MockStaticValue) IsNumber() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNumber")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNumber indicates an expected call of IsNumber.
func (mr *MockStaticValueMockRecorder) IsNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNumber", reflect.TypeOf((*MockStaticValue)(nil).IsNumber))
}

// IsSlice mocks base method.
func (m *MockStaticValue) IsSlice() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSlice")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSlice indicates an expected call of IsSlice.
func (mr *MockStaticValueMockRecorder) IsSlice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSlice", reflect.TypeOf((*MockStaticValue)(nil).IsSlice))
}

// IsString mocks base method.
func (m *MockStaticValue) IsString() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsString")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsString indicates an expected call of IsString.
func (mr *MockStaticValueMockRecorder) IsString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsString", reflect.TypeOf((*MockStaticValue)(nil).IsString))
}

// MapValue mocks base method.
func (m *MockStaticValue) MapValue() (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapValue")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MapValue indicates an expected call of MapValue.
func (mr *MockStaticValueMockRecorder) MapValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapValue", reflect.TypeOf((*MockStaticValue)(nil).MapValue))
}

// Resolve mocks base method.
func (m *MockStaticValue) Resolve(arg0 ...Machine) (Expression, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Resolve", varargs...)
	ret0, _ := ret[0].(Expression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve.
func (mr *MockStaticValueMockRecorder) Resolve(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockStaticValue)(nil).Resolve), arg0...)
}

// SafeResolve mocks base method.
func (m *MockStaticValue) SafeResolve(arg0 ...Machine) (Expression, bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SafeResolve", varargs...)
	ret0, _ := ret[0].(Expression)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SafeResolve indicates an expected call of SafeResolve.
func (mr *MockStaticValueMockRecorder) SafeResolve(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeResolve", reflect.TypeOf((*MockStaticValue)(nil).SafeResolve), arg0...)
}

// SafeString mocks base method.
func (m *MockStaticValue) SafeString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SafeString")
	ret0, _ := ret[0].(string)
	return ret0
}

// SafeString indicates an expected call of SafeString.
func (mr *MockStaticValueMockRecorder) SafeString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeString", reflect.TypeOf((*MockStaticValue)(nil).SafeString))
}

// SliceValue mocks base method.
func (m *MockStaticValue) SliceValue() ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SliceValue")
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SliceValue indicates an expected call of SliceValue.
func (mr *MockStaticValueMockRecorder) SliceValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SliceValue", reflect.TypeOf((*MockStaticValue)(nil).SliceValue))
}

// Static mocks base method.
func (m *MockStaticValue) Static() StaticValue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Static")
	ret0, _ := ret[0].(StaticValue)
	return ret0
}

// Static indicates an expected call of Static.
func (mr *MockStaticValueMockRecorder) Static() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Static", reflect.TypeOf((*MockStaticValue)(nil).Static))
}

// String mocks base method.
func (m *MockStaticValue) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockStaticValueMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockStaticValue)(nil).String))
}

// StringValue mocks base method.
func (m *MockStaticValue) StringValue() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringValue")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StringValue indicates an expected call of StringValue.
func (mr *MockStaticValueMockRecorder) StringValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringValue", reflect.TypeOf((*MockStaticValue)(nil).StringValue))
}

// Template mocks base method.
func (m *MockStaticValue) Template() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template")
	ret0, _ := ret[0].(string)
	return ret0
}

// Template indicates an expected call of Template.
func (mr *MockStaticValueMockRecorder) Template() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockStaticValue)(nil).Template))
}

// Type mocks base method.
func (m *MockStaticValue) Type() Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(Type)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockStaticValueMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockStaticValue)(nil).Type))
}

// Value mocks base method.
func (m *MockStaticValue) Value() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockStaticValueMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockStaticValue)(nil).Value))
}

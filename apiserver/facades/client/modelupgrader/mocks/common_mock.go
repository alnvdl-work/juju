// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common (interfaces: BlockCheckerInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBlockCheckerInterface is a mock of BlockCheckerInterface interface.
type MockBlockCheckerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBlockCheckerInterfaceMockRecorder
}

// MockBlockCheckerInterfaceMockRecorder is the mock recorder for MockBlockCheckerInterface.
type MockBlockCheckerInterfaceMockRecorder struct {
	mock *MockBlockCheckerInterface
}

// NewMockBlockCheckerInterface creates a new mock instance.
func NewMockBlockCheckerInterface(ctrl *gomock.Controller) *MockBlockCheckerInterface {
	mock := &MockBlockCheckerInterface{ctrl: ctrl}
	mock.recorder = &MockBlockCheckerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockCheckerInterface) EXPECT() *MockBlockCheckerInterfaceMockRecorder {
	return m.recorder
}

// ChangeAllowed mocks base method.
func (m *MockBlockCheckerInterface) ChangeAllowed() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeAllowed")
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeAllowed indicates an expected call of ChangeAllowed.
func (mr *MockBlockCheckerInterfaceMockRecorder) ChangeAllowed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAllowed", reflect.TypeOf((*MockBlockCheckerInterface)(nil).ChangeAllowed))
}

// DestroyAllowed mocks base method.
func (m *MockBlockCheckerInterface) DestroyAllowed() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyAllowed")
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyAllowed indicates an expected call of DestroyAllowed.
func (mr *MockBlockCheckerInterfaceMockRecorder) DestroyAllowed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyAllowed", reflect.TypeOf((*MockBlockCheckerInterface)(nil).DestroyAllowed))
}

// RemoveAllowed mocks base method.
func (m *MockBlockCheckerInterface) RemoveAllowed() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllowed")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllowed indicates an expected call of RemoveAllowed.
func (mr *MockBlockCheckerInterfaceMockRecorder) RemoveAllowed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllowed", reflect.TypeOf((*MockBlockCheckerInterface)(nil).RemoveAllowed))
}
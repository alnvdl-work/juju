// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/meterstatus (interfaces: UnitStateAPI,StateReadWriter)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/interface_mock.go github.com/juju/juju/worker/meterstatus UnitStateAPI,StateReadWriter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	params "github.com/juju/juju/rpc/params"
	meterstatus "github.com/juju/juju/worker/meterstatus"
	gomock "go.uber.org/mock/gomock"
)

// MockUnitStateAPI is a mock of UnitStateAPI interface.
type MockUnitStateAPI struct {
	ctrl     *gomock.Controller
	recorder *MockUnitStateAPIMockRecorder
}

// MockUnitStateAPIMockRecorder is the mock recorder for MockUnitStateAPI.
type MockUnitStateAPIMockRecorder struct {
	mock *MockUnitStateAPI
}

// NewMockUnitStateAPI creates a new mock instance.
func NewMockUnitStateAPI(ctrl *gomock.Controller) *MockUnitStateAPI {
	mock := &MockUnitStateAPI{ctrl: ctrl}
	mock.recorder = &MockUnitStateAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitStateAPI) EXPECT() *MockUnitStateAPIMockRecorder {
	return m.recorder
}

// SetState mocks base method.
func (m *MockUnitStateAPI) SetState(arg0 params.SetUnitStateArg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetState indicates an expected call of SetState.
func (mr *MockUnitStateAPIMockRecorder) SetState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockUnitStateAPI)(nil).SetState), arg0)
}

// State mocks base method.
func (m *MockUnitStateAPI) State() (params.UnitStateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(params.UnitStateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State.
func (mr *MockUnitStateAPIMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockUnitStateAPI)(nil).State))
}

// MockStateReadWriter is a mock of StateReadWriter interface.
type MockStateReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockStateReadWriterMockRecorder
}

// MockStateReadWriterMockRecorder is the mock recorder for MockStateReadWriter.
type MockStateReadWriterMockRecorder struct {
	mock *MockStateReadWriter
}

// NewMockStateReadWriter creates a new mock instance.
func NewMockStateReadWriter(ctrl *gomock.Controller) *MockStateReadWriter {
	mock := &MockStateReadWriter{ctrl: ctrl}
	mock.recorder = &MockStateReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateReadWriter) EXPECT() *MockStateReadWriterMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockStateReadWriter) Read() (*meterstatus.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].(*meterstatus.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockStateReadWriterMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStateReadWriter)(nil).Read))
}

// Write mocks base method.
func (m *MockStateReadWriter) Write(arg0 *meterstatus.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockStateReadWriterMockRecorder) Write(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStateReadWriter)(nil).Write), arg0)
}

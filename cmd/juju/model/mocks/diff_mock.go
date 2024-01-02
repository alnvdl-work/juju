// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/model (interfaces: DiffCommandAPI)
//
// Generated by this command:
//
//	mockgen -package mocks -destination ./mocks/diff_mock.go github.com/juju/juju/cmd/juju/model DiffCommandAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	model "github.com/juju/juju/core/model"
	gomock "go.uber.org/mock/gomock"
)

// MockDiffCommandAPI is a mock of DiffCommandAPI interface.
type MockDiffCommandAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDiffCommandAPIMockRecorder
}

// MockDiffCommandAPIMockRecorder is the mock recorder for MockDiffCommandAPI.
type MockDiffCommandAPIMockRecorder struct {
	mock *MockDiffCommandAPI
}

// NewMockDiffCommandAPI creates a new mock instance.
func NewMockDiffCommandAPI(ctrl *gomock.Controller) *MockDiffCommandAPI {
	mock := &MockDiffCommandAPI{ctrl: ctrl}
	mock.recorder = &MockDiffCommandAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiffCommandAPI) EXPECT() *MockDiffCommandAPIMockRecorder {
	return m.recorder
}

// BranchInfo mocks base method.
func (m *MockDiffCommandAPI) BranchInfo(arg0 string, arg1 bool, arg2 func(time.Time) string) (map[string]model.Generation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BranchInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]model.Generation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BranchInfo indicates an expected call of BranchInfo.
func (mr *MockDiffCommandAPIMockRecorder) BranchInfo(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BranchInfo", reflect.TypeOf((*MockDiffCommandAPI)(nil).BranchInfo), arg0, arg1, arg2)
}

// Close mocks base method.
func (m *MockDiffCommandAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDiffCommandAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDiffCommandAPI)(nil).Close))
}

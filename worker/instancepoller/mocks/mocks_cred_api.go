// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/common (interfaces: CredentialAPI)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/mocks_cred_api.go github.com/juju/juju/worker/common CredentialAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCredentialAPI is a mock of CredentialAPI interface.
type MockCredentialAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialAPIMockRecorder
}

// MockCredentialAPIMockRecorder is the mock recorder for MockCredentialAPI.
type MockCredentialAPIMockRecorder struct {
	mock *MockCredentialAPI
}

// NewMockCredentialAPI creates a new mock instance.
func NewMockCredentialAPI(ctrl *gomock.Controller) *MockCredentialAPI {
	mock := &MockCredentialAPI{ctrl: ctrl}
	mock.recorder = &MockCredentialAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialAPI) EXPECT() *MockCredentialAPIMockRecorder {
	return m.recorder
}

// InvalidateModelCredential mocks base method.
func (m *MockCredentialAPI) InvalidateModelCredential(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateModelCredential", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateModelCredential indicates an expected call of InvalidateModelCredential.
func (mr *MockCredentialAPIMockRecorder) InvalidateModelCredential(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateModelCredential", reflect.TypeOf((*MockCredentialAPI)(nil).InvalidateModelCredential), arg0)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: SecretsTriggerWatcher)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	watcher "github.com/juju/juju/core/watcher"
)

// MockSecretsTriggerWatcher is a mock of SecretsTriggerWatcher interface.
type MockSecretsTriggerWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsTriggerWatcherMockRecorder
}

// MockSecretsTriggerWatcherMockRecorder is the mock recorder for MockSecretsTriggerWatcher.
type MockSecretsTriggerWatcherMockRecorder struct {
	mock *MockSecretsTriggerWatcher
}

// NewMockSecretsTriggerWatcher creates a new mock instance.
func NewMockSecretsTriggerWatcher(ctrl *gomock.Controller) *MockSecretsTriggerWatcher {
	mock := &MockSecretsTriggerWatcher{ctrl: ctrl}
	mock.recorder = &MockSecretsTriggerWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsTriggerWatcher) EXPECT() *MockSecretsTriggerWatcherMockRecorder {
	return m.recorder
}

// Changes mocks base method.
func (m *MockSecretsTriggerWatcher) Changes() watcher.SecretTriggerChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(watcher.SecretTriggerChannel)
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockSecretsTriggerWatcherMockRecorder) Changes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockSecretsTriggerWatcher)(nil).Changes))
}

// Err mocks base method.
func (m *MockSecretsTriggerWatcher) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockSecretsTriggerWatcherMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSecretsTriggerWatcher)(nil).Err))
}

// Kill mocks base method.
func (m *MockSecretsTriggerWatcher) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockSecretsTriggerWatcherMockRecorder) Kill() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockSecretsTriggerWatcher)(nil).Kill))
}

// Stop mocks base method.
func (m *MockSecretsTriggerWatcher) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockSecretsTriggerWatcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockSecretsTriggerWatcher)(nil).Stop))
}

// Wait mocks base method.
func (m *MockSecretsTriggerWatcher) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockSecretsTriggerWatcherMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockSecretsTriggerWatcher)(nil).Wait))
}

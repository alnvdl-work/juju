// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/store (interfaces: CharmAdder,CharmrepoForDeploy,CharmsAPI,DownloadBundleClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	url "net/url"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	charm "github.com/juju/charm/v9"
	params "github.com/juju/charmrepo/v7/csclient/params"
	charms "github.com/juju/juju/api/client/charms"
	charm0 "github.com/juju/juju/api/common/charm"
	charmhub "github.com/juju/juju/charmhub"
)

// MockCharmAdder is a mock of CharmAdder interface.
type MockCharmAdder struct {
	ctrl     *gomock.Controller
	recorder *MockCharmAdderMockRecorder
}

// MockCharmAdderMockRecorder is the mock recorder for MockCharmAdder.
type MockCharmAdderMockRecorder struct {
	mock *MockCharmAdder
}

// NewMockCharmAdder creates a new mock instance.
func NewMockCharmAdder(ctrl *gomock.Controller) *MockCharmAdder {
	mock := &MockCharmAdder{ctrl: ctrl}
	mock.recorder = &MockCharmAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmAdder) EXPECT() *MockCharmAdderMockRecorder {
	return m.recorder
}

// AddCharm mocks base method.
func (m *MockCharmAdder) AddCharm(arg0 *charm.URL, arg1 charm0.Origin, arg2 bool) (charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm0.Origin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCharm indicates an expected call of AddCharm.
func (mr *MockCharmAdderMockRecorder) AddCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharm", reflect.TypeOf((*MockCharmAdder)(nil).AddCharm), arg0, arg1, arg2)
}

// AddLocalCharm mocks base method.
func (m *MockCharmAdder) AddLocalCharm(arg0 *charm.URL, arg1 charm.Charm, arg2 bool) (*charm.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLocalCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLocalCharm indicates an expected call of AddLocalCharm.
func (mr *MockCharmAdderMockRecorder) AddLocalCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLocalCharm", reflect.TypeOf((*MockCharmAdder)(nil).AddLocalCharm), arg0, arg1, arg2)
}

// CheckCharmPlacement mocks base method.
func (m *MockCharmAdder) CheckCharmPlacement(arg0 string, arg1 *charm.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCharmPlacement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCharmPlacement indicates an expected call of CheckCharmPlacement.
func (mr *MockCharmAdderMockRecorder) CheckCharmPlacement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCharmPlacement", reflect.TypeOf((*MockCharmAdder)(nil).CheckCharmPlacement), arg0, arg1)
}

// MockCharmrepoForDeploy is a mock of CharmrepoForDeploy interface.
type MockCharmrepoForDeploy struct {
	ctrl     *gomock.Controller
	recorder *MockCharmrepoForDeployMockRecorder
}

// MockCharmrepoForDeployMockRecorder is the mock recorder for MockCharmrepoForDeploy.
type MockCharmrepoForDeployMockRecorder struct {
	mock *MockCharmrepoForDeploy
}

// NewMockCharmrepoForDeploy creates a new mock instance.
func NewMockCharmrepoForDeploy(ctrl *gomock.Controller) *MockCharmrepoForDeploy {
	mock := &MockCharmrepoForDeploy{ctrl: ctrl}
	mock.recorder = &MockCharmrepoForDeployMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmrepoForDeploy) EXPECT() *MockCharmrepoForDeployMockRecorder {
	return m.recorder
}

// GetBundle mocks base method.
func (m *MockCharmrepoForDeploy) GetBundle(arg0 *charm.URL, arg1 string) (charm.Bundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBundle", arg0, arg1)
	ret0, _ := ret[0].(charm.Bundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBundle indicates an expected call of GetBundle.
func (mr *MockCharmrepoForDeployMockRecorder) GetBundle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBundle", reflect.TypeOf((*MockCharmrepoForDeploy)(nil).GetBundle), arg0, arg1)
}

// ResolveWithPreferredChannel mocks base method.
func (m *MockCharmrepoForDeploy) ResolveWithPreferredChannel(arg0 *charm.URL, arg1 params.Channel) (*charm.URL, params.Channel, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveWithPreferredChannel", arg0, arg1)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(params.Channel)
	ret2, _ := ret[2].([]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveWithPreferredChannel indicates an expected call of ResolveWithPreferredChannel.
func (mr *MockCharmrepoForDeployMockRecorder) ResolveWithPreferredChannel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveWithPreferredChannel", reflect.TypeOf((*MockCharmrepoForDeploy)(nil).ResolveWithPreferredChannel), arg0, arg1)
}

// MockCharmsAPI is a mock of CharmsAPI interface.
type MockCharmsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCharmsAPIMockRecorder
}

// MockCharmsAPIMockRecorder is the mock recorder for MockCharmsAPI.
type MockCharmsAPIMockRecorder struct {
	mock *MockCharmsAPI
}

// NewMockCharmsAPI creates a new mock instance.
func NewMockCharmsAPI(ctrl *gomock.Controller) *MockCharmsAPI {
	mock := &MockCharmsAPI{ctrl: ctrl}
	mock.recorder = &MockCharmsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmsAPI) EXPECT() *MockCharmsAPIMockRecorder {
	return m.recorder
}

// GetDownloadInfo mocks base method.
func (m *MockCharmsAPI) GetDownloadInfo(arg0 *charm.URL, arg1 charm0.Origin) (charms.DownloadInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDownloadInfo", arg0, arg1)
	ret0, _ := ret[0].(charms.DownloadInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDownloadInfo indicates an expected call of GetDownloadInfo.
func (mr *MockCharmsAPIMockRecorder) GetDownloadInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadInfo", reflect.TypeOf((*MockCharmsAPI)(nil).GetDownloadInfo), arg0, arg1)
}

// ResolveCharms mocks base method.
func (m *MockCharmsAPI) ResolveCharms(arg0 []charms.CharmToResolve) ([]charms.ResolvedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveCharms", arg0)
	ret0, _ := ret[0].([]charms.ResolvedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveCharms indicates an expected call of ResolveCharms.
func (mr *MockCharmsAPIMockRecorder) ResolveCharms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveCharms", reflect.TypeOf((*MockCharmsAPI)(nil).ResolveCharms), arg0)
}

// MockDownloadBundleClient is a mock of DownloadBundleClient interface.
type MockDownloadBundleClient struct {
	ctrl     *gomock.Controller
	recorder *MockDownloadBundleClientMockRecorder
}

// MockDownloadBundleClientMockRecorder is the mock recorder for MockDownloadBundleClient.
type MockDownloadBundleClientMockRecorder struct {
	mock *MockDownloadBundleClient
}

// NewMockDownloadBundleClient creates a new mock instance.
func NewMockDownloadBundleClient(ctrl *gomock.Controller) *MockDownloadBundleClient {
	mock := &MockDownloadBundleClient{ctrl: ctrl}
	mock.recorder = &MockDownloadBundleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloadBundleClient) EXPECT() *MockDownloadBundleClientMockRecorder {
	return m.recorder
}

// DownloadAndReadBundle mocks base method.
func (m *MockDownloadBundleClient) DownloadAndReadBundle(arg0 context.Context, arg1 *url.URL, arg2 string, arg3 ...charmhub.DownloadOption) (charm.Bundle, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadAndReadBundle", varargs...)
	ret0, _ := ret[0].(charm.Bundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadAndReadBundle indicates an expected call of DownloadAndReadBundle.
func (mr *MockDownloadBundleClientMockRecorder) DownloadAndReadBundle(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadAndReadBundle", reflect.TypeOf((*MockDownloadBundleClient)(nil).DownloadAndReadBundle), varargs...)
}

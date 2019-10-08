// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/network/containerizer (interfaces: Container,Unit,Application,Spaces)

// Package containerizer is a generated GoMock package.
package containerizer

import (
	gomock "github.com/golang/mock/gomock"
	set "github.com/juju/collections/set"
	constraints "github.com/juju/juju/core/constraints"
	instance "github.com/juju/juju/core/instance"
	network "github.com/juju/juju/core/network"
	state "github.com/juju/juju/state"
	reflect "reflect"
)

// MockContainer is a mock of Container interface
type MockContainer struct {
	ctrl     *gomock.Controller
	recorder *MockContainerMockRecorder
}

// MockContainerMockRecorder is the mock recorder for MockContainer
type MockContainerMockRecorder struct {
	mock *MockContainer
}

// NewMockContainer creates a new mock instance
func NewMockContainer(ctrl *gomock.Controller) *MockContainer {
	mock := &MockContainer{ctrl: ctrl}
	mock.recorder = &MockContainerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContainer) EXPECT() *MockContainerMockRecorder {
	return m.recorder
}

// AllLinkLayerDevices mocks base method
func (m *MockContainer) AllLinkLayerDevices() ([]LinkLayerDevice, error) {
	ret := m.ctrl.Call(m, "AllLinkLayerDevices")
	ret0, _ := ret[0].([]LinkLayerDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllLinkLayerDevices indicates an expected call of AllLinkLayerDevices
func (mr *MockContainerMockRecorder) AllLinkLayerDevices() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllLinkLayerDevices", reflect.TypeOf((*MockContainer)(nil).AllLinkLayerDevices))
}

// AllSpaces mocks base method
func (m *MockContainer) AllSpaces() (set.Strings, error) {
	ret := m.ctrl.Call(m, "AllSpaces")
	ret0, _ := ret[0].(set.Strings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllSpaces indicates an expected call of AllSpaces
func (mr *MockContainerMockRecorder) AllSpaces() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSpaces", reflect.TypeOf((*MockContainer)(nil).AllSpaces))
}

// Constraints mocks base method
func (m *MockContainer) Constraints() (constraints.Value, error) {
	ret := m.ctrl.Call(m, "Constraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Constraints indicates an expected call of Constraints
func (mr *MockContainerMockRecorder) Constraints() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Constraints", reflect.TypeOf((*MockContainer)(nil).Constraints))
}

// ContainerType mocks base method
func (m *MockContainer) ContainerType() instance.ContainerType {
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(instance.ContainerType)
	return ret0
}

// ContainerType indicates an expected call of ContainerType
func (mr *MockContainerMockRecorder) ContainerType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockContainer)(nil).ContainerType))
}

// Id mocks base method
func (m *MockContainer) Id() string {
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id
func (mr *MockContainerMockRecorder) Id() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockContainer)(nil).Id))
}

// LinkLayerDevicesForSpaces mocks base method
func (m *MockContainer) LinkLayerDevicesForSpaces(arg0 network.SpaceInfos) (map[string][]LinkLayerDevice, error) {
	ret := m.ctrl.Call(m, "LinkLayerDevicesForSpaces", arg0)
	ret0, _ := ret[0].(map[string][]LinkLayerDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkLayerDevicesForSpaces indicates an expected call of LinkLayerDevicesForSpaces
func (mr *MockContainerMockRecorder) LinkLayerDevicesForSpaces(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkLayerDevicesForSpaces", reflect.TypeOf((*MockContainer)(nil).LinkLayerDevicesForSpaces), arg0)
}

// Raw mocks base method
func (m *MockContainer) Raw() *state.Machine {
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(*state.Machine)
	return ret0
}

// Raw indicates an expected call of Raw
func (mr *MockContainerMockRecorder) Raw() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockContainer)(nil).Raw))
}

// RemoveAllAddresses mocks base method
func (m *MockContainer) RemoveAllAddresses() error {
	ret := m.ctrl.Call(m, "RemoveAllAddresses")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllAddresses indicates an expected call of RemoveAllAddresses
func (mr *MockContainerMockRecorder) RemoveAllAddresses() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllAddresses", reflect.TypeOf((*MockContainer)(nil).RemoveAllAddresses))
}

// SetConstraints mocks base method
func (m *MockContainer) SetConstraints(arg0 constraints.Value) error {
	ret := m.ctrl.Call(m, "SetConstraints", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConstraints indicates an expected call of SetConstraints
func (mr *MockContainerMockRecorder) SetConstraints(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConstraints", reflect.TypeOf((*MockContainer)(nil).SetConstraints), arg0)
}

// SetDevicesAddresses mocks base method
func (m *MockContainer) SetDevicesAddresses(arg0 ...state.LinkLayerDeviceAddress) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetDevicesAddresses", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDevicesAddresses indicates an expected call of SetDevicesAddresses
func (mr *MockContainerMockRecorder) SetDevicesAddresses(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDevicesAddresses", reflect.TypeOf((*MockContainer)(nil).SetDevicesAddresses), arg0...)
}

// SetLinkLayerDevices mocks base method
func (m *MockContainer) SetLinkLayerDevices(arg0 ...state.LinkLayerDeviceArgs) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetLinkLayerDevices", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLinkLayerDevices indicates an expected call of SetLinkLayerDevices
func (mr *MockContainerMockRecorder) SetLinkLayerDevices(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLinkLayerDevices", reflect.TypeOf((*MockContainer)(nil).SetLinkLayerDevices), arg0...)
}

// SetParentLinkLayerDevicesBeforeTheirChildren mocks base method
func (m *MockContainer) SetParentLinkLayerDevicesBeforeTheirChildren(arg0 []state.LinkLayerDeviceArgs) error {
	ret := m.ctrl.Call(m, "SetParentLinkLayerDevicesBeforeTheirChildren", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetParentLinkLayerDevicesBeforeTheirChildren indicates an expected call of SetParentLinkLayerDevicesBeforeTheirChildren
func (mr *MockContainerMockRecorder) SetParentLinkLayerDevicesBeforeTheirChildren(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParentLinkLayerDevicesBeforeTheirChildren", reflect.TypeOf((*MockContainer)(nil).SetParentLinkLayerDevicesBeforeTheirChildren), arg0)
}

// Units mocks base method
func (m *MockContainer) Units() ([]Unit, error) {
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units
func (mr *MockContainerMockRecorder) Units() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockContainer)(nil).Units))
}

// MockUnit is a mock of Unit interface
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// Application mocks base method
func (m *MockUnit) Application() (Application, error) {
	ret := m.ctrl.Call(m, "Application")
	ret0, _ := ret[0].(Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application
func (mr *MockUnitMockRecorder) Application() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockUnit)(nil).Application))
}

// Name mocks base method
func (m *MockUnit) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockUnitMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockUnit)(nil).Name))
}

// MockApplication is a mock of Application interface
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// Charm mocks base method
func (m *MockApplication) Charm() (Charm, bool, error) {
	ret := m.ctrl.Call(m, "Charm")
	ret0, _ := ret[0].(Charm)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Charm indicates an expected call of Charm
func (mr *MockApplicationMockRecorder) Charm() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockApplication)(nil).Charm))
}

// EndpointBindings mocks base method
func (m *MockApplication) EndpointBindings() (map[string]string, error) {
	ret := m.ctrl.Call(m, "EndpointBindings")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndpointBindings indicates an expected call of EndpointBindings
func (mr *MockApplicationMockRecorder) EndpointBindings() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointBindings", reflect.TypeOf((*MockApplication)(nil).EndpointBindings))
}

// Name mocks base method
func (m *MockApplication) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockApplicationMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockApplication)(nil).Name))
}

// MockSpaces is a mock of Spaces interface
type MockSpaces struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesMockRecorder
}

// MockSpacesMockRecorder is the mock recorder for MockSpaces
type MockSpacesMockRecorder struct {
	mock *MockSpaces
}

// NewMockSpaces creates a new mock instance
func NewMockSpaces(ctrl *gomock.Controller) *MockSpaces {
	mock := &MockSpaces{ctrl: ctrl}
	mock.recorder = &MockSpacesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSpaces) EXPECT() *MockSpacesMockRecorder {
	return m.recorder
}

// GetByID mocks base method
func (m *MockSpaces) GetByID(arg0 string) (network.SpaceInfo, error) {
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockSpacesMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockSpaces)(nil).GetByID), arg0)
}

// GetByName mocks base method
func (m *MockSpaces) GetByName(arg0 string) (network.SpaceInfo, error) {
	ret := m.ctrl.Call(m, "GetByName", arg0)
	ret0, _ := ret[0].(network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockSpacesMockRecorder) GetByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockSpaces)(nil).GetByName), arg0)
}

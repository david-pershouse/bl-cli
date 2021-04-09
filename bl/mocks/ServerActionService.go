// Code generated by MockGen. DO NOT EDIT.
// Source: server_actions.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	bl "github.com/binarylane/bl-cli/bl"
	gomock "github.com/golang/mock/gomock"
)

// MockServerActionsService is a mock of ServerActionsService interface.
type MockServerActionsService struct {
	ctrl     *gomock.Controller
	recorder *MockServerActionsServiceMockRecorder
}

// MockServerActionsServiceMockRecorder is the mock recorder for MockServerActionsService.
type MockServerActionsServiceMockRecorder struct {
	mock *MockServerActionsService
}

// NewMockServerActionsService creates a new mock instance.
func NewMockServerActionsService(ctrl *gomock.Controller) *MockServerActionsService {
	mock := &MockServerActionsService{ctrl: ctrl}
	mock.recorder = &MockServerActionsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerActionsService) EXPECT() *MockServerActionsServiceMockRecorder {
	return m.recorder
}

// Shutdown mocks base method.
func (m *MockServerActionsService) Shutdown(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockServerActionsServiceMockRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockServerActionsService)(nil).Shutdown), arg0)
}

// ShutdownByTag mocks base method.
func (m *MockServerActionsService) ShutdownByTag(arg0 string) (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownByTag", arg0)
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShutdownByTag indicates an expected call of ShutdownByTag.
func (mr *MockServerActionsServiceMockRecorder) ShutdownByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownByTag", reflect.TypeOf((*MockServerActionsService)(nil).ShutdownByTag), arg0)
}

// PowerOff mocks base method.
func (m *MockServerActionsService) PowerOff(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOff", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOff indicates an expected call of PowerOff.
func (mr *MockServerActionsServiceMockRecorder) PowerOff(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOff", reflect.TypeOf((*MockServerActionsService)(nil).PowerOff), arg0)
}

// PowerOffByTag mocks base method.
func (m *MockServerActionsService) PowerOffByTag(arg0 string) (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOffByTag", arg0)
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOffByTag indicates an expected call of PowerOffByTag.
func (mr *MockServerActionsServiceMockRecorder) PowerOffByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOffByTag", reflect.TypeOf((*MockServerActionsService)(nil).PowerOffByTag), arg0)
}

// PowerOn mocks base method.
func (m *MockServerActionsService) PowerOn(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOn", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOn indicates an expected call of PowerOn.
func (mr *MockServerActionsServiceMockRecorder) PowerOn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOn", reflect.TypeOf((*MockServerActionsService)(nil).PowerOn), arg0)
}

// PowerOnByTag mocks base method.
func (m *MockServerActionsService) PowerOnByTag(arg0 string) (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOnByTag", arg0)
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOnByTag indicates an expected call of PowerOnByTag.
func (mr *MockServerActionsServiceMockRecorder) PowerOnByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOnByTag", reflect.TypeOf((*MockServerActionsService)(nil).PowerOnByTag), arg0)
}

// PowerCycle mocks base method.
func (m *MockServerActionsService) PowerCycle(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerCycle", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerCycle indicates an expected call of PowerCycle.
func (mr *MockServerActionsServiceMockRecorder) PowerCycle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerCycle", reflect.TypeOf((*MockServerActionsService)(nil).PowerCycle), arg0)
}

// PowerCycleByTag mocks base method.
func (m *MockServerActionsService) PowerCycleByTag(arg0 string) (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerCycleByTag", arg0)
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerCycleByTag indicates an expected call of PowerCycleByTag.
func (mr *MockServerActionsServiceMockRecorder) PowerCycleByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerCycleByTag", reflect.TypeOf((*MockServerActionsService)(nil).PowerCycleByTag), arg0)
}

// Reboot mocks base method.
func (m *MockServerActionsService) Reboot(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reboot", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reboot indicates an expected call of Reboot.
func (mr *MockServerActionsServiceMockRecorder) Reboot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reboot", reflect.TypeOf((*MockServerActionsService)(nil).Reboot), arg0)
}

// Restore mocks base method.
func (m *MockServerActionsService) Restore(arg0, arg1 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", arg0, arg1)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Restore indicates an expected call of Restore.
func (mr *MockServerActionsServiceMockRecorder) Restore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockServerActionsService)(nil).Restore), arg0, arg1)
}

// Resize mocks base method.
func (m *MockServerActionsService) Resize(arg0 int, arg1 string, arg2 bool) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resize", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resize indicates an expected call of Resize.
func (mr *MockServerActionsServiceMockRecorder) Resize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockServerActionsService)(nil).Resize), arg0, arg1, arg2)
}

// Rename mocks base method.
func (m *MockServerActionsService) Rename(arg0 int, arg1 string) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", arg0, arg1)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rename indicates an expected call of Rename.
func (mr *MockServerActionsServiceMockRecorder) Rename(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockServerActionsService)(nil).Rename), arg0, arg1)
}

// Snapshot mocks base method.
func (m *MockServerActionsService) Snapshot(arg0 int, arg1 string) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", arg0, arg1)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockServerActionsServiceMockRecorder) Snapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockServerActionsService)(nil).Snapshot), arg0, arg1)
}

// SnapshotByTag mocks base method.
func (m *MockServerActionsService) SnapshotByTag(arg0, arg1 string) (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotByTag", arg0, arg1)
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotByTag indicates an expected call of SnapshotByTag.
func (mr *MockServerActionsServiceMockRecorder) SnapshotByTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotByTag", reflect.TypeOf((*MockServerActionsService)(nil).SnapshotByTag), arg0, arg1)
}

// EnableBackups mocks base method.
func (m *MockServerActionsService) EnableBackups(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableBackups", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableBackups indicates an expected call of EnableBackups.
func (mr *MockServerActionsServiceMockRecorder) EnableBackups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableBackups", reflect.TypeOf((*MockServerActionsService)(nil).EnableBackups), arg0)
}

// EnableBackupsByTag mocks base method.
func (m *MockServerActionsService) EnableBackupsByTag(arg0 string) (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableBackupsByTag", arg0)
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableBackupsByTag indicates an expected call of EnableBackupsByTag.
func (mr *MockServerActionsServiceMockRecorder) EnableBackupsByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableBackupsByTag", reflect.TypeOf((*MockServerActionsService)(nil).EnableBackupsByTag), arg0)
}

// DisableBackups mocks base method.
func (m *MockServerActionsService) DisableBackups(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableBackups", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableBackups indicates an expected call of DisableBackups.
func (mr *MockServerActionsServiceMockRecorder) DisableBackups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableBackups", reflect.TypeOf((*MockServerActionsService)(nil).DisableBackups), arg0)
}

// DisableBackupsByTag mocks base method.
func (m *MockServerActionsService) DisableBackupsByTag(arg0 string) (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableBackupsByTag", arg0)
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableBackupsByTag indicates an expected call of DisableBackupsByTag.
func (mr *MockServerActionsServiceMockRecorder) DisableBackupsByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableBackupsByTag", reflect.TypeOf((*MockServerActionsService)(nil).DisableBackupsByTag), arg0)
}

// PasswordReset mocks base method.
func (m *MockServerActionsService) PasswordReset(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordReset", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PasswordReset indicates an expected call of PasswordReset.
func (mr *MockServerActionsServiceMockRecorder) PasswordReset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordReset", reflect.TypeOf((*MockServerActionsService)(nil).PasswordReset), arg0)
}

// RebuildByImageID mocks base method.
func (m *MockServerActionsService) RebuildByImageID(arg0, arg1 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RebuildByImageID", arg0, arg1)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RebuildByImageID indicates an expected call of RebuildByImageID.
func (mr *MockServerActionsServiceMockRecorder) RebuildByImageID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebuildByImageID", reflect.TypeOf((*MockServerActionsService)(nil).RebuildByImageID), arg0, arg1)
}

// RebuildByImageSlug mocks base method.
func (m *MockServerActionsService) RebuildByImageSlug(arg0 int, arg1 string) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RebuildByImageSlug", arg0, arg1)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RebuildByImageSlug indicates an expected call of RebuildByImageSlug.
func (mr *MockServerActionsServiceMockRecorder) RebuildByImageSlug(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebuildByImageSlug", reflect.TypeOf((*MockServerActionsService)(nil).RebuildByImageSlug), arg0, arg1)
}

// ChangeKernel mocks base method.
func (m *MockServerActionsService) ChangeKernel(arg0, arg1 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeKernel", arg0, arg1)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeKernel indicates an expected call of ChangeKernel.
func (mr *MockServerActionsServiceMockRecorder) ChangeKernel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeKernel", reflect.TypeOf((*MockServerActionsService)(nil).ChangeKernel), arg0, arg1)
}

// EnableIPv6 mocks base method.
func (m *MockServerActionsService) EnableIPv6(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableIPv6", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableIPv6 indicates an expected call of EnableIPv6.
func (mr *MockServerActionsServiceMockRecorder) EnableIPv6(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableIPv6", reflect.TypeOf((*MockServerActionsService)(nil).EnableIPv6), arg0)
}

// EnableIPv6ByTag mocks base method.
func (m *MockServerActionsService) EnableIPv6ByTag(arg0 string) (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableIPv6ByTag", arg0)
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableIPv6ByTag indicates an expected call of EnableIPv6ByTag.
func (mr *MockServerActionsServiceMockRecorder) EnableIPv6ByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableIPv6ByTag", reflect.TypeOf((*MockServerActionsService)(nil).EnableIPv6ByTag), arg0)
}

// EnablePrivateNetworking mocks base method.
func (m *MockServerActionsService) EnablePrivateNetworking(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePrivateNetworking", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnablePrivateNetworking indicates an expected call of EnablePrivateNetworking.
func (mr *MockServerActionsServiceMockRecorder) EnablePrivateNetworking(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePrivateNetworking", reflect.TypeOf((*MockServerActionsService)(nil).EnablePrivateNetworking), arg0)
}

// EnablePrivateNetworkingByTag mocks base method.
func (m *MockServerActionsService) EnablePrivateNetworkingByTag(arg0 string) (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePrivateNetworkingByTag", arg0)
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnablePrivateNetworkingByTag indicates an expected call of EnablePrivateNetworkingByTag.
func (mr *MockServerActionsServiceMockRecorder) EnablePrivateNetworkingByTag(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePrivateNetworkingByTag", reflect.TypeOf((*MockServerActionsService)(nil).EnablePrivateNetworkingByTag), arg0)
}

// Get mocks base method.
func (m *MockServerActionsService) Get(arg0, arg1 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServerActionsServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServerActionsService)(nil).Get), arg0, arg1)
}

// GetByURI mocks base method.
func (m *MockServerActionsService) GetByURI(arg0 string) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByURI", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByURI indicates an expected call of GetByURI.
func (mr *MockServerActionsServiceMockRecorder) GetByURI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByURI", reflect.TypeOf((*MockServerActionsService)(nil).GetByURI), arg0)
}

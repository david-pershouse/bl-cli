// Code generated by MockGen. DO NOT EDIT.
// Source: sshkeys.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	bl "github.com/binarylane/bl-cli/bl"
	binarylane "github.com/binarylane/go-binarylane"
	gomock "github.com/golang/mock/gomock"
)

// MockKeysService is a mock of KeysService interface.
type MockKeysService struct {
	ctrl     *gomock.Controller
	recorder *MockKeysServiceMockRecorder
}

// MockKeysServiceMockRecorder is the mock recorder for MockKeysService.
type MockKeysServiceMockRecorder struct {
	mock *MockKeysService
}

// NewMockKeysService creates a new mock instance.
func NewMockKeysService(ctrl *gomock.Controller) *MockKeysService {
	mock := &MockKeysService{ctrl: ctrl}
	mock.recorder = &MockKeysServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeysService) EXPECT() *MockKeysServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockKeysService) Create(kcr *binarylane.KeyCreateRequest) (*bl.SSHKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", kcr)
	ret0, _ := ret[0].(*bl.SSHKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockKeysServiceMockRecorder) Create(kcr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockKeysService)(nil).Create), kcr)
}

// Delete mocks base method.
func (m *MockKeysService) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockKeysServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKeysService)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockKeysService) Get(id string) (*bl.SSHKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*bl.SSHKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKeysServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKeysService)(nil).Get), id)
}

// List mocks base method.
func (m *MockKeysService) List() (bl.SSHKeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(bl.SSHKeys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockKeysServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKeysService)(nil).List))
}

// Update mocks base method.
func (m *MockKeysService) Update(id string, kur *binarylane.KeyUpdateRequest) (*bl.SSHKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, kur)
	ret0, _ := ret[0].(*bl.SSHKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockKeysServiceMockRecorder) Update(id, kur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockKeysService)(nil).Update), id, kur)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: vpcs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	bl "github.com/binarylane/bl-cli/bl"
	binarylane "github.com/binarylane/go-binarylane"
	gomock "github.com/golang/mock/gomock"
)

// MockVPCsService is a mock of VPCsService interface.
type MockVPCsService struct {
	ctrl     *gomock.Controller
	recorder *MockVPCsServiceMockRecorder
}

// MockVPCsServiceMockRecorder is the mock recorder for MockVPCsService.
type MockVPCsServiceMockRecorder struct {
	mock *MockVPCsService
}

// NewMockVPCsService creates a new mock instance.
func NewMockVPCsService(ctrl *gomock.Controller) *MockVPCsService {
	mock := &MockVPCsService{ctrl: ctrl}
	mock.recorder = &MockVPCsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVPCsService) EXPECT() *MockVPCsServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockVPCsService) Get(id int) (*bl.VPC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*bl.VPC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVPCsServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVPCsService)(nil).Get), id)
}

// List mocks base method.
func (m *MockVPCsService) List() (bl.VPCs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(bl.VPCs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVPCsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVPCsService)(nil).List))
}

// Create mocks base method.
func (m *MockVPCsService) Create(vpcr *binarylane.VPCCreateRequest) (*bl.VPC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", vpcr)
	ret0, _ := ret[0].(*bl.VPC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVPCsServiceMockRecorder) Create(vpcr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVPCsService)(nil).Create), vpcr)
}

// Update mocks base method.
func (m *MockVPCsService) Update(id int, vpcr *binarylane.VPCUpdateRequest) (*bl.VPC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, vpcr)
	ret0, _ := ret[0].(*bl.VPC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockVPCsServiceMockRecorder) Update(id, vpcr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVPCsService)(nil).Update), id, vpcr)
}

// Delete mocks base method.
func (m *MockVPCsService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVPCsServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVPCsService)(nil).Delete), id)
}

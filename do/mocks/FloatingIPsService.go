// Code generated by MockGen. DO NOT EDIT.
// Source: floating_ips.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "git.mammoth.com.au/github/bl-cli/do"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockFloatingIPsService is a mock of FloatingIPsService interface.
type MockFloatingIPsService struct {
	ctrl     *gomock.Controller
	recorder *MockFloatingIPsServiceMockRecorder
}

// MockFloatingIPsServiceMockRecorder is the mock recorder for MockFloatingIPsService.
type MockFloatingIPsServiceMockRecorder struct {
	mock *MockFloatingIPsService
}

// NewMockFloatingIPsService creates a new mock instance.
func NewMockFloatingIPsService(ctrl *gomock.Controller) *MockFloatingIPsService {
	mock := &MockFloatingIPsService{ctrl: ctrl}
	mock.recorder = &MockFloatingIPsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFloatingIPsService) EXPECT() *MockFloatingIPsServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockFloatingIPsService) List() (do.FloatingIPs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(do.FloatingIPs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockFloatingIPsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFloatingIPsService)(nil).List))
}

// Get mocks base method.
func (m *MockFloatingIPsService) Get(ip string) (*do.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ip)
	ret0, _ := ret[0].(*do.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFloatingIPsServiceMockRecorder) Get(ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFloatingIPsService)(nil).Get), ip)
}

// Create mocks base method.
func (m *MockFloatingIPsService) Create(ficr *godo.FloatingIPCreateRequest) (*do.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ficr)
	ret0, _ := ret[0].(*do.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFloatingIPsServiceMockRecorder) Create(ficr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFloatingIPsService)(nil).Create), ficr)
}

// Delete mocks base method.
func (m *MockFloatingIPsService) Delete(ip string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ip)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFloatingIPsServiceMockRecorder) Delete(ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFloatingIPsService)(nil).Delete), ip)
}

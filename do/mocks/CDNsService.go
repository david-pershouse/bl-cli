// Code generated by MockGen. DO NOT EDIT.
// Source: cdns.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "git.mammoth.com.au/github/bl-cli/do"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockCDNsService is a mock of CDNsService interface.
type MockCDNsService struct {
	ctrl     *gomock.Controller
	recorder *MockCDNsServiceMockRecorder
}

// MockCDNsServiceMockRecorder is the mock recorder for MockCDNsService.
type MockCDNsServiceMockRecorder struct {
	mock *MockCDNsService
}

// NewMockCDNsService creates a new mock instance.
func NewMockCDNsService(ctrl *gomock.Controller) *MockCDNsService {
	mock := &MockCDNsService{ctrl: ctrl}
	mock.recorder = &MockCDNsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCDNsService) EXPECT() *MockCDNsServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockCDNsService) List() ([]do.CDN, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]do.CDN)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCDNsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCDNsService)(nil).List))
}

// Get mocks base method.
func (m *MockCDNsService) Get(arg0 string) (*do.CDN, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*do.CDN)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCDNsServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCDNsService)(nil).Get), arg0)
}

// Create mocks base method.
func (m *MockCDNsService) Create(arg0 *godo.CDNCreateRequest) (*do.CDN, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*do.CDN)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCDNsServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCDNsService)(nil).Create), arg0)
}

// UpdateTTL mocks base method.
func (m *MockCDNsService) UpdateTTL(arg0 string, arg1 *godo.CDNUpdateTTLRequest) (*do.CDN, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTTL", arg0, arg1)
	ret0, _ := ret[0].(*do.CDN)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTTL indicates an expected call of UpdateTTL.
func (mr *MockCDNsServiceMockRecorder) UpdateTTL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTTL", reflect.TypeOf((*MockCDNsService)(nil).UpdateTTL), arg0, arg1)
}

// UpdateCustomDomain mocks base method.
func (m *MockCDNsService) UpdateCustomDomain(arg0 string, arg1 *godo.CDNUpdateCustomDomainRequest) (*do.CDN, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomDomain", arg0, arg1)
	ret0, _ := ret[0].(*do.CDN)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCustomDomain indicates an expected call of UpdateCustomDomain.
func (mr *MockCDNsServiceMockRecorder) UpdateCustomDomain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomDomain", reflect.TypeOf((*MockCDNsService)(nil).UpdateCustomDomain), arg0, arg1)
}

// FlushCache mocks base method.
func (m *MockCDNsService) FlushCache(arg0 string, arg1 *godo.CDNFlushCacheRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushCache", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushCache indicates an expected call of FlushCache.
func (mr *MockCDNsServiceMockRecorder) FlushCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushCache", reflect.TypeOf((*MockCDNsService)(nil).FlushCache), arg0, arg1)
}

// Delete mocks base method.
func (m *MockCDNsService) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCDNsServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCDNsService)(nil).Delete), arg0)
}

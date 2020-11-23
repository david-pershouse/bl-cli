// Code generated by MockGen. DO NOT EDIT.
// Source: domains.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	bl "git.mammoth.com.au/github/bl-cli/bl"
	binarylane "git.mammoth.com.au/github/go-binarylane"
	gomock "github.com/golang/mock/gomock"
)

// MockDomainsService is a mock of DomainsService interface.
type MockDomainsService struct {
	ctrl     *gomock.Controller
	recorder *MockDomainsServiceMockRecorder
}

// MockDomainsServiceMockRecorder is the mock recorder for MockDomainsService.
type MockDomainsServiceMockRecorder struct {
	mock *MockDomainsService
}

// NewMockDomainsService creates a new mock instance.
func NewMockDomainsService(ctrl *gomock.Controller) *MockDomainsService {
	mock := &MockDomainsService{ctrl: ctrl}
	mock.recorder = &MockDomainsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomainsService) EXPECT() *MockDomainsServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockDomainsService) List() (bl.Domains, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(bl.Domains)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDomainsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDomainsService)(nil).List))
}

// Get mocks base method.
func (m *MockDomainsService) Get(arg0 string) (*bl.Domain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*bl.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDomainsServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDomainsService)(nil).Get), arg0)
}

// Create mocks base method.
func (m *MockDomainsService) Create(arg0 *binarylane.DomainCreateRequest) (*bl.Domain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*bl.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDomainsServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDomainsService)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockDomainsService) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDomainsServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDomainsService)(nil).Delete), arg0)
}

// Records mocks base method.
func (m *MockDomainsService) Records(arg0 string) (bl.DomainRecords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Records", arg0)
	ret0, _ := ret[0].(bl.DomainRecords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Records indicates an expected call of Records.
func (mr *MockDomainsServiceMockRecorder) Records(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Records", reflect.TypeOf((*MockDomainsService)(nil).Records), arg0)
}

// Record mocks base method.
func (m *MockDomainsService) Record(arg0 string, arg1 int) (*bl.DomainRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Record", arg0, arg1)
	ret0, _ := ret[0].(*bl.DomainRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Record indicates an expected call of Record.
func (mr *MockDomainsServiceMockRecorder) Record(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockDomainsService)(nil).Record), arg0, arg1)
}

// DeleteRecord mocks base method.
func (m *MockDomainsService) DeleteRecord(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecord", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecord indicates an expected call of DeleteRecord.
func (mr *MockDomainsServiceMockRecorder) DeleteRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecord", reflect.TypeOf((*MockDomainsService)(nil).DeleteRecord), arg0, arg1)
}

// EditRecord mocks base method.
func (m *MockDomainsService) EditRecord(arg0 string, arg1 int, arg2 *bl.DomainRecordEditRequest) (*bl.DomainRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(*bl.DomainRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditRecord indicates an expected call of EditRecord.
func (mr *MockDomainsServiceMockRecorder) EditRecord(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditRecord", reflect.TypeOf((*MockDomainsService)(nil).EditRecord), arg0, arg1, arg2)
}

// CreateRecord mocks base method.
func (m *MockDomainsService) CreateRecord(arg0 string, arg1 *bl.DomainRecordEditRequest) (*bl.DomainRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", arg0, arg1)
	ret0, _ := ret[0].(*bl.DomainRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRecord indicates an expected call of CreateRecord.
func (mr *MockDomainsServiceMockRecorder) CreateRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockDomainsService)(nil).CreateRecord), arg0, arg1)
}

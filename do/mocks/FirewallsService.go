// Code generated by MockGen. DO NOT EDIT.
// Source: firewalls.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "git.mammoth.com.au/github/bl-cli/do"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockFirewallsService is a mock of FirewallsService interface.
type MockFirewallsService struct {
	ctrl     *gomock.Controller
	recorder *MockFirewallsServiceMockRecorder
}

// MockFirewallsServiceMockRecorder is the mock recorder for MockFirewallsService.
type MockFirewallsServiceMockRecorder struct {
	mock *MockFirewallsService
}

// NewMockFirewallsService creates a new mock instance.
func NewMockFirewallsService(ctrl *gomock.Controller) *MockFirewallsService {
	mock := &MockFirewallsService{ctrl: ctrl}
	mock.recorder = &MockFirewallsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirewallsService) EXPECT() *MockFirewallsServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockFirewallsService) Get(fID string) (*do.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", fID)
	ret0, _ := ret[0].(*do.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFirewallsServiceMockRecorder) Get(fID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFirewallsService)(nil).Get), fID)
}

// Create mocks base method.
func (m *MockFirewallsService) Create(fr *godo.FirewallRequest) (*do.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", fr)
	ret0, _ := ret[0].(*do.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFirewallsServiceMockRecorder) Create(fr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFirewallsService)(nil).Create), fr)
}

// Update mocks base method.
func (m *MockFirewallsService) Update(fID string, fr *godo.FirewallRequest) (*do.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", fID, fr)
	ret0, _ := ret[0].(*do.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockFirewallsServiceMockRecorder) Update(fID, fr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFirewallsService)(nil).Update), fID, fr)
}

// List mocks base method.
func (m *MockFirewallsService) List() (do.Firewalls, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(do.Firewalls)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockFirewallsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFirewallsService)(nil).List))
}

// ListByDroplet mocks base method.
func (m *MockFirewallsService) ListByDroplet(dID int) (do.Firewalls, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDroplet", dID)
	ret0, _ := ret[0].(do.Firewalls)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByDroplet indicates an expected call of ListByDroplet.
func (mr *MockFirewallsServiceMockRecorder) ListByDroplet(dID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDroplet", reflect.TypeOf((*MockFirewallsService)(nil).ListByDroplet), dID)
}

// Delete mocks base method.
func (m *MockFirewallsService) Delete(fID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", fID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFirewallsServiceMockRecorder) Delete(fID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFirewallsService)(nil).Delete), fID)
}

// AddDroplets mocks base method.
func (m *MockFirewallsService) AddDroplets(fID string, dIDs ...int) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{fID}
	for _, a := range dIDs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddDroplets", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDroplets indicates an expected call of AddDroplets.
func (mr *MockFirewallsServiceMockRecorder) AddDroplets(fID interface{}, dIDs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{fID}, dIDs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDroplets", reflect.TypeOf((*MockFirewallsService)(nil).AddDroplets), varargs...)
}

// RemoveDroplets mocks base method.
func (m *MockFirewallsService) RemoveDroplets(fID string, dIDs ...int) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{fID}
	for _, a := range dIDs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveDroplets", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDroplets indicates an expected call of RemoveDroplets.
func (mr *MockFirewallsServiceMockRecorder) RemoveDroplets(fID interface{}, dIDs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{fID}, dIDs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDroplets", reflect.TypeOf((*MockFirewallsService)(nil).RemoveDroplets), varargs...)
}

// AddTags mocks base method.
func (m *MockFirewallsService) AddTags(fID string, tags ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{fID}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddTags", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTags indicates an expected call of AddTags.
func (mr *MockFirewallsServiceMockRecorder) AddTags(fID interface{}, tags ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{fID}, tags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTags", reflect.TypeOf((*MockFirewallsService)(nil).AddTags), varargs...)
}

// RemoveTags mocks base method.
func (m *MockFirewallsService) RemoveTags(fID string, tags ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{fID}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveTags", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTags indicates an expected call of RemoveTags.
func (mr *MockFirewallsServiceMockRecorder) RemoveTags(fID interface{}, tags ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{fID}, tags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTags", reflect.TypeOf((*MockFirewallsService)(nil).RemoveTags), varargs...)
}

// AddRules mocks base method.
func (m *MockFirewallsService) AddRules(fID string, rr *godo.FirewallRulesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRules", fID, rr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRules indicates an expected call of AddRules.
func (mr *MockFirewallsServiceMockRecorder) AddRules(fID, rr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRules", reflect.TypeOf((*MockFirewallsService)(nil).AddRules), fID, rr)
}

// RemoveRules mocks base method.
func (m *MockFirewallsService) RemoveRules(fID string, rr *godo.FirewallRulesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRules", fID, rr)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRules indicates an expected call of RemoveRules.
func (mr *MockFirewallsServiceMockRecorder) RemoveRules(fID, rr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRules", reflect.TypeOf((*MockFirewallsService)(nil).RemoveRules), fID, rr)
}

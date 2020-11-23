// Code generated by MockGen. DO NOT EDIT.
// Source: actions.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	bl "git.mammoth.com.au/github/bl-cli/bl"
	gomock "github.com/golang/mock/gomock"
)

// MockActionsService is a mock of ActionsService interface.
type MockActionsService struct {
	ctrl     *gomock.Controller
	recorder *MockActionsServiceMockRecorder
}

// MockActionsServiceMockRecorder is the mock recorder for MockActionsService.
type MockActionsServiceMockRecorder struct {
	mock *MockActionsService
}

// NewMockActionsService creates a new mock instance.
func NewMockActionsService(ctrl *gomock.Controller) *MockActionsService {
	mock := &MockActionsService{ctrl: ctrl}
	mock.recorder = &MockActionsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionsService) EXPECT() *MockActionsServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockActionsService) List() (bl.Actions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(bl.Actions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockActionsServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockActionsService)(nil).List))
}

// Get mocks base method.
func (m *MockActionsService) Get(arg0 int) (*bl.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*bl.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockActionsServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockActionsService)(nil).Get), arg0)
}

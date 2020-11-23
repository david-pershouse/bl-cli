// Code generated by MockGen. DO NOT EDIT.
// Source: images.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	bl "git.mammoth.com.au/github/bl-cli/bl"
	binarylane "git.mammoth.com.au/github/go-binarylane"
	gomock "github.com/golang/mock/gomock"
)

// MockImagesService is a mock of ImagesService interface.
type MockImagesService struct {
	ctrl     *gomock.Controller
	recorder *MockImagesServiceMockRecorder
}

// MockImagesServiceMockRecorder is the mock recorder for MockImagesService.
type MockImagesServiceMockRecorder struct {
	mock *MockImagesService
}

// NewMockImagesService creates a new mock instance.
func NewMockImagesService(ctrl *gomock.Controller) *MockImagesService {
	mock := &MockImagesService{ctrl: ctrl}
	mock.recorder = &MockImagesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImagesService) EXPECT() *MockImagesServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockImagesService) List(public bool) (bl.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", public)
	ret0, _ := ret[0].(bl.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockImagesServiceMockRecorder) List(public interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockImagesService)(nil).List), public)
}

// ListDistribution mocks base method.
func (m *MockImagesService) ListDistribution(public bool) (bl.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDistribution", public)
	ret0, _ := ret[0].(bl.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDistribution indicates an expected call of ListDistribution.
func (mr *MockImagesServiceMockRecorder) ListDistribution(public interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDistribution", reflect.TypeOf((*MockImagesService)(nil).ListDistribution), public)
}

// ListApplication mocks base method.
func (m *MockImagesService) ListApplication(public bool) (bl.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplication", public)
	ret0, _ := ret[0].(bl.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplication indicates an expected call of ListApplication.
func (mr *MockImagesServiceMockRecorder) ListApplication(public interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplication", reflect.TypeOf((*MockImagesService)(nil).ListApplication), public)
}

// ListUser mocks base method.
func (m *MockImagesService) ListUser(public bool) (bl.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUser", public)
	ret0, _ := ret[0].(bl.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUser indicates an expected call of ListUser.
func (mr *MockImagesServiceMockRecorder) ListUser(public interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser", reflect.TypeOf((*MockImagesService)(nil).ListUser), public)
}

// GetByID mocks base method.
func (m *MockImagesService) GetByID(id int) (*bl.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*bl.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockImagesServiceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockImagesService)(nil).GetByID), id)
}

// GetBySlug mocks base method.
func (m *MockImagesService) GetBySlug(slug string) (*bl.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySlug", slug)
	ret0, _ := ret[0].(*bl.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug.
func (mr *MockImagesServiceMockRecorder) GetBySlug(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockImagesService)(nil).GetBySlug), slug)
}

// Update mocks base method.
func (m *MockImagesService) Update(id int, iur *binarylane.ImageUpdateRequest) (*bl.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, iur)
	ret0, _ := ret[0].(*bl.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockImagesServiceMockRecorder) Update(id, iur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockImagesService)(nil).Update), id, iur)
}

// Delete mocks base method.
func (m *MockImagesService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockImagesServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockImagesService)(nil).Delete), id)
}

// Create mocks base method.
func (m *MockImagesService) Create(icr *binarylane.CustomImageCreateRequest) (*bl.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", icr)
	ret0, _ := ret[0].(*bl.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockImagesServiceMockRecorder) Create(icr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockImagesService)(nil).Create), icr)
}

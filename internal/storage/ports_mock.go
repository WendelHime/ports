// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/WendelHime/ports/internal/storage (interfaces: PortRepository)

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	reflect "reflect"

	models "github.com/WendelHime/ports/internal/shared/models"
	gomock "github.com/golang/mock/gomock"
)

// MockPortRepository is a mock of PortRepository interface.
type MockPortRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPortRepositoryMockRecorder
}

// MockPortRepositoryMockRecorder is the mock recorder for MockPortRepository.
type MockPortRepositoryMockRecorder struct {
	mock *MockPortRepository
}

// NewMockPortRepository creates a new mock instance.
func NewMockPortRepository(ctrl *gomock.Controller) *MockPortRepository {
	mock := &MockPortRepository{ctrl: ctrl}
	mock.recorder = &MockPortRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortRepository) EXPECT() *MockPortRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPortRepository) Create(arg0 context.Context, arg1 models.Port) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPortRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPortRepository)(nil).Create), arg0, arg1)
}

// Get mocks base method.
func (m *MockPortRepository) Get(arg0 context.Context, arg1 string) (models.Port, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(models.Port)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPortRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPortRepository)(nil).Get), arg0, arg1)
}

// Update mocks base method.
func (m *MockPortRepository) Update(arg0 context.Context, arg1 models.Port) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPortRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPortRepository)(nil).Update), arg0, arg1)
}
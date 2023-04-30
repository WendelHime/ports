// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/WendelHime/ports/internal/logic (interfaces: PortDomainService)

// Package logic is a generated GoMock package.
package logic

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPortDomainService is a mock of PortDomainService interface.
type MockPortDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockPortDomainServiceMockRecorder
}

// MockPortDomainServiceMockRecorder is the mock recorder for MockPortDomainService.
type MockPortDomainServiceMockRecorder struct {
	mock *MockPortDomainService
}

// NewMockPortDomainService creates a new mock instance.
func NewMockPortDomainService(ctrl *gomock.Controller) *MockPortDomainService {
	mock := &MockPortDomainService{ctrl: ctrl}
	mock.recorder = &MockPortDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortDomainService) EXPECT() *MockPortDomainServiceMockRecorder {
	return m.recorder
}

// SyncPorts mocks base method.
func (m *MockPortDomainService) SyncPorts(arg0 context.Context, arg1 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncPorts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncPorts indicates an expected call of SyncPorts.
func (mr *MockPortDomainServiceMockRecorder) SyncPorts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPorts", reflect.TypeOf((*MockPortDomainService)(nil).SyncPorts), arg0, arg1)
}
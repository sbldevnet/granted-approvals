// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/common-fate/granted-approvals/pkg/service/accesssvc (interfaces: Granter)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	access "github.com/common-fate/granted-approvals/pkg/access"
	grantsvc "github.com/common-fate/granted-approvals/pkg/service/grantsvc"
	gomock "github.com/golang/mock/gomock"
)

// MockGranter is a mock of Granter interface.
type MockGranter struct {
	ctrl     *gomock.Controller
	recorder *MockGranterMockRecorder
}

// MockGranterMockRecorder is the mock recorder for MockGranter.
type MockGranterMockRecorder struct {
	mock *MockGranter
}

// NewMockGranter creates a new mock instance.
func NewMockGranter(ctrl *gomock.Controller) *MockGranter {
	mock := &MockGranter{ctrl: ctrl}
	mock.recorder = &MockGranterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGranter) EXPECT() *MockGranterMockRecorder {
	return m.recorder
}

// CreateGrant mocks base method.
func (m *MockGranter) CreateGrant(arg0 context.Context, arg1 grantsvc.CreateGrantOpts) (*access.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGrant", arg0, arg1)
	ret0, _ := ret[0].(*access.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGrant indicates an expected call of CreateGrant.
func (mr *MockGranterMockRecorder) CreateGrant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGrant", reflect.TypeOf((*MockGranter)(nil).CreateGrant), arg0, arg1)
}

// RevokeGrant mocks base method.
func (m *MockGranter) RevokeGrant(arg0 context.Context, arg1 grantsvc.RevokeGrantOpts) (*access.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeGrant", arg0, arg1)
	ret0, _ := ret[0].(*access.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeGrant indicates an expected call of RevokeGrant.
func (mr *MockGranterMockRecorder) RevokeGrant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeGrant", reflect.TypeOf((*MockGranter)(nil).RevokeGrant), arg0, arg1)
}

// ValidateGrant mocks base method.
func (m *MockGranter) ValidateGrant(arg0 context.Context, arg1 grantsvc.CreateGrantOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGrant", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateGrant indicates an expected call of ValidateGrant.
func (mr *MockGranterMockRecorder) ValidateGrant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGrant", reflect.TypeOf((*MockGranter)(nil).ValidateGrant), arg0, arg1)
}

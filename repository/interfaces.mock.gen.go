// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockRepositoryInterface) CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, input)
	ret0, _ := ret[0].(*CreateUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryInterfaceMockRecorder) CreateUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateUser), ctx, input)
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}

// GetUserById mocks base method.
func (m *MockRepositoryInterface) GetUserById(ctx context.Context, id uint) (*UserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(*UserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserById), ctx, id)
}

// GetUserByPhoneNumber mocks base method.
func (m *MockRepositoryInterface) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*UserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhoneNumber", ctx, phoneNumber)
	ret0, _ := ret[0].(*UserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhoneNumber indicates an expected call of GetUserByPhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserByPhoneNumber(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserByPhoneNumber), ctx, phoneNumber)
}

// IncrementUserLoginCount mocks base method.
func (m *MockRepositoryInterface) IncrementUserLoginCount(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementUserLoginCount", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementUserLoginCount indicates an expected call of IncrementUserLoginCount.
func (mr *MockRepositoryInterfaceMockRecorder) IncrementUserLoginCount(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementUserLoginCount", reflect.TypeOf((*MockRepositoryInterface)(nil).IncrementUserLoginCount), ctx, id)
}

// IsPhoneNumberExists mocks base method.
func (m *MockRepositoryInterface) IsPhoneNumberExists(ctx context.Context, phoneNumber string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPhoneNumberExists", ctx, phoneNumber)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPhoneNumberExists indicates an expected call of IsPhoneNumberExists.
func (mr *MockRepositoryInterfaceMockRecorder) IsPhoneNumberExists(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPhoneNumberExists", reflect.TypeOf((*MockRepositoryInterface)(nil).IsPhoneNumberExists), ctx, phoneNumber)
}
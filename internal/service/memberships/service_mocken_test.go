// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mocken_test.go -package=memberships
//

// Package memberships is a generated GoMock package.
package memberships

import (
	reflect "reflect"

	memberships "github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	gomock "go.uber.org/mock/gomock"
)

// MockrepositoryInterface is a mock of repositoryInterface interface.
type MockrepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockrepositoryInterfaceMockRecorder is the mock recorder for MockrepositoryInterface.
type MockrepositoryInterfaceMockRecorder struct {
	mock *MockrepositoryInterface
}

// NewMockrepositoryInterface creates a new mock instance.
func NewMockrepositoryInterface(ctrl *gomock.Controller) *MockrepositoryInterface {
	mock := &MockrepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockrepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockrepositoryInterface) EXPECT() *MockrepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockrepositoryInterface) CreateUser(model memberships.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockrepositoryInterfaceMockRecorder) CreateUser(model any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockrepositoryInterface)(nil).CreateUser), model)
}

// GetUser mocks base method.
func (m *MockrepositoryInterface) GetUser(email, username string, id uint) (*memberships.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", email, username, id)
	ret0, _ := ret[0].(*memberships.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockrepositoryInterfaceMockRecorder) GetUser(email, username, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockrepositoryInterface)(nil).GetUser), email, username, id)
}

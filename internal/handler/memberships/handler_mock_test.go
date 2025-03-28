// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go
//
// Generated by this command:
//
//	mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
//

// Package memberships is a generated GoMock package.
package memberships

import (
	reflect "reflect"

	memberships "github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	gomock "go.uber.org/mock/gomock"
)

// MockserviceInterface is a mock of serviceInterface interface.
type MockserviceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockserviceInterfaceMockRecorder
	isgomock struct{}
}

// MockserviceInterfaceMockRecorder is the mock recorder for MockserviceInterface.
type MockserviceInterfaceMockRecorder struct {
	mock *MockserviceInterface
}

// NewMockserviceInterface creates a new mock instance.
func NewMockserviceInterface(ctrl *gomock.Controller) *MockserviceInterface {
	mock := &MockserviceInterface{ctrl: ctrl}
	mock.recorder = &MockserviceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockserviceInterface) EXPECT() *MockserviceInterfaceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockserviceInterface) Login(request *memberships.LoginRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockserviceInterfaceMockRecorder) Login(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockserviceInterface)(nil).Login), request)
}

// Signup mocks base method.
func (m *MockserviceInterface) Signup(request memberships.SignUpRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signup indicates an expected call of Signup.
func (mr *MockserviceInterfaceMockRecorder) Signup(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockserviceInterface)(nil).Signup), request)
}

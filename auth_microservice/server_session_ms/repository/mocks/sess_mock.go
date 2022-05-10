// Code generated by MockGen. DO NOT EDIT.
// Source: sess.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "PLANEXA_backend/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSessionRedis is a mock of SessionRedis interface.
type MockSessionRedis struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRedisMockRecorder
}

// MockSessionRedisMockRecorder is the mock recorder for MockSessionRedis.
type MockSessionRedisMockRecorder struct {
	mock *MockSessionRedis
}

// NewMockSessionRedis creates a new mock instance.
func NewMockSessionRedis(ctrl *gomock.Controller) *MockSessionRedis {
	mock := &MockSessionRedis{ctrl: ctrl}
	mock.recorder = &MockSessionRedisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionRedis) EXPECT() *MockSessionRedisMockRecorder {
	return m.recorder
}

// DeleteSession mocks base method.
func (m *MockSessionRedis) DeleteSession(cookieValue string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", cookieValue)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockSessionRedisMockRecorder) DeleteSession(cookieValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionRedis)(nil).DeleteSession), cookieValue)
}

// GetSession mocks base method.
func (m *MockSessionRedis) GetSession(cookieValue string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", cookieValue)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockSessionRedisMockRecorder) GetSession(cookieValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockSessionRedis)(nil).GetSession), cookieValue)
}

// SetSession mocks base method.
func (m *MockSessionRedis) SetSession(session models.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSession", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSession indicates an expected call of SetSession.
func (mr *MockSessionRedisMockRecorder) SetSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockSessionRedis)(nil).SetSession), session)
}
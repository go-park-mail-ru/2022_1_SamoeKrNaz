// Code generated by MockGen. DO NOT EDIT.
// Source: list_usecase.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	models "PLANEXA_backend/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockListUseCase is a mock of ListUseCase interface.
type MockListUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockListUseCaseMockRecorder
}

// MockListUseCaseMockRecorder is the mock recorder for MockListUseCase.
type MockListUseCaseMockRecorder struct {
	mock *MockListUseCase
}

// NewMockListUseCase creates a new mock instance.
func NewMockListUseCase(ctrl *gomock.Controller) *MockListUseCase {
	mock := &MockListUseCase{ctrl: ctrl}
	mock.recorder = &MockListUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListUseCase) EXPECT() *MockListUseCaseMockRecorder {
	return m.recorder
}

// CreateList mocks base method.
func (m *MockListUseCase) CreateList(list models.List, userId, boardId uint) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", list, userId, boardId)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockListUseCaseMockRecorder) CreateList(list, userId, boardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockListUseCase)(nil).CreateList), list, userId, boardId)
}

// DeleteList mocks base method.
func (m *MockListUseCase) DeleteList(listId, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteList", listId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteList indicates an expected call of DeleteList.
func (mr *MockListUseCaseMockRecorder) DeleteList(listId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteList", reflect.TypeOf((*MockListUseCase)(nil).DeleteList), listId, userId)
}

// GetLists mocks base method.
func (m *MockListUseCase) GetLists(boardId, userId uint) ([]models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLists", boardId, userId)
	ret0, _ := ret[0].([]models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLists indicates an expected call of GetLists.
func (mr *MockListUseCaseMockRecorder) GetLists(boardId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLists", reflect.TypeOf((*MockListUseCase)(nil).GetLists), boardId, userId)
}

// GetSingleList mocks base method.
func (m *MockListUseCase) GetSingleList(listId, userId uint) (models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSingleList", listId, userId)
	ret0, _ := ret[0].(models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSingleList indicates an expected call of GetSingleList.
func (mr *MockListUseCaseMockRecorder) GetSingleList(listId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSingleList", reflect.TypeOf((*MockListUseCase)(nil).GetSingleList), listId, userId)
}

// RefactorList mocks base method.
func (m *MockListUseCase) RefactorList(list models.List, userId, boardId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefactorList", list, userId, boardId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefactorList indicates an expected call of RefactorList.
func (mr *MockListUseCaseMockRecorder) RefactorList(list, userId, boardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefactorList", reflect.TypeOf((*MockListUseCase)(nil).RefactorList), list, userId, boardId)
}
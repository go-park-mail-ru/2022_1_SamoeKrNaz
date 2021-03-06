// Code generated by MockGen. DO NOT EDIT.
// Source: checkList_repo.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	models "PLANEXA_backend/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCheckListRepository is a mock of CheckListRepository interface.
type MockCheckListRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCheckListRepositoryMockRecorder
}

// MockCheckListRepositoryMockRecorder is the mock recorder for MockCheckListRepository.
type MockCheckListRepositoryMockRecorder struct {
	mock *MockCheckListRepository
}

// NewMockCheckListRepository creates a new mock instance.
func NewMockCheckListRepository(ctrl *gomock.Controller) *MockCheckListRepository {
	mock := &MockCheckListRepository{ctrl: ctrl}
	mock.recorder = &MockCheckListRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckListRepository) EXPECT() *MockCheckListRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCheckListRepository) Create(checkList *models.CheckList) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", checkList)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCheckListRepositoryMockRecorder) Create(checkList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCheckListRepository)(nil).Create), checkList)
}

// Delete mocks base method.
func (m *MockCheckListRepository) Delete(IdCl uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", IdCl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCheckListRepositoryMockRecorder) Delete(IdCl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCheckListRepository)(nil).Delete), IdCl)
}

// GetById mocks base method.
func (m *MockCheckListRepository) GetById(IdCl uint) (*models.CheckList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", IdCl)
	ret0, _ := ret[0].(*models.CheckList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCheckListRepositoryMockRecorder) GetById(IdCl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCheckListRepository)(nil).GetById), IdCl)
}

// GetCheckListItems mocks base method.
func (m *MockCheckListRepository) GetCheckListItems(IdCl uint) (*[]models.CheckListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckListItems", IdCl)
	ret0, _ := ret[0].(*[]models.CheckListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckListItems indicates an expected call of GetCheckListItems.
func (mr *MockCheckListRepositoryMockRecorder) GetCheckListItems(IdCl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckListItems", reflect.TypeOf((*MockCheckListRepository)(nil).GetCheckListItems), IdCl)
}

// Update mocks base method.
func (m *MockCheckListRepository) Update(checkList models.CheckList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", checkList)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCheckListRepositoryMockRecorder) Update(checkList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCheckListRepository)(nil).Update), checkList)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: checklistitem_repo.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	models "PLANEXA_backend/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCheckListItemRepository is a mock of CheckListItemRepository interface.
type MockCheckListItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCheckListItemRepositoryMockRecorder
}

// MockCheckListItemRepositoryMockRecorder is the mock recorder for MockCheckListItemRepository.
type MockCheckListItemRepositoryMockRecorder struct {
	mock *MockCheckListItemRepository
}

// NewMockCheckListItemRepository creates a new mock instance.
func NewMockCheckListItemRepository(ctrl *gomock.Controller) *MockCheckListItemRepository {
	mock := &MockCheckListItemRepository{ctrl: ctrl}
	mock.recorder = &MockCheckListItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckListItemRepository) EXPECT() *MockCheckListItemRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCheckListItemRepository) Create(checkListItem *models.CheckListItem) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", checkListItem)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCheckListItemRepositoryMockRecorder) Create(checkListItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCheckListItemRepository)(nil).Create), checkListItem)
}

// Delete mocks base method.
func (m *MockCheckListItemRepository) Delete(IdClIt uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", IdClIt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCheckListItemRepositoryMockRecorder) Delete(IdClIt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCheckListItemRepository)(nil).Delete), IdClIt)
}

// GetById mocks base method.
func (m *MockCheckListItemRepository) GetById(IdCl uint) (*models.CheckListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", IdCl)
	ret0, _ := ret[0].(*models.CheckListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCheckListItemRepositoryMockRecorder) GetById(IdCl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCheckListItemRepository)(nil).GetById), IdCl)
}

// Update mocks base method.
func (m *MockCheckListItemRepository) Update(checkListItem models.CheckListItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", checkListItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCheckListItemRepositoryMockRecorder) Update(checkListItem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCheckListItemRepository)(nil).Update), checkListItem)
}

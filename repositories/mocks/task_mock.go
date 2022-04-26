// Code generated by MockGen. DO NOT EDIT.
// Source: task_repo.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	models "PLANEXA_backend/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskRepository is a mock of TaskRepository interface.
type MockTaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryMockRecorder
}

// MockTaskRepositoryMockRecorder is the mock recorder for MockTaskRepository.
type MockTaskRepositoryMockRecorder struct {
	mock *MockTaskRepository
}

// NewMockTaskRepository creates a new mock instance.
func NewMockTaskRepository(ctrl *gomock.Controller) *MockTaskRepository {
	mock := &MockTaskRepository{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRepository) EXPECT() *MockTaskRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTaskRepository) Create(task *models.Task, IdL, IdB uint) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", task, IdL, IdB)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTaskRepositoryMockRecorder) Create(task, IdL, IdB interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskRepository)(nil).Create), task, IdL, IdB)
}

// Delete mocks base method.
func (m *MockTaskRepository) Delete(IdT uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", IdT)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTaskRepositoryMockRecorder) Delete(IdT interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTaskRepository)(nil).Delete), IdT)
}

// GetById mocks base method.
func (m *MockTaskRepository) GetById(IdT uint) (*models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", IdT)
	ret0, _ := ret[0].(*models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockTaskRepositoryMockRecorder) GetById(IdT interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTaskRepository)(nil).GetById), IdT)
}

// GetCheckLists mocks base method.
func (m *MockTaskRepository) GetCheckLists(IdT uint) (*[]models.CheckList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckLists", IdT)
	ret0, _ := ret[0].(*[]models.CheckList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckLists indicates an expected call of GetCheckLists.
func (mr *MockTaskRepositoryMockRecorder) GetCheckLists(IdT interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckLists", reflect.TypeOf((*MockTaskRepository)(nil).GetCheckLists), IdT)
}

// GetImportantTasks mocks base method.
func (m *MockTaskRepository) GetImportantTasks(IdU uint) (*[]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImportantTasks", IdU)
	ret0, _ := ret[0].(*[]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImportantTasks indicates an expected call of GetImportantTasks.
func (mr *MockTaskRepositoryMockRecorder) GetImportantTasks(IdU interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImportantTasks", reflect.TypeOf((*MockTaskRepository)(nil).GetImportantTasks), IdU)
}

// GetTasks mocks base method.
func (m *MockTaskRepository) GetTasks(IdL uint) (*[]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks", IdL)
	ret0, _ := ret[0].(*[]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockTaskRepositoryMockRecorder) GetTasks(IdL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockTaskRepository)(nil).GetTasks), IdL)
}

// IsAccessToTask mocks base method.
func (m *MockTaskRepository) IsAccessToTask(IdU, IdT uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAccessToTask", IdU, IdT)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAccessToTask indicates an expected call of IsAccessToTask.
func (mr *MockTaskRepositoryMockRecorder) IsAccessToTask(IdU, IdT interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAccessToTask", reflect.TypeOf((*MockTaskRepository)(nil).IsAccessToTask), IdU, IdT)
}

// Update mocks base method.
func (m *MockTaskRepository) Update(task models.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTaskRepositoryMockRecorder) Update(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskRepository)(nil).Update), task)
}

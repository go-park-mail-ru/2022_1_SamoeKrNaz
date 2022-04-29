// Code generated by MockGen. DO NOT EDIT.
// Source: board_repo.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	models "PLANEXA_backend/models"
	multipart "mime/multipart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBoardRepository is a mock of BoardRepository interface.
type MockBoardRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBoardRepositoryMockRecorder
}

// MockBoardRepositoryMockRecorder is the mock recorder for MockBoardRepository.
type MockBoardRepositoryMockRecorder struct {
	mock *MockBoardRepository
}

// NewMockBoardRepository creates a new mock instance.
func NewMockBoardRepository(ctrl *gomock.Controller) *MockBoardRepository {
	mock := &MockBoardRepository{ctrl: ctrl}
	mock.recorder = &MockBoardRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBoardRepository) EXPECT() *MockBoardRepositoryMockRecorder {
	return m.recorder
}

// AppendUser mocks base method.
func (m *MockBoardRepository) AppendUser(boardId, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendUser", boardId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendUser indicates an expected call of AppendUser.
func (mr *MockBoardRepositoryMockRecorder) AppendUser(boardId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendUser", reflect.TypeOf((*MockBoardRepository)(nil).AppendUser), boardId, userId)
}

// Create mocks base method.
func (m *MockBoardRepository) Create(board *models.Board) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", board)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBoardRepositoryMockRecorder) Create(board interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBoardRepository)(nil).Create), board)
}

// Delete mocks base method.
func (m *MockBoardRepository) Delete(IdB uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", IdB)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBoardRepositoryMockRecorder) Delete(IdB interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBoardRepository)(nil).Delete), IdB)
}

// DeleteUser mocks base method.
func (m *MockBoardRepository) DeleteUser(boardId, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", boardId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockBoardRepositoryMockRecorder) DeleteUser(boardId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockBoardRepository)(nil).DeleteUser), boardId, userId)
}

// GetBoardUser mocks base method.
func (m *MockBoardRepository) GetBoardUser(IdB uint) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoardUser", IdB)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoardUser indicates an expected call of GetBoardUser.
func (mr *MockBoardRepositoryMockRecorder) GetBoardUser(IdB interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoardUser", reflect.TypeOf((*MockBoardRepository)(nil).GetBoardUser), IdB)
}

// GetById mocks base method.
func (m *MockBoardRepository) GetById(IdB uint) (*models.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", IdB)
	ret0, _ := ret[0].(*models.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBoardRepositoryMockRecorder) GetById(IdB interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBoardRepository)(nil).GetById), IdB)
}

// GetLists mocks base method.
func (m *MockBoardRepository) GetLists(IdB uint) ([]models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLists", IdB)
	ret0, _ := ret[0].([]models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLists indicates an expected call of GetLists.
func (mr *MockBoardRepositoryMockRecorder) GetLists(IdB interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLists", reflect.TypeOf((*MockBoardRepository)(nil).GetLists), IdB)
}

// GetUserBoards mocks base method.
func (m *MockBoardRepository) GetUserBoards(IdU uint) ([]models.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBoards", IdU)
	ret0, _ := ret[0].([]models.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBoards indicates an expected call of GetUserBoards.
func (mr *MockBoardRepositoryMockRecorder) GetUserBoards(IdU interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBoards", reflect.TypeOf((*MockBoardRepository)(nil).GetUserBoards), IdU)
}

// IsAccessToBoard mocks base method.
func (m *MockBoardRepository) IsAccessToBoard(IdU, IdB uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAccessToBoard", IdU, IdB)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAccessToBoard indicates an expected call of IsAccessToBoard.
func (mr *MockBoardRepositoryMockRecorder) IsAccessToBoard(IdU, IdB interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAccessToBoard", reflect.TypeOf((*MockBoardRepository)(nil).IsAccessToBoard), IdU, IdB)
}

// SaveImage mocks base method.
func (m *MockBoardRepository) SaveImage(board *models.Board, header *multipart.FileHeader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveImage", board, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveImage indicates an expected call of SaveImage.
func (mr *MockBoardRepositoryMockRecorder) SaveImage(board, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveImage", reflect.TypeOf((*MockBoardRepository)(nil).SaveImage), board, header)
}

// Update mocks base method.
func (m *MockBoardRepository) Update(board models.Board) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", board)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBoardRepositoryMockRecorder) Update(board interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBoardRepository)(nil).Update), board)
}

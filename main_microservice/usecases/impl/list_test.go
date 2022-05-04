package impl

import (
	customErrors "PLANEXA_backend/errors"
	mock_repositories2 "PLANEXA_backend/main_microservice/repositories/mocks"
	"PLANEXA_backend/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLists(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories2.NewMockBoardRepository(controller)
	listRepo := mock_repositories2.NewMockListRepository(controller)
	listUseCase := MakeListUsecase(listRepo, boardRepo)

	lists := []models.List{{Title: "title1", Position: 1}, {Title: "title2", Position: 2}}

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(11)).Return(true, nil)
	boardRepo.EXPECT().GetLists(uint(11)).Return(lists, nil)
	newLists, err := listUseCase.GetLists(uint(11), uint(22))
	assert.Equal(t, lists, newLists)
	assert.Equal(t, err, nil)

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(11)).Return(false, customErrors.ErrNoAccess)
	newLists, err = listUseCase.GetLists(uint(11), uint(22))
	assert.Equal(t, []models.List(nil), newLists)
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestGetSingleList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories2.NewMockBoardRepository(controller)
	listRepo := mock_repositories2.NewMockListRepository(controller)
	listUseCase := MakeListUsecase(listRepo, boardRepo)

	list := models.List{Title: "title2", Position: 2}
	board := models.Board{IdB: 1, Title: "title1", Description: "desc1"}

	listRepo.EXPECT().GetBoard(uint(11)).Return(&board, nil)
	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(1)).Return(true, nil)
	listRepo.EXPECT().GetById(uint(11)).Return(&list, nil)
	newList, err := listUseCase.GetSingleList(uint(11), uint(22))
	assert.Equal(t, list, newList)
	assert.Equal(t, err, nil)

	listRepo.EXPECT().GetBoard(uint(11)).Return(nil, customErrors.ErrNoAccess)
	newList, err = listUseCase.GetSingleList(uint(11), uint(22))
	assert.Equal(t, models.List{}, newList)
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestCreateList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories2.NewMockBoardRepository(controller)
	listRepo := mock_repositories2.NewMockListRepository(controller)
	listUseCase := MakeListUsecase(listRepo, boardRepo)

	list := models.List{Title: "title2", Position: 2}

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(1)).Return(true, nil)
	listRepo.EXPECT().Create(&list, uint(1)).Return(uint(10), nil)
	listRepo.EXPECT().GetById(uint(10)).Return(&list, nil)
	newList, err := listUseCase.CreateList(list, uint(1), uint(22))
	assert.Equal(t, &list, newList)
	assert.Equal(t, err, nil)

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(1)).Return(false, customErrors.ErrNoAccess)
	newList, err = listUseCase.CreateList(list, uint(1), uint(22))
	assert.Equal(t, (*models.List)(nil), newList)
	assert.Equal(t, err, customErrors.ErrNoAccess)

}

func TestRefactorList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories2.NewMockBoardRepository(controller)
	listRepo := mock_repositories2.NewMockListRepository(controller)
	listUseCase := MakeListUsecase(listRepo, boardRepo)

	list := models.List{Title: "title2", Position: 2}

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(1)).Return(true, nil)
	listRepo.EXPECT().Update(list).Return(nil)
	err := listUseCase.RefactorList(list, uint(22), uint(1))
	assert.Equal(t, err, nil)

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(1)).Return(false, customErrors.ErrNoAccess)
	err = listUseCase.RefactorList(list, uint(22), uint(1))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestDeleteList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories2.NewMockBoardRepository(controller)
	listRepo := mock_repositories2.NewMockListRepository(controller)
	listUseCase := MakeListUsecase(listRepo, boardRepo)

	list := models.List{Title: "title2", Position: 2}
	board := models.Board{IdB: 1, Title: "title1", Description: "desc1"}

	listRepo.EXPECT().GetBoard(uint(11)).Return(&board, nil)
	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(1)).Return(true, nil)
	listRepo.EXPECT().GetById(uint(11)).Return(&list, nil)
	listRepo.EXPECT().Delete(uint(11)).Return(nil)
	err := listUseCase.DeleteList(uint(11), uint(22))
	assert.Equal(t, nil, err)

	listRepo.EXPECT().GetBoard(uint(11)).Return(nil, customErrors.ErrNoAccess)
	err = listUseCase.DeleteList(uint(11), uint(22))
	assert.Equal(t, customErrors.ErrNoAccess, err)
}

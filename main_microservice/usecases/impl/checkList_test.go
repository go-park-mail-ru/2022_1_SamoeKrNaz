package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories/mocks"
	"PLANEXA_backend/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCheckLists(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListUseCase := MakeCheckListUsecase(checkListRepo, taskRepo)

	checkLists := []models.CheckList{{Title: "title1"}, {Title: "title2"}}
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(22)).Return(true, nil)
	taskRepo.EXPECT().GetCheckLists(uint(22)).Return(&checkLists, nil)
	newCheckLists, err := checkListUseCase.GetCheckLists(uint(11), uint(22))
	assert.Equal(t, &checkLists, newCheckLists)
	assert.Equal(t, err, nil)

	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(22)).Return(false, customErrors.ErrBadInputData)
	newCheckLists, err = checkListUseCase.GetCheckLists(uint(11), uint(22))
	assert.Equal(t, (*[]models.CheckList)(nil), newCheckLists)
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestGetSingleCheckList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListUseCase := MakeCheckListUsecase(checkListRepo, taskRepo)

	checkLists := models.CheckList{Title: "title1", IdCl: 0}
	checkListRepo.EXPECT().GetById(uint(11)).Return(&checkLists, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(22), checkLists.IdCl).Return(true, nil)
	newCheckList, err := checkListUseCase.GetSingleCheckList(uint(22), uint(11))
	assert.Equal(t, &checkLists, newCheckList)
	assert.Equal(t, err, nil)

	checkListRepo.EXPECT().GetById(uint(11)).Return(&checkLists, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(22), checkLists.IdCl).Return(false, customErrors.ErrBadInputData)
	newCheckList, err = checkListUseCase.GetSingleCheckList(uint(22), uint(11))
	assert.Equal(t, (*models.CheckList)(nil), newCheckList)
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestCreateCheckList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListUseCase := MakeCheckListUsecase(checkListRepo, taskRepo)

	checkLists := models.CheckList{Title: "title1", IdCl: 0}
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(1)).Return(true, nil)
	checkListRepo.EXPECT().Create(&checkLists).Return(uint(0), nil)
	checkListRepo.EXPECT().GetById(uint(0)).Return(&checkLists, nil)
	newCheckList, err := checkListUseCase.CreateCheckList(&checkLists, uint(1), uint(11))
	assert.Equal(t, &checkLists, newCheckList)
	assert.Equal(t, err, nil)

	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(1)).Return(false, customErrors.ErrBadInputData)
	newCheckList, err = checkListUseCase.CreateCheckList(&checkLists, uint(1), uint(11))
	assert.Equal(t, (*models.CheckList)(nil), newCheckList)
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestRefactorCheckList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListUseCase := MakeCheckListUsecase(checkListRepo, taskRepo)

	checkLists := models.CheckList{Title: "title1", IdCl: 0, IdT: 0}

	checkListRepo.EXPECT().GetById(uint(0)).Return(&checkLists, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(true, nil)
	checkListRepo.EXPECT().Update(checkLists).Return(nil)
	err := checkListUseCase.RefactorCheckList(&checkLists, uint(11))
	assert.Equal(t, err, nil)

	checkListRepo.EXPECT().GetById(uint(0)).Return(&checkLists, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(false, customErrors.ErrNoAccess)
	err = checkListUseCase.RefactorCheckList(&checkLists, uint(11))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestDeleteCheckList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListUseCase := MakeCheckListUsecase(checkListRepo, taskRepo)

	checkLists := models.CheckList{Title: "title1", IdCl: 1, IdT: 0}

	checkListRepo.EXPECT().GetById(uint(1)).Return(&checkLists, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(true, nil)
	checkListRepo.EXPECT().Delete(uint(1)).Return(nil)
	err := checkListUseCase.DeleteCheckList(1, uint(11))
	assert.Equal(t, err, nil)

	checkListRepo.EXPECT().GetById(uint(1)).Return(&checkLists, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(false, customErrors.ErrNoAccess)
	err = checkListUseCase.DeleteCheckList(1, uint(11))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCheckListItems(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListItemRepo := mock_repositories.NewMockCheckListItemRepository(controller)
	checkListUseCase := MakeCheckListItemUsecase(checkListItemRepo, checkListRepo, taskRepo)

	checkListItems := []models.CheckListItem{{Description: "title1"}, {Description: "title2"}}
	checkListRepo.EXPECT().GetCheckListItems(uint(22)).Return(&checkListItems, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(true, nil)
	newCheckListItems, err := checkListUseCase.GetCheckListItems(uint(11), uint(22))
	assert.Equal(t, &checkListItems, newCheckListItems)
	assert.Equal(t, err, nil)

	checkListRepo.EXPECT().GetCheckListItems(uint(22)).Return(&checkListItems, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(false, customErrors.ErrBadInputData)
	newCheckListItems, err = checkListUseCase.GetCheckListItems(uint(11), uint(22))
	assert.Equal(t, (*[]models.CheckListItem)(nil), newCheckListItems)
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestGetSingleCheckListItem(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListItemRepo := mock_repositories.NewMockCheckListItemRepository(controller)
	checkListUseCase := MakeCheckListItemUsecase(checkListItemRepo, checkListRepo, taskRepo)

	checkListItems := models.CheckListItem{Description: "title1"}
	checkListItemRepo.EXPECT().GetById(uint(22)).Return(&checkListItems, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(true, nil)
	newCheckListItem, err := checkListUseCase.GetSingleCheckListItem(uint(22), uint(11))
	assert.Equal(t, &checkListItems, newCheckListItem)
	assert.Equal(t, err, nil)

	checkListItemRepo.EXPECT().GetById(uint(22)).Return(&checkListItems, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(false, customErrors.ErrBadInputData)
	newCheckListItem, err = checkListUseCase.GetSingleCheckListItem(uint(22), uint(11))
	assert.Equal(t, (*models.CheckListItem)(nil), newCheckListItem)
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestCreateCheckListItem(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListItemRepo := mock_repositories.NewMockCheckListItemRepository(controller)
	checkListUseCase := MakeCheckListItemUsecase(checkListItemRepo, checkListRepo, taskRepo)

	checkListItems := models.CheckListItem{Description: "title1"}
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(true, nil)
	checkListItemRepo.EXPECT().Create(&checkListItems).Return(uint(0), nil)
	checkListItemRepo.EXPECT().GetById(uint(0)).Return(&checkListItems, nil)
	newCheckListItem, err := checkListUseCase.CreateCheckListItem(&checkListItems, uint(0), uint(11))
	assert.Equal(t, &checkListItems, newCheckListItem)
	assert.Equal(t, err, nil)

	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(false, customErrors.ErrBadInputData)
	newCheckListItem, err = checkListUseCase.CreateCheckListItem(&checkListItems, uint(1), uint(11))
	assert.Equal(t, (*models.CheckListItem)(nil), newCheckListItem)
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestRefactorCheckListItem(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListItemRepo := mock_repositories.NewMockCheckListItemRepository(controller)
	checkListUseCase := MakeCheckListItemUsecase(checkListItemRepo, checkListRepo, taskRepo)

	checkListItems := models.CheckListItem{Description: "title1"}
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(true, nil)
	checkListItemRepo.EXPECT().Update(checkListItems).Return(nil)
	err := checkListUseCase.RefactorCheckListItem(&checkListItems, uint(11))
	assert.Equal(t, err, nil)

	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(false, customErrors.ErrNoAccess)
	err = checkListUseCase.RefactorCheckListItem(&checkListItems, uint(11))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestDeleteCheckListItem(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	checkListItemRepo := mock_repositories.NewMockCheckListItemRepository(controller)
	checkListUseCase := MakeCheckListItemUsecase(checkListItemRepo, checkListRepo, taskRepo)

	checkListItems := models.CheckListItem{Description: "title1"}

	checkListItemRepo.EXPECT().GetById(uint(1)).Return(&checkListItems, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(true, nil)
	checkListItemRepo.EXPECT().Delete(uint(1)).Return(nil)
	err := checkListUseCase.DeleteCheckListItem(1, uint(11))
	assert.Equal(t, err, nil)

	checkListItemRepo.EXPECT().GetById(uint(1)).Return(&checkListItems, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(false, customErrors.ErrNoAccess)
	err = checkListUseCase.DeleteCheckListItem(1, uint(11))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

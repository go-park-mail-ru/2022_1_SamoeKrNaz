package impl

import (
	customErrors "PLANEXA_backend/errors"
	repositories2 "PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"github.com/microcosm-cc/bluemonday"
)

type CheckListItemUseCaseImpl struct {
	repCheckListItem repositories2.CheckListItemRepository
	repCheckList     repositories2.CheckListRepository
	repTask          repositories2.TaskRepository
}

func MakeCheckListItemUsecase(repCheckListItem_ repositories2.CheckListItemRepository, repCheckList_ repositories2.CheckListRepository,
	repTask_ repositories2.TaskRepository) usecases.CheckListItemUseCase {
	return &CheckListItemUseCaseImpl{repCheckListItem: repCheckListItem_, repCheckList: repCheckList_, repTask: repTask_}
}

func (checkListItemUseCase *CheckListItemUseCaseImpl) GetCheckListItems(userId uint, IdCl uint) (*[]models.CheckListItem, error) {
	checkListItems, err := checkListItemUseCase.repCheckList.GetCheckListItems(IdCl)
	if err != nil {
		return nil, err
	}
	isAccess, err := checkListItemUseCase.repTask.IsAccessToTask(userId, (*checkListItems)[0].IdT)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, checkList := range *checkListItems {
		checkList.Description = sanitizer.Sanitize(checkList.Description)
	}
	return checkListItems, nil
}

func (checkListItemUseCase *CheckListItemUseCaseImpl) GetSingleCheckListItem(IdClIt uint, userId uint) (*models.CheckListItem, error) {
	checkListItem, err := checkListItemUseCase.repCheckListItem.GetById(IdClIt)
	if err != nil {
		return nil, err
	}
	isAccess, err := checkListItemUseCase.repTask.IsAccessToTask(userId, checkListItem.IdT)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	sanitizer := bluemonday.UGCPolicy()
	checkListItem.Description = sanitizer.Sanitize(checkListItem.Description)
	return checkListItem, nil
}

func (checkListItemUseCase *CheckListItemUseCaseImpl) CreateCheckListItem(checkListItem *models.CheckListItem, IdCl uint, userId uint) (*models.CheckListItem, error) {
	checkListItem.IdCl = IdCl
	isAccess, err := checkListItemUseCase.repTask.IsAccessToTask(userId, checkListItem.IdT)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	checkListItemId, err := checkListItemUseCase.repCheckListItem.Create(checkListItem)
	if err != nil {
		return nil, err
	}
	createdCheckListItem, err := checkListItemUseCase.repCheckListItem.GetById(checkListItemId)
	return createdCheckListItem, err
}

func (checkListItemUseCase *CheckListItemUseCaseImpl) RefactorCheckListItem(checkListItem *models.CheckListItem, userId uint) error {
	isAccess, err := checkListItemUseCase.repTask.IsAccessToTask(userId, checkListItem.IdT)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	return checkListItemUseCase.repCheckListItem.Update(*checkListItem)
}

func (checkListItemUseCase *CheckListItemUseCaseImpl) DeleteCheckListItem(IdClIt uint, userId uint) error {
	checkListItem, err := checkListItemUseCase.repCheckListItem.GetById(IdClIt)
	if err != nil {
		return err
	}
	isAccess, err := checkListItemUseCase.repTask.IsAccessToTask(userId, checkListItem.IdT)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	return checkListItemUseCase.repCheckListItem.Delete(IdClIt)
}

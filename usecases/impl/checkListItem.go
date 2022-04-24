package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/usecases"
	"github.com/microcosm-cc/bluemonday"
)

type CheckListItemUseCaseImpl struct {
	repCheckListItem repositories.CheckListItemRepository
	repTask          repositories.TaskRepository
}

func MakeCheckListItemUsecase(repCheckListItem_ repositories.CheckListItemRepository, repTask_ repositories.TaskRepository) usecases.CheckListItemUseCase {
	return &CheckListItemUseCaseImpl{repCheckListItem: repCheckListItem_, repTask: repTask_}
}

func (checkListItemUseCase *CheckListItemUseCaseImpl) GetCheckLists(userId uint, IdCl uint) (*[]models.CheckListItem, error) {
	checkListItems, err := checkListItemUseCase.repCheckListItem.GetCheckListItems(IdCl)
	if err != nil {
		return nil, err
	}
	isAccess, err := checkListItemUseCase.repTask.IsAccessToTask(userId, checkListItems[0].IdT)
	if err != nil {
		return nil, err
	} else if isAccess == false {
		return nil, customErrors.ErrNoAccess
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, checkList := range checkListItems {
		checkList.Description = sanitizer.Sanitize(checkList.Description)
	}
	return &checkListItems, nil
}

func (checkListItemUseCase *CheckListItemUseCaseImpl) GetSingleCheckListItem(IdClIt uint, userId uint) (*models.CheckListItem, error) {
	checkListItem, err := checkListItemUseCase.repCheckListItem.GetById(IdClIt)
	isAccess, err := checkListItemUseCase.repTask.IsAccessToTask(userId, checkListItem.IdT)
	if err != nil {
		return nil, err
	} else if isAccess == false {
		return nil, customErrors.ErrNoAccess
	}
	if err != nil {
		return nil, err
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
	} else if isAccess == false {
		return nil, customErrors.ErrNoAccess
	}
	checkListItemId, err := checkListItemUseCase.repCheckListItem.Create(checkListItem)
	if err != nil {
		return nil, err
	}
	createdBoard, err := checkListItemUseCase.repCheckListItem.GetById(checkListItemId)
	return createdBoard, nil
}

func (checkListItemUseCase *CheckListItemUseCaseImpl) RefactorCheckListItem(checkListItem *models.CheckListItem, userId uint) error {
	isAccess, err := checkListItemUseCase.repTask.IsAccessToTask(userId, checkListItem.IdT)
	if err != nil {
		return err
	} else if isAccess == false {
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
	} else if isAccess == false {
		return customErrors.ErrNoAccess
	}
	return checkListItemUseCase.repCheckListItem.Delete(IdClIt)
}

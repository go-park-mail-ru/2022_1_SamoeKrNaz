package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/usecases"
	"github.com/microcosm-cc/bluemonday"
)

type CheckListUseCaseImpl struct {
	repCheckList repositories.CheckListRepository
	repTask      repositories.TaskRepository
}

func MakeCheckListUsecase(repCheckList_ repositories.CheckListRepository, repTask_ repositories.TaskRepository) usecases.CheckListUseCase {
	return &CheckListUseCaseImpl{repCheckList: repCheckList_, repTask: repTask_}
}

func (checkListUseCase *CheckListUseCaseImpl) GetSingleCheckList(userId uint, IdCl uint) (*models.CheckList, error) {
	checkList, err := checkListUseCase.repCheckList.GetById(IdCl)
	if err != nil {
		return nil, err
	}
	isAccess, err := checkListUseCase.repTask.IsAccessToTask(userId, checkList.IdT)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	sanitizer := bluemonday.UGCPolicy()
	checkList.Title = sanitizer.Sanitize(checkList.Title)
	return checkList, nil
}

func (checkListUseCase *CheckListUseCaseImpl) GetCheckLists(userId uint, IdT uint) (*[]models.CheckList, error) {
	isAccess, err := checkListUseCase.repTask.IsAccessToTask(userId, IdT)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	checkLists, err := checkListUseCase.repTask.GetCheckLists(IdT)
	if err != nil {
		return nil, err
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, checkList := range *checkLists {
		checkList.Title = sanitizer.Sanitize(checkList.Title)
	}
	return checkLists, nil
}

func (checkListUseCase *CheckListUseCaseImpl) CreateCheckList(checkList *models.CheckList, IdT uint, userId uint) (*models.CheckList, error) {
	checkList.IdT = IdT
	isAccess, err := checkListUseCase.repTask.IsAccessToTask(userId, checkList.IdT)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	checkListId, err := checkListUseCase.repCheckList.Create(checkList)
	if err != nil {
		return nil, err
	}
	createdCheckList, err := checkListUseCase.repCheckList.GetById(checkListId)
	return createdCheckList, err
}

func (checkListUseCase *CheckListUseCaseImpl) RefactorCheckList(checkList *models.CheckList, userId uint) error {
	currentData, err := checkListUseCase.repCheckList.GetById(checkList.IdCl)
	if err != nil {
		return err
	}
	isAccess, err := checkListUseCase.repTask.IsAccessToTask(userId, currentData.IdT)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	return checkListUseCase.repCheckList.Update(*checkList)
}

func (checkListUseCase *CheckListUseCaseImpl) DeleteCheckList(IdCl uint, userId uint) error {
	checkList, err := checkListUseCase.repCheckList.GetById(IdCl)
	if err != nil {
		return err
	}
	isAccess, err := checkListUseCase.repTask.IsAccessToTask(userId, checkList.IdT)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	err = checkListUseCase.repCheckList.Delete(IdCl)
	return err
}

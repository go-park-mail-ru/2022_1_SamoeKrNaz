package usecases

import "PLANEXA_backend/models"

type CheckListUseCase interface {
	GetCheckList(IdCl uint, userId uint) (*models.CheckList, error)
	CreateCheckList(checkList *models.CheckList, IdT uint, userId uint) (*models.CheckList, error)
	RefactorCheckList(checkList *models.CheckList, userId uint) error
	DeleteCheckList(IdCl uint, userId uint) error
}

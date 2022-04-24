package usecases

import "PLANEXA_backend/models"

type CheckListItemUseCase interface {
	GetCheckLists(userId uint, IdCl uint) (*[]models.CheckListItem, error)
	GetSingleCheckListItem(IdClIt uint, userId uint) (*models.CheckListItem, error)
	CreateCheckListItem(checkListItem *models.CheckListItem, IdCl uint, userId uint) (*models.CheckListItem, error)
	RefactorCheckListItem(checkListItem *models.CheckListItem, userId uint) error
	DeleteCheckListItem(IdClIt uint, userId uint) error
}

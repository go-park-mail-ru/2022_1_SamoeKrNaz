package repositories

import "PLANEXA_backend/models"

type CheckListRepository interface {
	Create(checkList *models.CheckList) (uint, error)
	GetById(IdCl uint) (*models.CheckList, error)
	Update(checkList models.CheckList) error
	Delete(IdCl uint) error
	GetCheckListItems(IdCl uint) (*[]models.CheckListItem, error)
	GetCheckLists(IdT uint) ([]models.CheckList, error)
}

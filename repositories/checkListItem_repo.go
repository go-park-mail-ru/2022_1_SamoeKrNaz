package repositories

import "PLANEXA_backend/models"

type CheckListItemRepository interface {
	Create(checkListItem *models.CheckListItem) (uint, error)
	GetById(IdCl uint) (*models.CheckListItem, error)
	Update(checkListItem models.CheckListItem) error
	Delete(IdClIt uint) error
	GetCheckListItems(IdCl uint) ([]models.CheckListItem, error)
}

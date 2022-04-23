package repositories

import "PLANEXA_backend/models"

type CheckListRepository interface {
	Create(checkList *models.CheckList) (uint, error)
	GetById(IdCl uint) (*models.CheckList, error)
	Update(checkList models.CheckList) error
	Delete(IdCl uint) error
}

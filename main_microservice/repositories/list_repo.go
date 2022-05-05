package repositories

import "PLANEXA_backend/models"

type ListRepository interface {
	Create(list *models.List, IdB uint) (uint, error)
	Update(list models.List) error
	Delete(IdL uint) error
	GetTasks(IdL uint) (*[]models.Task, error)
	GetById(IdL uint) (*models.List, error)
	GetBoard(IdL uint) (*models.Board, error)
}

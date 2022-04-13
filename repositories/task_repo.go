package repositories

import "PLANEXA_backend/models"

type TaskRepository interface {
	Create(task *models.Task, IdL uint, IdB uint) (uint, error)
	GetTasks(IdL uint) (*[]models.Task, error)
	Update(task models.Task) error
	Delete(IdT uint) error
	GetById(IdT uint) (*models.Task, error)
}

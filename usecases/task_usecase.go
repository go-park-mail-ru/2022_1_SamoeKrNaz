package usecases

import "PLANEXA_backend/models"

type TaskUseCase interface {
	GetTasks(listId uint, userId uint) ([]models.Task, error)
	GetSingleTask(taskId uint, userId uint) (models.Task, error)
	CreateTask(task models.Task, idB uint, idL uint, idU uint) (*models.Task, error)
	RefactorTask(task models.Task, userId uint) error
	DeleteTask(taskId uint, userId uint) error
}

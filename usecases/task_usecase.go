package usecases

import "PLANEXA_backend/models"

type TaskUseCase interface {
	GetTasks(listId uint, userId uint) ([]models.List, error)
	GetSingleTask(listId uint, userId uint) (models.List, error)
	CreateTask(list models.Task, userId uint) (uint, error)
	RefactorTask(task models.Task, userId uint, listId uint) error
	DeleteTask(taskId uint, userId uint) error
}

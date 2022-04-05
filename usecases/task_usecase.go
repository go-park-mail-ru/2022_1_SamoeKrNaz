package usecases

import "PLANEXA_backend/models"

type TaskUseCase interface {
	GetTasks(listId uint, userId uint) ([]models.Task, error)
	GetSingleTask(listId uint, userId uint) (models.Task, error)
	CreateTask(list models.Task, idB uint, idL uint) (uint, error)
	RefactorTask(task models.Task, userId uint, taskId uint) error
	DeleteTask(taskId uint, userId uint) error
}

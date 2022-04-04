package impl

import (
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
)

type TaskUseCaseImpl struct {
	rep *repositories.TaskRepository
}

func MakeTaskUsecase(rep_ *repositories.TaskRepository) *TaskUsecase {
	return &TaskUsecase{rep: rep_}
}

func GetTasks(listId uint, userId uint) ([]models.List, error) {
	// достаю все таски из БД по айди доски
	var err error
	return []models.List{}, err
}

func GetSingleTask(listId uint, userId uint) (models.List, error) {
	// доставю таск из бд
	var err error
	return models.List{}, err
}

func CreateTask(list models.Task, userId uint) (uint, error) {
	// создаю таск в бд, получаю айди таска
	var err error
	return 0, err
}

func RefactorTask(task models.Task, userId uint, listId uint) error {
	// проверяю может ли юзер редачить
	// вношу изменения в бд
	var err error
	return err
}

func DeleteTask(taskId uint, userId uint) error {
	// проверяю есть ли такой таск и может ли юзер удалить его
	// удаляю таск
	var err error
	return err
}

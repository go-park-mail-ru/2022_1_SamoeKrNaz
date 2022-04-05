package impl

import (
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/usecases"
)

type TaskUseCaseImpl struct {
	rep *repositories.TaskRepository
}

func MakeTaskUsecase(rep_ *repositories.TaskRepository) usecases.TaskUseCase {
	return &TaskUseCaseImpl{rep: rep_}
}

func (taskUseCase *TaskUseCaseImpl) GetTasks(listId uint, userId uint) ([]models.Task, error) {
	// достаю все таски из БД по айди доски
	tasks, err := taskUseCase.rep.GetTasks(listId)
	return *tasks, err
}

func (taskUseCase *TaskUseCaseImpl) GetSingleTask(taskId uint, userId uint) (models.Task, error) {
	// доставю таск из бд
	task, err := taskUseCase.rep.GetById(taskId)
	return *task, err
}

func (taskUseCase *TaskUseCaseImpl) CreateTask(task models.Task, idB uint, idL uint) (uint, error) {
	// создаю таск в бд, получаю айди таска
	err := taskUseCase.rep.Create(task, idL, idB)
	return 0, err
}

func (taskUseCase *TaskUseCaseImpl) RefactorTask(task models.Task, userId uint, listId uint) error {
	// проверяю может ли юзер редачить
	// вношу изменения в бд
	err := taskUseCase.rep.Update(task)
	return err
}

func (taskUseCase *TaskUseCaseImpl) DeleteTask(taskId uint, userId uint) error {
	// проверяю есть ли такой таск и может ли юзер удалить его
	// удаляю таск
	_, err := taskUseCase.rep.GetById(taskId)
	if err != nil {
		return err
	}

	err = taskUseCase.rep.Delete(taskId)
	return err
}

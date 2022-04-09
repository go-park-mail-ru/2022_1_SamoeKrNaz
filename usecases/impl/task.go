package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/usecases"
)

type TaskUseCaseImpl struct {
	repTask  *repositories.TaskRepository
	repBoard *repositories.BoardRepository
	repList  *repositories.ListRepository
}

func MakeTaskUsecase(repTask_ *repositories.TaskRepository, repBoard_ *repositories.BoardRepository, repList_ *repositories.ListRepository) usecases.TaskUseCase {
	return &TaskUseCaseImpl{repTask: repTask_, repBoard: repBoard_, repList: repList_}
}

func (taskUseCase *TaskUseCaseImpl) GetTasks(listId uint, userId uint) ([]models.Task, error) {
	// достаю список из бд, чтобы получить айдишник доски
	list, err := taskUseCase.repList.GetById(listId)
	if err != nil {
		return nil, err
	}
	isAccess, err := taskUseCase.repBoard.IsAccessToBoard(userId, list.IdB)
	if err != nil {
		return nil, err
	} else if isAccess == false {
		return nil, customErrors.ErrNoAccess
	}
	tasks, err := taskUseCase.repTask.GetTasks(listId)
	return *tasks, err
}

func (taskUseCase *TaskUseCaseImpl) GetSingleTask(taskId uint, userId uint) (models.Task, error) {
	// доставю таск из бд
	task, err := taskUseCase.repTask.GetById(taskId)
	if err != nil {
		return models.Task{}, err
	}
	isAccess, err := taskUseCase.repBoard.IsAccessToBoard(userId, task.IdB)
	if err != nil {
		return models.Task{}, err
	} else if isAccess == false {
		return models.Task{}, customErrors.ErrNoAccess
	}
	return *task, err
}

func (taskUseCase *TaskUseCaseImpl) CreateTask(task models.Task, idB uint, idL uint, idU uint) (uint, error) {
	isAccess, err := taskUseCase.repBoard.IsAccessToBoard(idU, task.IdB)
	if err != nil {
		return 0, err
	} else if isAccess == false {
		return 0, customErrors.ErrNoAccess
	}
	// создаю таск в бд, получаю айди таска
	taskId, err := taskUseCase.repTask.Create(&task, idL, idB)
	return taskId, err
}

func (taskUseCase *TaskUseCaseImpl) RefactorTask(task models.Task, userId uint, listId uint) error {
	// проверяю может ли юзер редачить
	isAccess, err := taskUseCase.repBoard.IsAccessToBoard(userId, task.IdB)
	if err != nil {
		return err
	} else if isAccess == false {
		return customErrors.ErrNoAccess
	}
	// вношу изменения в бд
	err = taskUseCase.repTask.Update(task)
	return err
}

func (taskUseCase *TaskUseCaseImpl) DeleteTask(taskId uint, userId uint) error {
	task, err := taskUseCase.repTask.GetById(taskId)
	if err != nil {
		return err
	}
	// проверяю есть ли такой таск и может ли юзер удалить его
	isAccess, err := taskUseCase.repBoard.IsAccessToBoard(userId, task.IdB)
	if err != nil {
		return err
	} else if isAccess == false {
		return customErrors.ErrNoAccess
	}
	// удаляю таск
	err = taskUseCase.repTask.Delete(taskId)
	return err
}

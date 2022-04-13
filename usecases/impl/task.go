package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/usecases"
	"github.com/microcosm-cc/bluemonday"
)

type TaskUseCaseImpl struct {
	repTask  repositories.TaskRepository
	repBoard repositories.BoardRepository
	repList  repositories.ListRepository
}

func MakeTaskUsecase(repTask_ repositories.TaskRepository, repBoard_ repositories.BoardRepository, repList_ repositories.ListRepository) usecases.TaskUseCase {
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
	if err != nil {
		return nil, err
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, task := range *tasks {
		task.Title = sanitizer.Sanitize(task.Title)
		task.DateCreated = sanitizer.Sanitize(task.DateCreated)
		task.Description = sanitizer.Sanitize(task.Description)
	}
	return *tasks, err
}

func (taskUseCase *TaskUseCaseImpl) GetSingleTask(taskId uint, userId uint) (models.Task, error) {
	// доставю таск из бд
	task, err := taskUseCase.repTask.GetById(taskId)
	if err != nil {
		return models.Task{}, err
	}
	sanitizer := bluemonday.UGCPolicy()
	task.Title = sanitizer.Sanitize(task.Title)
	task.DateCreated = sanitizer.Sanitize(task.DateCreated)
	task.Description = sanitizer.Sanitize(task.Description)
	isAccess, err := taskUseCase.repBoard.IsAccessToBoard(userId, task.IdB)
	if err != nil {
		return models.Task{}, err
	} else if isAccess == false {
		return models.Task{}, customErrors.ErrNoAccess
	}
	return *task, err
}

func (taskUseCase *TaskUseCaseImpl) CreateTask(task models.Task, idB uint, idL uint, idU uint) (*models.Task, error) {
	isAccess, err := taskUseCase.repBoard.IsAccessToBoard(idU, task.IdB)
	if err != nil {
		return nil, err
	} else if isAccess == false {
		return nil, customErrors.ErrNoAccess
	}
	// создаю таск в бд, получаю айди таска
	taskId, err := taskUseCase.repTask.Create(&task, idL, idB)
	if err != nil {
		return nil, err
	}
	createdTask, err := taskUseCase.repTask.GetById(taskId)
	return createdTask, err
}

func (taskUseCase *TaskUseCaseImpl) RefactorTask(task models.Task, userId uint) error {
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

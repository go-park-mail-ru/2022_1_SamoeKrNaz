package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"time"
)

type TaskUseCaseImpl struct {
	repTask      repositories.TaskRepository
	repBoard     repositories.BoardRepository
	repList      repositories.ListRepository
	repUser      repositories.UserRepository
	repCheckList repositories.CheckListRepository
	repComment   repositories.CommentRepository
}

func MakeTaskUsecase(repTask_ repositories.TaskRepository, repBoard_ repositories.BoardRepository,
	repList_ repositories.ListRepository, repUser_ repositories.UserRepository,
	repCheckList_ repositories.CheckListRepository, repComment_ repositories.CommentRepository) usecases.TaskUseCase {
	return &TaskUseCaseImpl{repTask: repTask_, repBoard: repBoard_,
		repList: repList_, repUser: repUser_, repCheckList: repCheckList_,
		repComment: repComment_}
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
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	tasks, err := taskUseCase.repTask.GetTasks(listId)
	if err != nil {
		return nil, err
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, task := range *tasks {
		task.Title = sanitizer.Sanitize(task.Title)
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
	task.Description = sanitizer.Sanitize(task.Description)
	isAccess, err := taskUseCase.repBoard.IsAccessToBoard(userId, task.IdB)
	if err != nil {
		return models.Task{}, err
	} else if !isAccess {
		return models.Task{}, customErrors.ErrNoAccess
	}
	appendedUsers, err := taskUseCase.repTask.GetTaskUser(taskId)
	if err != nil {
		return models.Task{}, err
	}
	for _, user := range *appendedUsers {
		user.Password = ""
	}
	checkLists, err := taskUseCase.repTask.GetCheckLists(taskId)
	if err != nil {
		return models.Task{}, err
	}
	for i, checkList := range *checkLists {
		checkListItems, err := taskUseCase.repCheckList.GetCheckListItems(checkList.IdCl)
		if err != nil {
			return models.Task{}, err
		}
		(*checkLists)[i].CheckListItems = *checkListItems
	}
	comments, err := taskUseCase.repComment.GetComments(taskId)
	for i, comment := range *comments {
		userComment, err := taskUseCase.repUser.GetUserById(comment.IdU)
		if err != nil {
			return models.Task{}, err
		}
		(*comments)[i].User = *userComment
	}
	if err != nil {
		return models.Task{}, err
	}
	task.Comments = *comments
	task.CheckLists = *checkLists
	task.Users = *appendedUsers
	return *task, err
}

func (taskUseCase *TaskUseCaseImpl) CreateTask(task models.Task, idB uint, idL uint, idU uint) (*models.Task, error) {
	isAccess, err := taskUseCase.repBoard.IsAccessToBoard(idU, task.IdB)
	task.IdU = idU
	task.DateToOrder = time.Now()
	task.Link = uuid.NewString()
	task.IsReady = false
	task.IsImportant = false
	if err != nil {
		return nil, err
	} else if !isAccess {
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
	} else if !isAccess {
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
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	// удаляю таск
	err = taskUseCase.repTask.Delete(taskId)
	return err
}

func (taskUseCase *TaskUseCaseImpl) GetImportantTask(userId uint) (*[]models.Task, error) {
	tasks, err := taskUseCase.repTask.GetImportantTasks(userId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (taskUseCase *TaskUseCaseImpl) AppendUserToTask(userId uint, appendedUserId uint, taskId uint) (models.User, error) {
	isAccess, err := taskUseCase.repTask.IsAccessToTask(userId, taskId)
	if err != nil {
		return models.User{}, err
	} else if !isAccess {
		return models.User{}, customErrors.ErrNoAccess
	}
	err = taskUseCase.repTask.AppendUser(taskId, appendedUserId)
	if err != nil {
		return models.User{}, err
	}
	user, err := taskUseCase.repUser.GetUserById(appendedUserId)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (taskUseCase *TaskUseCaseImpl) DeleteUserFromTask(userId uint, deletedUserId uint, taskId uint) error {
	isAccess, err := taskUseCase.repTask.IsAccessToTask(userId, taskId)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	err = taskUseCase.repTask.DeleteUser(taskId, deletedUserId)
	if err != nil {
		return err
	}
	return nil
}

func (taskUseCase *TaskUseCaseImpl) AppendUserToTaskByLink(userId uint, link string) error {
	task, err := taskUseCase.repTask.GetByLink(link)
	if err != nil {
		return err
	}
	isAccessToBoard, err := taskUseCase.repBoard.IsAccessToBoard(userId, task.IdB)
	if isAccessToBoard {
		return customErrors.ErrAlreadyAppended
	} else if err != nil {
		return err
	}
	err = taskUseCase.repBoard.AppendUser(task.IdB, userId)
	if err != nil {
		return err
	}
	isAccessToTask, err := taskUseCase.repTask.IsAccessToTask(userId, task.IdT)
	if isAccessToTask {
		return customErrors.ErrAlreadyAppended
	} else if err != nil {
		return err
	}
	err = taskUseCase.repTask.AppendUser(task.IdT, userId)
	if err != nil {
		return err
	}
	return nil
}

package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	usecase             usecases.TaskUseCase
	notificationUsecase usecases.NotificationUseCase
}

func MakeTaskHandler(usecase_ usecases.TaskUseCase, notificationUsecase_ usecases.NotificationUseCase) *TaskHandler {
	return &TaskHandler{usecase: usecase_, notificationUsecase: notificationUsecase_}
}

func (taskHandler *TaskHandler) GetTasks(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	listId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	var tasks models.Tasks
	tasks, err = taskHandler.usecase.GetTasks(uint(listId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	tasksJson, err := tasks.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", tasksJson)
}

func (taskHandler *TaskHandler) GetImportantTasks(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	tasks, err := taskHandler.usecase.GetImportantTask(uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newTasks := new(models.Tasks)
	*newTasks = *tasks

	tasksJson, err := newTasks.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", tasksJson)
}

func (taskHandler *TaskHandler) GetSingleTask(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	task, err := taskHandler.usecase.GetSingleTask(uint(taskId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	taskJson, err := task.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("IdB", task.IdB)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", task.IdT)
	c.Data(http.StatusOK, "application/json; charset=utf-8", taskJson)
}

func (taskHandler *TaskHandler) CreateTask(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var task models.Task
	err := easyjson.UnmarshalFromReader(c.Request.Body, &task)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	listId, err := strconv.ParseUint(c.Param("idL"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	boardId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	createdTask, err := taskHandler.usecase.CreateTask(task, uint(boardId), uint(listId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newTask := new(models.Task)
	*newTask = *createdTask

	taskJson, err := newTask.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("eventType", "UpdateBoard")
	c.Set("IdB", uint(boardId))
	c.Data(http.StatusOK, "application/json; charset=utf-8", taskJson)
}

func (taskHandler *TaskHandler) RefactorTask(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var task models.Task
	err := easyjson.UnmarshalFromReader(c.Request.Body, &task)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	task.IdT = uint(taskId)
	err = taskHandler.usecase.RefactorTask(task, uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newTask, err := taskHandler.usecase.GetSingleTask(task.IdT, uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var isUpdated models.Updated
	isUpdated.UpdatedInfo = true
	isUpdatedJson, err := isUpdated.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("IdB", newTask.IdB)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", newTask.IdT)
	c.Data(http.StatusCreated, "application/json; charset=utf-8", isUpdatedJson)
}

func (taskHandler *TaskHandler) DeleteTask(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}
	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	//вызываю юзкейс

	newTask, err := taskHandler.usecase.GetSingleTask(uint(taskId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = taskHandler.usecase.DeleteTask(uint(taskId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var isDeleted models.Deleted
	isDeleted.DeletedInfo = true
	isDeletedJson, err := isDeleted.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("IdB", newTask.IdB)
	c.Set("eventType", "DeleteTask")
	c.Set("IdT", newTask.IdT)
	c.Data(http.StatusOK, "application/json; charset=utf-8", isDeletedJson)
}

func (taskHandler *TaskHandler) AppendUserToTask(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	appendedUserId, err := strconv.ParseUint(c.Param("idU"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	//вызываю юзкейс

	appendedUser, err := taskHandler.usecase.AppendUserToTask(uint(userId.(uint64)), uint(appendedUserId), uint(taskId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	notification := models.Notification{IdU: uint(appendedUserId),
		NotificationType: "AppendUserToTask", IdT: uint(taskId), IdWh: uint(userId.(uint64))}

	err = taskHandler.notificationUsecase.Create(&notification)
	if err != nil {
		_ = c.Error(err)
		return
	}

	newTask, err := taskHandler.usecase.GetSingleTask(uint(taskId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	userJson, err := appendedUser.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("IdB", newTask.IdB)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", newTask.IdT)
	c.Set("Notification", "AppendUserToTask")
	c.Data(http.StatusOK, "application/json; charset=utf-8", userJson)
}

func (taskHandler *TaskHandler) DeleteUserFromTask(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	deletedUserId, err := strconv.ParseUint(c.Param("idU"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	notification := models.Notification{IdU: uint(deletedUserId),
		NotificationType: "DeleteUserFromTask", IdT: uint(taskId)}

	err = taskHandler.notificationUsecase.Create(&notification)
	if err != nil {
		_ = c.Error(err)
		return
	}

	//вызываю юзкейс

	err = taskHandler.usecase.DeleteUserFromTask(uint(userId.(uint64)), uint(deletedUserId), uint(taskId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newTask, err := taskHandler.usecase.GetSingleTask(uint(taskId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	var isDeleted models.Deleted
	isDeleted.DeletedInfo = true
	isDeletedJson, err := isDeleted.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("IdB", newTask.IdB)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", newTask.IdT)
	c.Set("Notification", "DeleteUserFromTask")
	c.Data(http.StatusOK, "application/json; charset=utf-8", isDeletedJson)
}

func (taskHandler *TaskHandler) AppendUserToTaskByLink(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	link := c.Param("link")

	//вызываю юзкейс

	appendedTask, err := taskHandler.usecase.AppendUserToTaskByLink(uint(userId.(uint64)), link)
	if err != nil {
		_ = c.Error(err)
		return
	}

	taskJson, err := (*appendedTask).MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("IdB", appendedTask.IdB)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", appendedTask.IdT)
	c.Set("Notification", "AppendUserToTask")
	c.Data(http.StatusOK, "application/json; charset=utf-8", taskJson)
}

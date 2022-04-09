package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	usecase usecases.TaskUseCase
}

func MakeTaskHandler(usecase_ usecases.TaskUseCase) *TaskHandler {
	return &TaskHandler{usecase: usecase_}
}

func (taskHandler *TaskHandler) GetTasks(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	listId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := taskHandler.usecase.GetTasks(uint(listId), userId.(uint))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
	return
}

func (taskHandler *TaskHandler) GetSingleTask(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list, err := taskHandler.usecase.GetSingleTask(uint(taskId), userId.(uint))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
	return
}

func (taskHandler *TaskHandler) CreateTask(c *gin.Context) {
	token, check := c.Get("content")
	if !check || token != "application/json" {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadContentType), gin.H{"error": customErrors.ErrBadContentType.Error()})
		return
	}
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	var list models.Task
	err := c.ShouldBindJSON(&list)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	listId, err := strconv.ParseUint(c.Param("idL"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boardId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskId, err := taskHandler.usecase.CreateTask(list, uint(boardId), uint(listId), userId.(uint))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"taskId": taskId})
	return
}

func (taskHandler *TaskHandler) RefactorTask(c *gin.Context) {
	token, check := c.Get("content")
	if !check || token != "application/json" {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadContentType), gin.H{"error": customErrors.ErrBadContentType.Error()})
		return
	}
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = taskHandler.usecase.RefactorTask(task, userId.(uint), uint(taskId))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"updated": true})
	return
}

func (taskHandler *TaskHandler) DeleteTask(c *gin.Context) {
	token, check := c.Get("content")
	if !check || token != "application/json" {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadContentType), gin.H{"error": customErrors.ErrBadContentType.Error()})
		return
	}
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}
	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//вызываю юзкейс

	err = taskHandler.usecase.DeleteTask(uint(taskId), userId.(uint))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
	return
}

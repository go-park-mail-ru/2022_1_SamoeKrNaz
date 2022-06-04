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

type CheckListItemHandler struct {
	usecase     usecases.CheckListItemUseCase
	taskUsecase usecases.TaskUseCase
}

func MakeCheckListItemHandler(usecase_ usecases.CheckListItemUseCase, taskUsecase_ usecases.TaskUseCase) *CheckListItemHandler {
	return &CheckListItemHandler{usecase: usecase_, taskUsecase: taskUsecase_}
}

func (checkListItemHandler *CheckListItemHandler) GetCheckListItems(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	checkListId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	checkListItems, err := checkListItemHandler.usecase.GetCheckListItems(uint(userId.(uint64)), uint(checkListId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newCheckListItems := new(models.CheckListItems)
	*newCheckListItems = *checkListItems
	checkListItemsJson, err := newCheckListItems.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", checkListItemsJson)
}

func (checkListItemHandler *CheckListItemHandler) GetSingleCheckListItem(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	checkListItemId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	checkListItem, err := checkListItemHandler.usecase.GetSingleCheckListItem(uint(checkListItemId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newCheckListItem := new(models.CheckListItem)
	*newCheckListItem = *checkListItem

	checkListItemJson, err := newCheckListItem.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", checkListItemJson)
}

func (checkListItemHandler *CheckListItemHandler) CreateCheckListItem(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var checkListItem models.CheckListItem
	err := easyjson.UnmarshalFromReader(c.Request.Body, &checkListItem)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	checkListId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	createdCheckListItem, err := checkListItemHandler.usecase.CreateCheckListItem(&checkListItem, uint(checkListId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	task, err := checkListItemHandler.taskUsecase.GetSingleTask(createdCheckListItem.IdT, uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newCheckListItem := new(models.CheckListItem)
	*newCheckListItem = *createdCheckListItem
	checkListItemJson, err := newCheckListItem.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("IdB", task.IdB)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", task.IdT)
	c.Data(http.StatusOK, "application/json; charset=utf-8", checkListItemJson)
}

func (checkListItemHandler *CheckListItemHandler) RefactorCheckListItem(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var checkListItem models.CheckListItem
	err := easyjson.UnmarshalFromReader(c.Request.Body, &checkListItem)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	checkListItemId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	checkListItem.IdClIt = uint(checkListItemId)
	err = checkListItemHandler.usecase.RefactorCheckListItem(&checkListItem, uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	getCheckList, err := checkListItemHandler.usecase.GetSingleCheckListItem(uint(checkListItemId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	task, err := checkListItemHandler.taskUsecase.GetSingleTask(getCheckList.IdT, uint(userId.(uint64)))
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
	c.Set("IdB", task.IdB)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", task.IdT)
	c.Data(http.StatusCreated, "application/json; charset=utf-8", isUpdatedJson)
}

func (checkListItemHandler *CheckListItemHandler) DeleteCheckListItem(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	checkListItemId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	getCheckList, err := checkListItemHandler.usecase.GetSingleCheckListItem(uint(checkListItemId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = checkListItemHandler.usecase.DeleteCheckListItem(uint(checkListItemId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	task, err := checkListItemHandler.taskUsecase.GetSingleTask(getCheckList.IdT, uint(userId.(uint64)))
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
	c.Set("IdB", task.IdB)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", task.IdT)
	c.Data(http.StatusOK, "application/json; charset=utf-8", isDeletedJson)
}

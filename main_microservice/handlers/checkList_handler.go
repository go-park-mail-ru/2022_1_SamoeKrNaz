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

type CheckListHandler struct {
	usecase usecases.CheckListUseCase
}

func MakeCheckListHandler(usecase_ usecases.CheckListUseCase) *CheckListHandler {
	return &CheckListHandler{usecase: usecase_}
}

func (checkListHandler *CheckListHandler) GetCheckLists(c *gin.Context) {
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

	checkListSlice, err := checkListHandler.usecase.GetCheckLists(uint(userId.(uint64)), uint(taskId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	checkLists := new(models.CheckLists)
	*checkLists = *checkListSlice

	checkListsJson, err := checkLists.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", checkListsJson)
}

func (checkListHandler *CheckListHandler) GetSingleCheckList(c *gin.Context) {
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

	checkList, err := checkListHandler.usecase.GetSingleCheckList(uint(checkListId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newCheckList := new(models.CheckList)
	*newCheckList = *checkList
	checkListJson, err := newCheckList.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", checkListJson)
}

func (checkListHandler *CheckListHandler) CreateCheckList(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var checkList models.CheckList
	err := easyjson.UnmarshalFromReader(c.Request.Body, &checkList)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	createdCheckList, err := checkListHandler.usecase.CreateCheckList(&checkList, uint(taskId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newCheckList := new(models.CheckList)
	*newCheckList = *createdCheckList
	checkListJson, err := newCheckList.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("IdU", userId)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", taskId)
	c.Data(http.StatusOK, "application/json; charset=utf-8", checkListJson)
}

func (checkListHandler *CheckListHandler) RefactorCheckList(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var checkList models.CheckList
	err := easyjson.UnmarshalFromReader(c.Request.Body, &checkList)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	checkListId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	checkList.IdCl = uint(checkListId)
	err = checkListHandler.usecase.RefactorCheckList(&checkList, uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	task, err := checkListHandler.usecase.GetSingleCheckList(uint(userId.(uint64)), checkList.IdCl)
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
	c.Set("IdU", userId)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", task.IdT)
	c.Data(http.StatusCreated, "application/json; charset=utf-8", isUpdatedJson)
}

func (checkListHandler *CheckListHandler) DeleteCheckList(c *gin.Context) {
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

	err = checkListHandler.usecase.DeleteCheckList(uint(checkListId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	task, err := checkListHandler.usecase.GetSingleCheckList(uint(userId.(uint64)), uint(checkListId))
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
	c.Set("IdU", userId)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", task.IdT)
	c.Data(http.StatusOK, "application/json; charset=utf-8", isDeletedJson)
}

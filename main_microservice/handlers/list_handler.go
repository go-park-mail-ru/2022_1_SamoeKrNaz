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

type ListHandler struct {
	usecase usecases.ListUseCase
}

func MakeListHandler(usecase_ usecases.ListUseCase) *ListHandler {
	return &ListHandler{usecase: usecase_}
}

func (listHandler *ListHandler) GetLists(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	boardId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	var lists models.Lists
	lists, err = listHandler.usecase.GetLists(uint(boardId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	listsJson, err := lists.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", listsJson)
}

func (listHandler *ListHandler) GetSingleList(c *gin.Context) {
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

	list, err := listHandler.usecase.GetSingleList(uint(listId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	listJson, err := list.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", listJson)
}

func (listHandler *ListHandler) CreateList(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var list models.List
	err := easyjson.UnmarshalFromReader(c.Request.Body, &list)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	boardId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	createdList, err := listHandler.usecase.CreateList(list, uint(boardId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newList := new(models.List)
	*newList = *createdList

	listJson, err := newList.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", listJson)
}

func (listHandler *ListHandler) RefactorList(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var list models.List
	err := easyjson.UnmarshalFromReader(c.Request.Body, &list)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	listId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	list.IdL = uint(listId)
	err = listHandler.usecase.RefactorList(list, uint(userId.(uint64)), uint(listId))
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
	c.Data(http.StatusCreated, "application/json; charset=utf-8", isUpdatedJson)
}

func (listHandler *ListHandler) DeleteList(c *gin.Context) {
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

	//вызываю юзкейс

	err = listHandler.usecase.DeleteList(uint(listId), uint(userId.(uint64)))
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
	c.Data(http.StatusOK, "application/json; charset=utf-8", isDeletedJson)
}

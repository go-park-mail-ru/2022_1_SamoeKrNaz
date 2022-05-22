package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AttachmentHandler struct {
	usecase usecases.AttachmentUseCase
}

func MakeAttachmentHandler(usecase_ usecases.AttachmentUseCase) *AttachmentHandler {
	return &AttachmentHandler{usecase: usecase_}
}

func (attachmentHandler *AttachmentHandler) CreateAttachment(c *gin.Context) {
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

	header, err := c.FormFile("attachment")
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	attachment, err := attachmentHandler.usecase.CreateAttachment(header, uint(taskId), uint(userId.(uint64)))

	if err != nil {
		_ = c.Error(err)
		return
	}

	attachmentJson, err := attachment.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", attachmentJson)
}

func (attachmentHandler *AttachmentHandler) GetSingleAttachment(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	attachmentId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	//вызываю юзкейс
	attachment, err := attachmentHandler.usecase.GetById(uint(attachmentId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	attachmentJson, err := attachment.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", attachmentJson)
}

func (attachmentHandler *AttachmentHandler) DeleteAttachment(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	attachemntId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	//вызываю юзкейс

	err = attachmentHandler.usecase.DeleteAttachment(uint(attachemntId), uint(userId.(uint64)))
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

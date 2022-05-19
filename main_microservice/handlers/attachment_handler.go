package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/usecases"
	"fmt"
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
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error in parse id")
		return
	}

	header, err := c.FormFile("attachment")
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		fmt.Println("error in formfile")
		return
	}

	attachment, err := attachmentHandler.usecase.CreateAttachment(header, uint(taskId), uint(userId.(uint64)))

	if err != nil {
		fmt.Println("error in createattachment")
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	c.JSON(http.StatusOK, attachment)
}

func (attachmentHandler *AttachmentHandler) GetSingleAttachment(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	attachmentId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//вызываю юзкейс
	attachment, err := attachmentHandler.usecase.GetById(uint(attachmentId), uint(userId.(uint64)))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attachment)
}

func (attachmentHandler *AttachmentHandler) DeleteAttachment(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	attachemntId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//вызываю юзкейс

	err = attachmentHandler.usecase.DeleteAttachment(uint(attachemntId), uint(userId.(uint64)))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

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

type CommentHandler struct {
	usecase usecases.CommentUseCase
}

func MakeCommentHandler(usecase_ usecases.CommentUseCase) *CommentHandler {
	return &CommentHandler{usecase: usecase_}
}

func (commentHandler *CommentHandler) GetComments(c *gin.Context) {
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

	comments, err := commentHandler.usecase.GetComments(uint(userId.(uint64)), uint(taskId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newComments := new(models.Comments)
	*newComments = *comments

	commentsJson, err := newComments.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", commentsJson)
}

func (commentHandler *CommentHandler) GetSingleComment(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	commentId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	comment, err := commentHandler.usecase.GetSingleComment(uint(commentId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newcomment := new(models.Comment)
	*newcomment = *comment

	commentJson, err := newcomment.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", commentJson)
}

func (commentHandler *CommentHandler) CreateComment(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var comment models.Comment
	err := easyjson.UnmarshalFromReader(c.Request.Body, &comment)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	taskId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	createdComment, err := commentHandler.usecase.CreateComment(&comment, uint(taskId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	newComment := new(models.Comment)
	*newComment = *createdComment

	commentJson, err := newComment.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("IdU", userId)
	c.Set("eventType", "UpdateTask")
	c.Set("IdT", taskId)
	c.Data(http.StatusOK, "application/json; charset=utf-8", commentJson)
}

func (commentHandler *CommentHandler) RefactorComment(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var comment models.Comment
	err := easyjson.UnmarshalFromReader(c.Request.Body, &comment)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	commentId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	comment.IdCm = uint(commentId)
	err = commentHandler.usecase.RefactorComment(&comment, uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	getComment, err := commentHandler.usecase.GetSingleComment(uint(userId.(uint64)), uint(commentId))
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
	c.Set("IdT", getComment.IdT)
	c.Data(http.StatusCreated, "application/json; charset=utf-8", isUpdatedJson)
}

func (commentHandler *CommentHandler) DeleteComment(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	commentId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	err = commentHandler.usecase.DeleteComment(uint(commentId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	getComment, err := commentHandler.usecase.GetSingleComment(uint(userId.(uint64)), uint(commentId))
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
	c.Set("IdT", getComment.IdT)
	c.Data(http.StatusOK, "application/json; charset=utf-8", isDeletedJson)
}

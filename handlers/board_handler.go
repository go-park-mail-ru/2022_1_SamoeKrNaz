package handlers

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BoardHandler struct {
	usecase usecases.BoardUseCase
}

func MakeBoardHandler(usecase_ usecases.BoardUseCase) *BoardHandler {
	return &BoardHandler{usecase: usecase_}
}

func (boardHandler *BoardHandler) GetBoards(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}

	//Получаю доски от БД
	boards, err := boardHandler.usecase.GetBoards(userId.(uint))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, boards)
	return
}

func (boardHandler *BoardHandler) GetSingleBoard(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrUnauthorized), gin.H{"error": customErrors.ErrUnauthorized.Error()})
		return
	}
	boardId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//вызываю юзкейс

	board, err := boardHandler.usecase.GetSingleBoard(uint(boardId), userId.(uint))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, board)
	return
}

func (boardHandler *BoardHandler) CreateBoard(c *gin.Context) {
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

	var board models.Board
	err := c.ShouldBindJSON(&board)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	boardId, err := boardHandler.usecase.CreateBoard(userId.(uint), board)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"boardId": boardId})
	return
}

func (boardHandler *BoardHandler) RefactorBoard(c *gin.Context) {
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

	var board models.Board
	err := c.ShouldBindJSON(&board)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(customErrors.ErrBadInputData), gin.H{"error": customErrors.ErrBadInputData.Error()})
		return
	}

	err = boardHandler.usecase.RefactorBoard(userId.(uint), board)
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"updated": true})
	return
}

func (boardHandler *BoardHandler) DeleteBoard(c *gin.Context) {
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
	boardId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//вызываю юзкейс

	err = boardHandler.usecase.DeleteBoard(uint(boardId), userId.(uint))
	if err != nil {
		c.JSON(customErrors.ConvertErrorToCode(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
	return
}

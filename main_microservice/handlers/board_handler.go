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

type BoardHandler struct {
	usecase usecases.BoardUseCase
}

func MakeBoardHandler(usecase_ usecases.BoardUseCase) *BoardHandler {
	return &BoardHandler{usecase: usecase_}
}

func (boardHandler *BoardHandler) GetBoards(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	//Получаю доски от БД
	var boards models.Boards
	boards, err := boardHandler.usecase.GetBoards(uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	boardsJson, err := boards.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", boardsJson)
}

func (boardHandler *BoardHandler) GetSingleBoard(c *gin.Context) {
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

	//вызываю юзкейс
	board, err := boardHandler.usecase.GetBoard(uint(boardId), uint(userId.(uint64)))
	if err != nil {
		_ = c.Error(err)
		return
	}

	boardJson, err := board.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", boardJson)
}

func (boardHandler *BoardHandler) CreateBoard(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var board models.Board
	err := easyjson.UnmarshalFromReader(c.Request.Body, &board)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	createdBoard, err := boardHandler.usecase.CreateBoard(uint(userId.(uint64)), board)
	if err != nil {
		_ = c.Error(err)
		return
	}

	boardJson, err := createdBoard.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusCreated, "application/json; charset=utf-8", boardJson)
}

func (boardHandler *BoardHandler) RefactorBoard(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var board models.Board
	err := easyjson.UnmarshalFromReader(c.Request.Body, &board)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	boardId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	board.IdB = uint(boardId)
	err = boardHandler.usecase.RefactorBoard(uint(userId.(uint64)), board)
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
	c.Set("eventType", "UpdateBoard")
	c.Set("IdB", boardId)
	c.Data(http.StatusCreated, "application/json; charset=utf-8", isUpdatedJson)
}

func (boardHandler *BoardHandler) DeleteBoard(c *gin.Context) {
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

	//вызываю юзкейс

	err = boardHandler.usecase.DeleteBoard(uint(boardId), uint(userId.(uint64)))
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
	c.Set("eventType", "UpdateBoard")
	c.Set("IdB", uint(boardId))
	c.Data(http.StatusOK, "application/json; charset=utf-8", isDeletedJson)
}

func (boardHandler *BoardHandler) SaveImage(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	var board models.Board
	err := easyjson.UnmarshalFromReader(c.Request.Body, &board)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	boardId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	board.IdB = uint(boardId)

	header, err := c.FormFile("img_board")
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	board.ImgDesk = header.Filename

	path, err := boardHandler.usecase.SaveImage(uint(userId.(uint64)), &board, header)

	if err != nil {
		_ = c.Error(err)
		return
	}

	var imgPath models.ImgBoard
	imgPath.ImgPath = path
	imgPathJson, err := imgPath.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", imgPathJson)
}

func (boardHandler *BoardHandler) AppendUserToBoard(c *gin.Context) {
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

	appendedUserId, err := strconv.ParseUint(c.Param("idU"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	//вызываю юзкейс

	appendedUser, err := boardHandler.usecase.AppendUserToBoard(uint(userId.(uint64)), uint(appendedUserId), uint(boardId))
	if err != nil {
		_ = c.Error(err)
		return
	}

	userJson, err := appendedUser.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("eventType", "UpdateBoard")
	c.Set("IdB", uint(boardId))
	c.Data(http.StatusOK, "application/json; charset=utf-8", userJson)
}

func (boardHandler *BoardHandler) DeleteUserToBoard(c *gin.Context) {
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

	deletedUserId, err := strconv.ParseUint(c.Param("idU"), 10, 32)
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}

	//вызываю юзкейс

	err = boardHandler.usecase.DeleteUserFromBoard(uint(userId.(uint64)), uint(deletedUserId), uint(boardId))
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
	c.Set("eventType", "UpdateBoard")
	c.Set("IdB", uint(boardId))
	c.Data(http.StatusOK, "application/json; charset=utf-8", isDeletedJson)
}

func (boardHandler *BoardHandler) AppendUserToBoardByLink(c *gin.Context) {
	userId, check := c.Get("Auth")
	if !check {
		_ = c.Error(customErrors.ErrUnauthorized)
		return
	}

	link := c.Param("link")

	//вызываю юзкейс

	appendedBoard, err := boardHandler.usecase.AppendUserByLink(uint(userId.(uint64)), link)
	if err != nil {
		_ = c.Error(err)
		return
	}

	boardJson, err := appendedBoard.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Set("eventType", "UpdateBoard")
	c.Set("IdB", appendedBoard.IdB)
	c.Data(http.StatusOK, "application/json; charset=utf-8", boardJson)
}

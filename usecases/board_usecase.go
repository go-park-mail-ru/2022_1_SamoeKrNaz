package usecases

import (
	"PLANEXA_backend/models"
	"mime/multipart"
)

type BoardUseCase interface {
	AppendUserToBoard(userId uint, boardId uint) (models.User, error)
	GetBoards(userId uint) ([]models.Board, error)
	GetSingleBoard(boardId uint, userId uint) (models.Board, error)
	GetBoard(boardId uint, userId uint) (models.Board, error)
	CreateBoard(userId uint, board models.Board) (*models.Board, error)
	RefactorBoard(userId uint, board models.Board) error
	DeleteBoard(boardId uint, userId uint) error
	SaveImage(userId uint, board *models.Board, header *multipart.FileHeader) (string, error)
}

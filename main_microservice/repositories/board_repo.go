package repositories

import (
	"PLANEXA_backend/models"
	"mime/multipart"
)

type BoardRepository interface {
	Create(board *models.Board) (uint, error)
	AppendUser(boardId uint, userId uint) error
	GetLists(IdB uint) ([]models.List, error)
	Update(board models.Board) error
	Delete(IdB uint) error
	GetUserBoards(IdU uint) ([]models.Board, error)
	GetById(IdB uint) (*models.Board, error)
	IsAccessToBoard(IdU uint, IdB uint) (bool, error)
	SaveImage(board *models.Board, header *multipart.FileHeader) error
	GetBoardUser(IdB uint) ([]models.User, error)
	DeleteUser(boardId uint, userId uint) error
	GetByLink(link string) (*models.Board, error)
}

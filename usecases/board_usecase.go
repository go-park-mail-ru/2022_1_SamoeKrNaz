package usecases

import "PLANEXA_backend/models"

type BoardUseCase interface {
	GetBoards(userId uint) ([]models.Board, error)
	GetSingleBoard(boardId uint, userId uint) (models.Board, error)
	CreateBoard(userId uint, board models.Board) error
	RefactorBoard(userId uint, board models.Board) error
	DeleteBoard(boardId uint, userId uint) error
}

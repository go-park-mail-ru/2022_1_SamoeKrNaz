package usecases

import "PLANEXA_backend/models"

type ListUsecase interface {
	GetLists(boardId uint, userId uint) ([]models.List, error)
	GetSingleList(listId uint, userId uint) (models.List, error)
	CreateList(list models.List, userId uint) (uint, error)
	RefactorList(list models.List, userId uint, boardId uint) error
	DeleteList(listId uint, userId uint) error
}

package usecases

import "PLANEXA_backend/models"

type CommentUseCase interface {
	GetComments(userId uint, IdT uint) (*[]models.Comment, error)
	GetSingleComment(userId uint, IdCm uint) (*models.Comment, error)
	CreateComment(comment *models.Comment, IdT uint, userId uint) (*models.Comment, error)
	RefactorComment(comment *models.Comment, userId uint) error
	DeleteComment(IdCm uint, userId uint) error
}

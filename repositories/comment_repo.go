package repositories

import "PLANEXA_backend/models"

type CommentRepository interface {
	Create(comment *models.Comment) (uint, error)
	GetById(IdCm uint) (*models.Comment, error)
	Update(comment models.Comment) error
	Delete(IdCm uint) error
	GetComments(IdT uint) (*[]models.Comment, error)
}

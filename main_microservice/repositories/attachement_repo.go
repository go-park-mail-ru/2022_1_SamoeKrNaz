package repositories

import (
	"PLANEXA_backend/models"
	"mime/multipart"
)

type AttachmentRepository interface {
	Create(header *multipart.FileHeader, IdT uint) (*models.Attachment, error)
	Delete(IdA uint) error
	GetById(IdA uint) (*models.Attachment, error)
}

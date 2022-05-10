package usecases

import (
	"PLANEXA_backend/models"
	"mime/multipart"
)

type AttachmentUseCase interface {
	CreateAttachment(header *multipart.FileHeader, taskId uint, userId uint) (*models.Attachment, error)
	DeleteAttachment(attachId uint, userId uint) error
	GetById(attachId uint, userId uint) (*models.Attachment, error)
}

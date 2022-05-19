package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"fmt"
	"mime/multipart"
)

type AttachmentUseCaseImpl struct {
	repAttachment repositories.AttachmentRepository
	repTask       repositories.TaskRepository
}

func MakeAttachmentUseCase(repAttachment_ repositories.AttachmentRepository, repTask_ repositories.TaskRepository) usecases.AttachmentUseCase {
	return &AttachmentUseCaseImpl{repAttachment: repAttachment_, repTask: repTask_}
}

func (attachmentUseCase AttachmentUseCaseImpl) CreateAttachment(header *multipart.FileHeader, taskId uint, userId uint) (*models.Attachment, error) {
	isAccess, err := attachmentUseCase.repTask.IsAccessToTask(userId, taskId)
	fmt.Println("usecase error 1", err)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	attachment, err := attachmentUseCase.repAttachment.Create(header, taskId)
	fmt.Println("usecase error 2", err)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}
func (attachmentUseCase AttachmentUseCaseImpl) DeleteAttachment(attachId uint, userId uint) error {
	attachment, err := attachmentUseCase.repAttachment.GetById(attachId)
	if err != nil {
		return err
	}
	isAccess, err := attachmentUseCase.repTask.IsAccessToTask(userId, attachment.IdT)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	err = attachmentUseCase.repAttachment.Delete(attachId)
	if err != nil {
		return err
	}
	return nil
}
func (attachmentUseCase AttachmentUseCaseImpl) GetById(attachId uint, userId uint) (*models.Attachment, error) {
	attachment, err := attachmentUseCase.repAttachment.GetById(attachId)
	if err != nil {
		return nil, err
	}
	isAccess, err := attachmentUseCase.repTask.IsAccessToTask(userId, attachment.IdT)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	return attachment, nil
}

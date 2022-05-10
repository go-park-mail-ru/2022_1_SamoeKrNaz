package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/models"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

const filePathAttach = "img_boards/"

type AttachmentRepositoryImpl struct {
	db *gorm.DB
}

func MakeAttachmentRepository(db *gorm.DB) repositories.AttachmentRepository {
	return &AttachmentRepositoryImpl{db: db}
}

func (attachmentRepository AttachmentRepositoryImpl) Create(header *multipart.FileHeader, IdT uint) (*models.Attachment, error) {
	var countFiles int64
	attachment := new(models.Attachment)
	attachment.IdT = IdT
	err := attachmentRepository.db.Model(&models.Attachment{}).Count(&countFiles).Error
	fileName := strings.Join([]string{filePathAttach, strconv.Itoa(int(countFiles))}, "")
	attachment.SystemName = fileName
	attachment.DefaultName = header.Filename

	output, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer output.Close()

	openFile, err := header.Open()
	if err != nil {
		return nil, err
	}

	defer openFile.Close()

	_, err = io.Copy(output, openFile)

	if err != nil {
		return nil, err
	}

	err = attachmentRepository.db.Create(attachment).Error
	return attachment, err
}

func (attachmentRepository AttachmentRepositoryImpl) Delete(IdA uint) error {
	attachment, err := attachmentRepository.GetById(IdA)
	if err != nil {
		return err
	}
	err = os.Remove(attachment.SystemName)
	if err != nil {
		return err
	}
	err = attachmentRepository.db.Delete(&models.Attachment{}, IdA).Error
	if err != nil {
		return err
	}
	return nil
}

func (attachmentRepository AttachmentRepositoryImpl) GetById(IdA uint) (*models.Attachment, error) {
	attachment := new(models.Attachment)
	result := attachmentRepository.db.Find(attachment, IdA)
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrAttachmentNotFound
	} else if result.Error != nil {
		return nil, result.Error
	} else {
		return attachment, nil
	}
}

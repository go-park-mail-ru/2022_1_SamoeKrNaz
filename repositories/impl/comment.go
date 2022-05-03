package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func MakeCommentRepository(db *gorm.DB) repositories.CommentRepository {
	return &CommentRepositoryImpl{db: db}
}

func (commentRepository *CommentRepositoryImpl) Create(comment *models.Comment) (uint, error) {
	err := commentRepository.db.Create(comment).Error
	return comment.IdCm, err
}

func (commentRepository *CommentRepositoryImpl) GetById(IdCm uint) (*models.Comment, error) {
	comment := new(models.Comment)
	result := commentRepository.db.Find(comment, IdCm)
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrCommentNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}

func (commentRepository *CommentRepositoryImpl) Update(comment models.Comment) error {
	currentData, err := commentRepository.GetById(comment.IdCm)
	if err != nil {
		return err
	}
	if currentData.Text != comment.Text && comment.Text != "" {
		currentData.Text = comment.Text
	}
	return commentRepository.db.Save(currentData).Error
}

func (commentRepository *CommentRepositoryImpl) Delete(IdCm uint) error {
	return commentRepository.db.Delete(&models.Comment{}, IdCm).Error
}

func (commentRepository *CommentRepositoryImpl) GetComments(IdT uint) (*[]models.Comment, error) {
	comments := new([]models.Comment)
	err := commentRepository.db.Where("id_t = ?", IdT).Find(comments).Error
	return comments, err
}

func (commentRepository *CommentRepositoryImpl) IsAccessToComment(IdCm uint, IdU uint) (bool, error) {
	comment := commentRepository.db.Where("id_cm = ? and id_u = ?", IdCm, IdU).Find(&models.Comment{})
	if comment.Error != nil {
		return false, comment.Error
	} else if comment == nil {
		return false, nil
	}
	return true, nil
}

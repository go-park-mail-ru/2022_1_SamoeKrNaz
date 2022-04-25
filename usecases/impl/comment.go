package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"PLANEXA_backend/usecases"
	"github.com/microcosm-cc/bluemonday"
)

type CommentUseCaseImpl struct {
	repComment repositories.CommentRepository
	repTask    repositories.TaskRepository
}

func MakeCommentUsecase(repComment_ repositories.CommentRepository, repTask_ repositories.TaskRepository) usecases.CommentUseCase {
	return &CommentUseCaseImpl{repComment: repComment_, repTask: repTask_}
}

func (commentUseCase *CommentUseCaseImpl) GetComments(userId uint, IdT uint) (*[]models.Comment, error) {
	comments, err := commentUseCase.repComment.GetComments(IdT)
	if err != nil {
		return nil, err
	}
	isAccess, err := commentUseCase.repTask.IsAccessToTask(userId, (*comments)[0].IdT)
	if err != nil {
		return nil, err
	} else if isAccess == false {
		return nil, customErrors.ErrNoAccess
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, comment := range *comments {
		comment.Text = sanitizer.Sanitize(comment.Text)
	}
	return comments, nil
}

func (commentUseCase *CommentUseCaseImpl) GetSingleComment(userId uint, IdCm uint) (*models.Comment, error) {
	comment, err := commentUseCase.repComment.GetById(IdCm)
	if err != nil {
		return nil, err
	}
	isAccess, err := commentUseCase.repTask.IsAccessToTask(userId, comment.IdT)
	if err != nil {
		return nil, err
	} else if isAccess == false {
		return nil, customErrors.ErrNoAccess
	}
	sanitizer := bluemonday.UGCPolicy()
	comment.Text = sanitizer.Sanitize(comment.Text)
	return comment, nil
}

func (commentUseCase *CommentUseCaseImpl) CreateComment(comment *models.Comment, IdT uint, userId uint) (*models.Comment, error) {
	comment.IdT = IdT
	isAccess, err := commentUseCase.repTask.IsAccessToTask(userId, comment.IdT)
	if err != nil {
		return nil, err
	} else if isAccess == false {
		return nil, customErrors.ErrNoAccess
	}
	checkListId, err := commentUseCase.repComment.Create(comment)
	if err != nil {
		return nil, err
	}
	createdComment, err := commentUseCase.repComment.GetById(checkListId)
	return createdComment, nil
}

func (commentUseCase *CommentUseCaseImpl) RefactorComment(comment *models.Comment, userId uint) error {
	currentData, err := commentUseCase.repComment.GetById(comment.IdCm)
	if err != nil {
		return err
	}
	isAccess, err := commentUseCase.repTask.IsAccessToTask(userId, currentData.IdT)
	if err != nil {
		return err
	} else if isAccess == false {
		return customErrors.ErrNoAccess
	}
	return commentUseCase.repComment.Update(*comment)
}

func (commentUseCase *CommentUseCaseImpl) DeleteComment(IdCm uint, userId uint) error {
	checkList, err := commentUseCase.repComment.GetById(IdCm)
	if err != nil {
		return err
	}
	isAccess, err := commentUseCase.repTask.IsAccessToTask(userId, checkList.IdT)
	if err != nil {
		return err
	} else if isAccess == false {
		return customErrors.ErrNoAccess
	}
	err = commentUseCase.repComment.Delete(IdCm)
	return err
}

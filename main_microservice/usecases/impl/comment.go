package impl

import (
	customErrors "PLANEXA_backend/errors"
	repositories2 "PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	rtime "github.com/ivahaev/russian-time"
	"github.com/microcosm-cc/bluemonday"
	"strconv"
	"time"
)

type CommentUseCaseImpl struct {
	repComment repositories2.CommentRepository
	repTask    repositories2.TaskRepository
	repUser    repositories2.UserRepository
}

func MakeCommentUsecase(repComment_ repositories2.CommentRepository, repTask_ repositories2.TaskRepository, repUser_ repositories2.UserRepository) usecases.CommentUseCase {
	return &CommentUseCaseImpl{repComment: repComment_, repTask: repTask_, repUser: repUser_}
}

func (commentUseCase *CommentUseCaseImpl) GetComments(userId uint, IdT uint) (*[]models.Comment, error) {
	comments, err := commentUseCase.repComment.GetComments(IdT)
	if err != nil {
		return nil, err
	}
	isAccess, err := commentUseCase.repTask.IsAccessToTask(userId, (*comments)[0].IdT)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, comment := range *comments {
		comment.Text = sanitizer.Sanitize(comment.Text)
		user, err := commentUseCase.repUser.GetUserById(comment.IdU)
		if err != nil {
			return nil, err
		}
		comment.User = *user
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
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	sanitizer := bluemonday.UGCPolicy()
	comment.Text = sanitizer.Sanitize(comment.Text)
	user, err := commentUseCase.repUser.GetUserById(comment.IdU)
	if err != nil {
		return nil, err
	}
	comment.User = *user
	return comment, nil
}

func (commentUseCase *CommentUseCaseImpl) CreateComment(comment *models.Comment, IdT uint, userId uint) (*models.Comment, error) {
	comment.IdT = IdT
	comment.IdU = userId
	isAccess, err := commentUseCase.repTask.IsAccessToTask(userId, comment.IdT)
	if err != nil {
		return nil, err
	} else if !isAccess {
		return nil, customErrors.ErrNoAccess
	}
	moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, err
	}
	comment.DateCreated = strconv.Itoa(time.Now().In(moscow).Day()) + " " + rtime.Now().Month().StringInCase() + " " + strconv.Itoa(time.Now().In(moscow).Year()) + ", " + time.Now().In(moscow).Format("15:04")
	commentId, err := commentUseCase.repComment.Create(comment)
	if err != nil {
		return nil, err
	}
	createdComment, err := commentUseCase.repComment.GetById(commentId)
	return createdComment, err
}

func (commentUseCase *CommentUseCaseImpl) RefactorComment(comment *models.Comment, userId uint) error {
	isAccess, err := commentUseCase.repComment.IsAccessToComment(userId, comment.IdCm)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	return commentUseCase.repComment.Update(*comment)
}

func (commentUseCase *CommentUseCaseImpl) DeleteComment(IdCm uint, userId uint) error {
	isAccess, err := commentUseCase.repComment.IsAccessToComment(userId, IdCm)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	err = commentUseCase.repComment.Delete(IdCm)
	return err
}

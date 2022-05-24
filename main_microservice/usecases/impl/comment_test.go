package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories/mocks"
	"PLANEXA_backend/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetComments(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentUseCase := MakeCommentUsecase(commentRepo, taskRepo, userRepo)

	user := models.User{}
	comments := []models.Comment{{Text: "title1"}}
	commentUser := []models.Comment{{Text: "title1"}}
	commentRepo.EXPECT().GetComments(uint(22)).Return(&comments, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(true, nil)
	userRepo.EXPECT().GetUserById(uint(0)).Return(&user, nil)
	newComments, err := commentUseCase.GetComments(uint(11), uint(22))
	assert.Equal(t, &commentUser, newComments)
	assert.Equal(t, err, nil)

	commentRepo.EXPECT().GetComments(uint(22)).Return(&comments, nil)
	taskRepo.EXPECT().IsAccessToTask(uint(11), uint(0)).Return(false, customErrors.ErrBadInputData)
	newComments, err = commentUseCase.GetComments(uint(11), uint(22))
	assert.Equal(t, (*[]models.Comment)(nil), newComments)
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestCreateComment(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentUseCase := MakeCommentUsecase(commentRepo, taskRepo, userRepo)

	comment := models.Comment{}

	taskRepo.EXPECT().IsAccessToTask(uint(15), uint(11)).Return(true, nil)
	commentRepo.EXPECT().Create(gomock.Any()).Return(uint(22), nil)
	commentRepo.EXPECT().GetById(uint(22)).Return(&comment, nil)
	_, err := commentUseCase.CreateComment(&comment, uint(11), uint(15))
	assert.Equal(t, err, nil)

	taskRepo.EXPECT().IsAccessToTask(uint(15), uint(11)).Return(false, customErrors.ErrNoAccess)
	_, err = commentUseCase.CreateComment(&comment, uint(11), uint(15))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestRefactorComment(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentUseCase := MakeCommentUsecase(commentRepo, taskRepo, userRepo)

	comment := models.Comment{IdCm: 15}

	commentRepo.EXPECT().IsAccessToComment(uint(15), uint(11)).Return(true, nil)
	commentRepo.EXPECT().Update(comment).Return(nil)
	err := commentUseCase.RefactorComment(&comment, uint(11))
	assert.Equal(t, err, nil)

	commentRepo.EXPECT().IsAccessToComment(uint(15), uint(11)).Return(false, customErrors.ErrNoAccess)
	err = commentUseCase.RefactorComment(&comment, uint(11))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestDeleteComment(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentUseCase := MakeCommentUsecase(commentRepo, taskRepo, userRepo)

	commentRepo.EXPECT().IsAccessToComment(uint(15), uint(11)).Return(true, nil)
	commentRepo.EXPECT().Delete(uint(11)).Return(nil)
	err := commentUseCase.DeleteComment(uint(11), uint(15))
	assert.Equal(t, err, nil)

	commentRepo.EXPECT().IsAccessToComment(uint(15), uint(11)).Return(false, customErrors.ErrNoAccess)
	err = commentUseCase.DeleteComment(uint(11), uint(15))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

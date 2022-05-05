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

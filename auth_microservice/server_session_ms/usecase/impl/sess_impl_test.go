package usecase_impl

import (
	mock_sess "PLANEXA_backend/auth_microservice/server_session_ms/repository/mocks"
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetSession(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	sessRepo := mock_sess.NewMockSessionRedis(controller)
	sessUseCase := CreateSessionUseCase(sessRepo)

	sessRepo.EXPECT().SetSession(models.Session{UserId: 22}).Return(nil)
	err := sessUseCase.SetSession(models.Session{UserId: 22})
	assert.Equal(t, err, nil)

	sessRepo.EXPECT().SetSession(models.Session{}).Return(customErrors.ErrBadInputData)
	err = sessUseCase.SetSession(models.Session{})
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestGetSession(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	sessRepo := mock_sess.NewMockSessionRedis(controller)
	sessUseCase := CreateSessionUseCase(sessRepo)

	sessRepo.EXPECT().GetSession("cookie").Return(uint64(1), nil)
	number, err := sessUseCase.GetSession("cookie")
	assert.Equal(t, number, uint64(1))
	assert.Equal(t, err, nil)

	sessRepo.EXPECT().GetSession("cookie").Return(uint64(0), customErrors.ErrBadInputData)
	number, err = sessUseCase.GetSession("cookie")
	assert.Equal(t, number, uint64(0))
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestDeleteSession(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	sessRepo := mock_sess.NewMockSessionRedis(controller)
	sessUseCase := CreateSessionUseCase(sessRepo)

	sessRepo.EXPECT().DeleteSession("cookie").Return(nil)
	err := sessUseCase.DeleteSession("cookie")
	assert.Equal(t, err, nil)

	sessRepo.EXPECT().DeleteSession("cookie").Return(customErrors.ErrBadInputData)
	err = sessUseCase.DeleteSession("cookie")
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

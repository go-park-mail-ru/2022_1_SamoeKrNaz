package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories/mocks"
	"PLANEXA_backend/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"testing"
)

func TestLogin(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Parallel()
	userRepo := mock_repositories.NewMockUserRepository(controller)
	sessRepo := mock_repositories.NewMockSessionRepository(controller)
	userUseCase := MakeUserUsecase(userRepo, sessRepo)

	user := models.User{IdU: 22, Username: "user1"}
	//session := models.Session{UserId: 22, CookieValue: "123"}

	userRepo.EXPECT().IsAbleToLogin(user.Username, user.Password).Return(true, nil)
	userRepo.EXPECT().GetUserByLogin(user.Username).Return(&user, nil)
	sessRepo.EXPECT().SetSession(gomock.Any()).Return(nil)
	_, _, err := userUseCase.Login(user)
	assert.Equal(t, err, nil)
}

func TestSaveAvatar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Parallel()
	userRepo := mock_repositories.NewMockUserRepository(controller)
	sessRepo := mock_repositories.NewMockSessionRepository(controller)
	userUseCase := MakeUserUsecase(userRepo, sessRepo)

	header := new(multipart.FileHeader)
	user := models.User{IdU: 22, Username: "user1"}
	userRepo.EXPECT().SaveAvatar(&user, header).Return(nil)
	_, err := userUseCase.SaveAvatar(&user, header)
	assert.Equal(t, nil, err)
}

func TestRegister(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Parallel()
	userRepo := mock_repositories.NewMockUserRepository(controller)
	sessRepo := mock_repositories.NewMockSessionRepository(controller)
	userUseCase := MakeUserUsecase(userRepo, sessRepo)

	user := models.User{IdU: 22, Username: "user1", Password: "123123123"}
	//session := models.Session{UserId: 22, CookieValue: "123"}

	userRepo.EXPECT().IsExist(user.Username).Return(false, nil)
	userRepo.EXPECT().Create(gomock.Any()).Return(uint(22), nil)
	sessRepo.EXPECT().SetSession(gomock.Any()).Return(nil)
	_, _, err := userUseCase.Register(user)
	assert.Equal(t, err, nil)
}

func TestLogout(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()

	userRepo := mock_repositories.NewMockUserRepository(controller)
	sessRepo := mock_repositories.NewMockSessionRepository(controller)
	userUseCase := MakeUserUsecase(userRepo, sessRepo)

	sessRepo.EXPECT().DeleteSession("123").Return(nil)
	err := userUseCase.Logout("123")
	assert.Equal(t, nil, err)

}

func TestGetInfoById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Parallel()
	userRepo := mock_repositories.NewMockUserRepository(controller)
	sessRepo := mock_repositories.NewMockSessionRepository(controller)
	userUseCase := MakeUserUsecase(userRepo, sessRepo)

	user := models.User{IdU: 22, Username: "user1"}

	userRepo.EXPECT().GetUserById(uint(22)).Return(&user, nil)
	newUser, err := userUseCase.GetInfoById(uint(22))
	assert.Equal(t, user, newUser)
	assert.Equal(t, err, nil)
}

func TestRefactorProfile(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Parallel()
	userRepo := mock_repositories.NewMockUserRepository(controller)
	sessRepo := mock_repositories.NewMockSessionRepository(controller)
	userUseCase := MakeUserUsecase(userRepo, sessRepo)
	user := models.User{IdU: 22, Username: "user1"}

	userRepo.EXPECT().Update(&user).Return(nil)
	err := userUseCase.RefactorProfile(user)
	assert.Equal(t, err, nil)
}

func TestGetUsersLike(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Parallel()
	userRepo := mock_repositories.NewMockUserRepository(controller)
	sessRepo := mock_repositories.NewMockSessionRepository(controller)
	userUseCase := MakeUserUsecase(userRepo, sessRepo)
	users := []models.User{{IdU: 22, Username: "user1"}, {IdU: 22, Username: "user1"}}

	userRepo.EXPECT().GetUsersLike("lol").Return(&users, nil)
	newUsers, err := userUseCase.GetUsersLike("lol")
	assert.Equal(t, err, nil)
	assert.Equal(t, users, *newUsers)

	userRepo.EXPECT().GetUsersLike("lol").Return(nil, customErrors.ErrNoAccess)
	newUsers, err = userUseCase.GetUsersLike("lol")
	assert.Equal(t, (*[]models.User)(nil), newUsers)
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

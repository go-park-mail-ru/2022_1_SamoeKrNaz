package usecase_impl

import (
	"PLANEXA_backend/models"
	"PLANEXA_backend/user_microservice/server_user/repository"
	"PLANEXA_backend/user_microservice/server_user/usecase"
)

type UserUseCaseImpl struct {
	userRepo repository.UserRepo
}

func CreateUserUseCase(userRep repository.UserRepo) usecase.UserUseCase {
	return &UserUseCaseImpl{userRepo: userRep}
}

func (userUseCase *UserUseCaseImpl) Create(user *models.User) (uint, error) {
	return userUseCase.userRepo.Create(user)
}

func (userUseCase *UserUseCaseImpl) Update(user *models.User) error {
	return userUseCase.userRepo.Update(user)
}

func (userUseCase *UserUseCaseImpl) IsAbleToLogin(password string, username string) (bool, error) {
	return userUseCase.userRepo.IsAbleToLogin(username, password)
}

func (userUseCase *UserUseCaseImpl) AddUserToBoard(idU uint, idB uint) error {
	return userUseCase.userRepo.AddUserToBoard(idB, idU)
}

func (userUseCase *UserUseCaseImpl) GetUserByLogin(username string) (*models.User, error) {
	return userUseCase.userRepo.GetUserByLogin(username)
}

func (userUseCase *UserUseCaseImpl) GetUserById(userId uint) (*models.User, error) {
	return userUseCase.userRepo.GetUserById(userId)
}

func (userUseCase *UserUseCaseImpl) IsExist(username string) (bool, error) {
	return userUseCase.userRepo.IsExist(username)
}

func (userUseCase *UserUseCaseImpl) GetUsersLike(username string) (*[]models.User, error) {
	return userUseCase.userRepo.GetUsersLike(username)
}

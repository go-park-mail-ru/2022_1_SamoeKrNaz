package usecase

import (
	"PLANEXA_backend/models"
)

type UserUseCase interface {
	Create(user *models.User) (uint, error)
	Update(user *models.User) error
	IsAbleToLogin(password string, username string) (bool, error)
	AddUserToBoard(idU uint, idB uint) error
	GetUserByLogin(username string) (*models.User, error)
	GetUserById(userId uint) (*models.User, error)
	IsExist(username string) (bool, error)
	GetUsersLike(username string) (*[]models.User, error)
}

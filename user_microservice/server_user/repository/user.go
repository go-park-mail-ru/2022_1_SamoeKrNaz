package repository

import (
	"PLANEXA_backend/models"
)

type UserRepo interface {
	Create(user *models.User) (uint, error)
	Update(user *models.User) error
	IsAbleToLogin(username string, password string) (bool, error)
	AddUserToBoard(IdB uint, IdU uint) error
	GetUserByLogin(username string) (*models.User, error)
	GetUserById(IdU uint) (*models.User, error)
	IsExist(username string) (bool, error)
	GetUsersLike(username string) (*[]models.User, error)
}

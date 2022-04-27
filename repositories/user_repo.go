package repositories

import (
	"PLANEXA_backend/models"
	"mime/multipart"
)

type UserRepository interface {
	Create(user *models.User) (uint, error)
	Update(user *models.User) error
	SaveAvatar(user *models.User, header *multipart.FileHeader) error
	IsAbleToLogin(username string, password string) (bool, error)
	AddUserToBoard(IdB uint, IdU uint) error
	GetUserByLogin(username string) (*models.User, error)
	GetUserById(IdU uint) (*models.User, error)
	IsExist(username string) (bool, error)
	GetUsersLike(username string) (*[]models.User, error)
}

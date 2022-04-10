package usecases

import (
	"PLANEXA_backend/models"
	"mime/multipart"
)

type UserUseCase interface {
	Login(user models.User) (string, error)
	Register(user models.User) (string, error)
	Logout(token string) error
	GetInfo(userId uint) (models.User, error)
	SaveAvatar(*models.User, *multipart.FileHeader) error
}

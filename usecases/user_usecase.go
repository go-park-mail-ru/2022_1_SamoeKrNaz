package usecases

import (
	"PLANEXA_backend/models"
	"mime/multipart"
)

type UserUseCase interface {
	Login(user models.User) (uint, string, error)
	Register(user models.User) (uint, string, error)
	Logout(token string) error
	GetInfoById(userId uint) (models.User, error)
	SaveAvatar(*models.User, *multipart.FileHeader) (string, error)
	RefactorProfile(user models.User) error
}

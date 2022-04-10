package usecases

import (
	"PLANEXA_backend/models"
)

type UserUseCase interface {
	Login(user models.User) (uint, string, error)
	Register(user models.User) (uint, string, error)
	Logout(token string) error
	GetInfoById(userId uint) (models.User, error)
	GetInfoByCookie(token string) (models.User, error)
}

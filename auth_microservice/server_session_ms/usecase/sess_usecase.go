package usecase

import "PLANEXA_backend/models"

type SessionUseCase interface {
	SetSession(session models.Session) error
	GetSession(cookieValue string) (uint64, error)
	DeleteSession(cookieValue string) error
}

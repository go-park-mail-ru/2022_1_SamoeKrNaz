package repository

import "PLANEXA_backend/models"

type SessionRedis interface {
	SetSession(session models.Session) error
	GetSession(cookieValue string) (uint64, error)
	DeleteSession(cookieValue string) error
}

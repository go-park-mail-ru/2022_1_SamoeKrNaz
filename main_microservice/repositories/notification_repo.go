package repositories

import (
	"PLANEXA_backend/models"
)

type NotificationRepository interface {
	Create(notification *models.Notification) error
	GetUsersNotifications(IdU uint) (*[]models.Notification, error)
	ReadNotifications(IdU uint) error
}

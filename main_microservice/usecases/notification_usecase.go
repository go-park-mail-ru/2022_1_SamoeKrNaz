package usecases

import (
	"PLANEXA_backend/models"
)

type NotificationUseCase interface {
	Create(notification *models.Notification) error
	GetUsersNotifications(IdU uint) (*models.Notifications, error)
	CreateBoardNotification(notification *models.Notification) error
	ReadNotifications(IdU uint) error
	DeleteNotifications(IdU uint) error
}

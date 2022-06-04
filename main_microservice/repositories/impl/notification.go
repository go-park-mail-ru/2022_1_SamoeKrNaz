package impl

import (
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/models"
	"gorm.io/gorm"
)

type NotificationRepositoryImpl struct {
	db *gorm.DB
}

func MakeNotificationRepository(db *gorm.DB) repositories.NotificationRepository {
	return &NotificationRepositoryImpl{db: db}
}

func (notificationRepository *NotificationRepositoryImpl) Create(notification *models.Notification) error {
	err := notificationRepository.db.Create(notification).Error
	return err
}

func (notificationRepository *NotificationRepositoryImpl) GetUsersNotifications(IdU uint) (*[]models.Notification, error) {
	notifications := new([]models.Notification)
	err := notificationRepository.db.Where("id_u = ?", IdU).Order("date_to_order DESC").Find(notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (notificationRepository *NotificationRepositoryImpl) ReadNotifications(IdU uint) error {
	err := notificationRepository.db.Model(&models.Notification{}).
		Where("id_u = ?", IdU).
		UpdateColumn("is_read", gorm.Expr("true")).Error
	return err
}

func (notificationRepository *NotificationRepositoryImpl) DeleteNotifications(IdU uint) error {
	err := notificationRepository.db.Where("id_u = ?", IdU).Delete(&models.Notification{}).Error
	if err != nil {
		return err
	}
	return nil
}

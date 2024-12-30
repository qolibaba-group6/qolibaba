// internal/notification-service/repository/notification_repository.go
package repository

import (
	"travel-booking-app/internal/notification-service/model"

	"travel-booking-app/internal/database"
)

type NotificationRepository interface {
	CreateNotification(notification *model.Notification) error
	GetNotificationByID(id string) (*model.Notification, error)
	GetNotificationsByUserID(userID string) ([]model.Notification, error)
	UpdateNotification(notification *model.Notification) error
	DeleteNotification(id string) error
}

type notificationRepository struct{}

func NewNotificationRepository() NotificationRepository {
	return &notificationRepository{}
}

func (r *notificationRepository) CreateNotification(notification *model.Notification) error {
	return database.DB.Create(notification).Error
}

func (r *notificationRepository) GetNotificationByID(id string) (*model.Notification, error) {
	var notification model.Notification
	err := database.DB.First(&notification, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *notificationRepository) GetNotificationsByUserID(userID string) ([]model.Notification, error) {
	var notifications []model.Notification
	err := database.DB.Where("user_id = ?", userID).Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *notificationRepository) UpdateNotification(notification *model.Notification) error {
	return database.DB.Save(notification).Error
}

func (r *notificationRepository) DeleteNotification(id string) error {
	return database.DB.Delete(&model.Notification{}, "id = ?", id).Error
}

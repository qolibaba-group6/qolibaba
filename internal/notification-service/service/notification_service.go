
package service
import (
	"errors"
	"travel-booking-app/internal/notification-service/model"
	"travel-booking-app/internal/notification-service/repository"
	"github.com/google/uuid"
)
type NotificationService interface {
	CreateNotification(userID, message, notificationType string) (string, error)
	GetNotificationByID(id string) (*model.Notification, error)
	GetNotificationsByUserID(userID string) ([]model.Notification, error)
	UpdateNotification(id, message, notificationType string, isRead bool) error
	DeleteNotification(id string) error
}
type notificationService struct {
	repo repository.NotificationRepository
}
func NewNotificationService(repo repository.NotificationRepository) NotificationService {
	return &notificationService{
		repo: repo,
	}
}
func (s *notificationService) CreateNotification(userID, message, notificationType string) (string, error) {
	if userID == "" || message == "" || notificationType == "" {
		return "", errors.New("invalid input data")
	}
	notification := &model.Notification{
		UserID:  uuid.MustParse(userID),
		Message: message,
		Type:    notificationType,
		IsRead:  false,
	}
	err := s.repo.CreateNotification(notification)
	if err != nil {
		return "", err
	}
	return notification.ID.String(), nil
}
func (s *notificationService) GetNotificationByID(id string) (*model.Notification, error) {
	return s.repo.GetNotificationByID(id)
}
func (s *notificationService) GetNotificationsByUserID(userID string) ([]model.Notification, error) {
	return s.repo.GetNotificationsByUserID(userID)
}
func (s *notificationService) UpdateNotification(id, message, notificationType string, isRead bool) error {
	notification, err := s.repo.GetNotificationByID(id)
	if err != nil {
		return err
	}
	if message != "" {
		notification.Message = message
	}
	if notificationType != "" {
		notification.Type = notificationType
	}
	notification.IsRead = isRead
	return s.repo.UpdateNotification(notification)
}
func (s *notificationService) DeleteNotification(id string) error {
	return s.repo.DeleteNotification(id)
}

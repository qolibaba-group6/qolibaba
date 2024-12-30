
package handler
import (
	"encoding/json"
	"net/http"
	"travel-booking-app/internal/notification-service/repository"
	"travel-booking-app/internal/notification-service/service"
	"time"
	"github.com/gorilla/mux"
)
type NotificationHandler struct {
	notificationService service.NotificationService
}
func NewNotificationHandler() *NotificationHandler {
	repo := repository.NewNotificationRepository()
	service := service.NewNotificationService(repo)
	return &NotificationHandler{
		notificationService: service,
	}
}
type CreateNotificationRequest struct {
	UserID           string `json:"user_id"`
	Message          string `json:"message"`
	NotificationType string `json:"type"`
}
type UpdateNotificationRequest struct {
	Message          string `json:"message,omitempty"`
	NotificationType string `json:"type,omitempty"`
	IsRead           bool   `json:"is_read,omitempty"`
}
type NotificationResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
	Type      string `json:"type"`
	IsRead    bool   `json:"is_read"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
func (h *NotificationHandler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	var req CreateNotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	notificationID, err := h.notificationService.CreateNotification(req.UserID, req.Message, req.NotificationType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := map[string]string{
		"message":         "Notification created successfully",
		"notification_id": notificationID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *NotificationHandler) GetNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notificationID := vars["id"]
	notification, err := h.notificationService.GetNotificationByID(notificationID)
	if err != nil {
		http.Error(w, "Notification not found", http.StatusNotFound)
		return
	}
	response := NotificationResponse{
		ID:        notification.ID.String(),
		UserID:    notification.UserID.String(),
		Message:   notification.Message,
		Type:      notification.Type,
		IsRead:    notification.IsRead,
		CreatedAt: notification.CreatedAt.Format(time.RFC3339),
		UpdatedAt: notification.UpdatedAt.Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *NotificationHandler) GetNotificationsByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}
	notifications, err := h.notificationService.GetNotificationsByUserID(userID)
	if err != nil {
		http.Error(w, "Error fetching notifications", http.StatusInternalServerError)
		return
	}
	var response []NotificationResponse
	for _, notification := range notifications {
		response = append(response, NotificationResponse{
			ID:        notification.ID.String(),
			UserID:    notification.UserID.String(),
			Message:   notification.Message,
			Type:      notification.Type,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt.Format(time.RFC3339),
			UpdatedAt: notification.UpdatedAt.Format(time.RFC3339),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *NotificationHandler) UpdateNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notificationID := vars["id"]
	var req UpdateNotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := h.notificationService.UpdateNotification(notificationID, req.Message, req.NotificationType, req.IsRead)
	if err != nil {
		http.Error(w, "Error updating notification", http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"message": "Notification updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *NotificationHandler) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notificationID := vars["id"]
	err := h.notificationService.DeleteNotification(notificationID)
	if err != nil {
		http.Error(w, "Error deleting notification", http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"message": "Notification deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

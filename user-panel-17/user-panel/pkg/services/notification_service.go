package services

import (
	"net/http"
	"user-panel/internal/db"

	"github.com/gin-gonic/gin"
)

// GetNotifications retrieves all notifications for the logged-in user
func GetNotifications(c *gin.Context) {
	userID := c.GetString("userID")

	rows, err := db.DB.Query("SELECT id, message, status, created_at FROM notifications WHERE user_id = $1 ORDER BY created_at DESC", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch notifications"})
		return
	}
	defer rows.Close()

	var notifications []map[string]interface{}
	for rows.Next() {
		var id, status, message string
		var createdAt string
		if err := rows.Scan(&id, &message, &status, &createdAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse notifications"})
			return
		}
		notifications = append(notifications, gin.H{
			"id":         id,
			"message":    message,
			"status":     status,
			"created_at": createdAt,
		})
	}

	c.JSON(http.StatusOK, notifications)
}

// MarkNotificationAsRead updates the status of a notification to 'READ'
func MarkNotificationAsRead(c *gin.Context) {
	userID := c.GetString("userID")
	notificationID := c.Param("id")

	_, err := db.DB.Exec("UPDATE notifications SET status = 'READ' WHERE id = $1 AND user_id = $2", notificationID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to mark notification as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "notification marked as read"})
}

// DeleteNotification deletes a specific notification
func DeleteNotification(c *gin.Context) {
    userID := c.GetString("userID")
    notificationID := c.Param("id")

    _, err := db.DB.Exec("DELETE FROM notifications WHERE id = $1 AND user_id = $2", notificationID, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete notification"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "notification deleted successfully"})
}

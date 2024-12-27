
package services

import (
	"log"
	"net/http"
	"user-panel/internal/db"

	"github.com/gin-gonic/gin"
)

// LogUserActivity logs user activities into the database
func LogUserActivity(userID, activity, description string) {
	_, err := db.DB.Exec("INSERT INTO user_activities (user_id, activity, description) VALUES ($1, $2, $3)",
		userID, activity, description)
	if err != nil {
		log.Printf("Failed to log user activity: %v", err)
	}
}

// GetUserActivities retrieves all activities for the user
func GetUserActivities(c *gin.Context) {
	userID := c.GetString("userID")

	rows, err := db.DB.Query("SELECT activity, description, created_at FROM user_activities WHERE user_id=$1 ORDER BY created_at DESC", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch activities"})
		return
	}
	defer rows.Close()

	var activities []map[string]interface{}
	for rows.Next() {
		var activity, description, createdAt string
		if err := rows.Scan(&activity, &description, &createdAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse activities"})
			return
		}
		activities = append(activities, gin.H{
			"activity":    activity,
			"description": description,
			"created_at":  createdAt,
		})
	}

	c.JSON(http.StatusOK, activities)
}

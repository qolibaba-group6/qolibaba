
package services

import (
	"log"
	"net/http"
	"user-panel/internal/db"

	"github.com/gin-gonic/gin"
)

// LogTransaction logs user transactions into the database
func LogTransaction(userID, transactionType, description string, amount float64) {
	_, err := db.DB.Exec("INSERT INTO transactions (user_id, transaction_type, amount, description) VALUES ($1, $2, $3, $4)",
		userID, transactionType, amount, description)
	if err != nil {
		log.Printf("Failed to log transaction: %v", err)
	}
}

// GetUserTransactions retrieves all transactions for the user
func GetUserTransactions(c *gin.Context) {
	userID := c.GetString("userID")

	rows, err := db.DB.Query("SELECT transaction_type, amount, description, created_at FROM transactions WHERE user_id=$1 ORDER BY created_at DESC", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transactions"})
		return
	}
	defer rows.Close()

	var transactions []map[string]interface{}
	for rows.Next() {
		var transactionType, description, createdAt string
		var amount float64
		if err := rows.Scan(&transactionType, &amount, &description, &createdAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse transactions"})
			return
		}
		transactions = append(transactions, gin.H{
			"transaction_type": transactionType,
			"amount":           amount,
			"description":      description,
			"created_at":       createdAt,
		})
	}

	c.JSON(http.StatusOK, transactions)
}

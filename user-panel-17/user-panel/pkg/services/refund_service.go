
func UpdateRefundStatus(c *gin.Context) {
    refundID := c.Param("id")
    var input struct {
        Status string `json:"status" binding:"required"` // 'APPROVED' or 'REJECTED'
    }

    if err := c.ShouldBindJSON(&input); err != nil || (input.Status != "APPROVED" && input.Status != "REJECTED") {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    // If status is APPROVED, ensure sufficient balance exists
    if input.Status == "APPROVED" {
        var userID string
        var amount float64
        err := db.DB.QueryRow("SELECT user_id, amount FROM refunds WHERE id=$1", refundID).Scan(&userID, &amount)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch refund details"})
            return
        }

        var balance float64
        err = db.DB.QueryRow("SELECT balance FROM wallets WHERE user_id=$1", userID).Scan(&balance)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch wallet balance"})
            return
        }

        if balance < amount {
            c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient wallet balance for refund"})
            return
        }

        // Deduct the refund amount from wallet
        _, err = db.DB.Exec("UPDATE wallets SET balance = balance - $1 WHERE user_id = $2", amount, userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to deduct refund amount from wallet"})
            return
        }
    }

    _, err := db.DB.Exec("UPDATE refunds SET status=$1, updated_at=NOW() WHERE id=$2", input.Status, refundID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update refund status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "refund status updated successfully"})
}


package main

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

// DB holds the database connection
var DB *sql.DB

// InitializeDB initializes the database connection
func InitializeDB() {
    var err error
    connStr := "postgres://user:password@localhost:5432/walletdb?sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    // Create tables if they don't exist
    _, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS wallets (
            id SERIAL PRIMARY KEY,
            user_id INT NOT NULL,
            balance NUMERIC(15, 2) DEFAULT 0.00
        );

        CREATE TABLE IF NOT EXISTS transactions (
            id SERIAL PRIMARY KEY,
            wallet_id INT NOT NULL,
            type VARCHAR(50) NOT NULL,
            amount NUMERIC(15, 2) NOT NULL,
            created_at TIMESTAMP DEFAULT NOW()
        )
    `)
    if err != nil {
        log.Fatalf("Could not create wallets or transactions table: %v", err)
    }
    log.Println("Database connected and tables created (if not exist).")
}

// DepositDB updates the wallet balance and logs the transaction
func DepositDB(walletID int, amount float64) error {
    tx, err := DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec("UPDATE wallets SET balance = balance + $1 WHERE id = $2", amount, walletID)
    if err != nil {
        tx.Rollback()
        return err
    }

    _, err = tx.Exec("INSERT INTO transactions (wallet_id, type, amount) VALUES ($1, $2, $3)", walletID, "deposit", amount)
    if err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}

// WithdrawDB updates the wallet balance and logs the transaction
func WithdrawDB(walletID int, amount float64) error {
    tx, err := DB.Begin()
    if err != nil {
        return err
    }

    var currentBalance float64
    err = tx.QueryRow("SELECT balance FROM wallets WHERE id = $1", walletID).Scan(&currentBalance)
    if err != nil {
        tx.Rollback()
        return err
    }

    if currentBalance < amount {
        tx.Rollback()
        return sql.ErrNoRows // Insufficient funds
    }

    _, err = tx.Exec("UPDATE wallets SET balance = balance - $1 WHERE id = $2", amount, walletID)
    if err != nil {
        tx.Rollback()
        return err
    }

    _, err = tx.Exec("INSERT INTO transactions (wallet_id, type, amount) VALUES ($1, $2, $3)", walletID, "withdrawal", amount)
    if err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}

// GetTransactionHistoryDB retrieves the transaction history for a wallet
func GetTransactionHistoryDB(walletID int) ([]map[string]interface{}, error) {
    rows, err := DB.Query("SELECT type, amount, created_at FROM transactions WHERE wallet_id = $1 ORDER BY created_at DESC", walletID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var history []map[string]interface{}
    for rows.Next() {
        var transactionType string
        var amount float64
        var createdAt string
        if err := rows.Scan(&transactionType, &amount, &createdAt); err != nil {
            return nil, err
        }
        history = append(history, map[string]interface{}{
            "type":       transactionType,
            "amount":     amount,
            "created_at": createdAt,
        })
    }
    return history, nil
}

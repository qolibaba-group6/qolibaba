package types

import (
	"gorm.io/gorm"
	"time"
)

// Constants for Wallet Role, Transaction Type, Claim Type, and Status
const (
	WalletRoleUser = "user"
	WalletRoleBank = "bank"

	TransactionTypeDeposit    = "deposit"
	TransactionTypeWithdrawal = "withdrawal"
	TransactionTypePayment    = "payment"
	TransactionTypeRefund     = "refund"

	ClaimTypeHotel     = "hotel"
	ClaimTypeFlight    = "flight"
	ClaimTypeTransport = "transport"
	ClaimTypeOther     = "other"

	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
)

// Wallet Model
type Wallet struct {
	gorm.Model
	UserID       *uint         `gorm:"uniqueIndex;not null" json:"user_id"`
	CardNumber   string        `gorm:"type:varchar(16);not null;unique" json:"card_number"`
	Balance      float64       `gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	Role         string        `gorm:"type:enum('user', 'bank');not null" json:"role"`
	Transactions []Transaction `gorm:"foreignKey:WalletID" json:"transactions"`
}

// Transaction Model
type Transaction struct {
	gorm.Model
	WalletID        uint       `gorm:"not null" json:"wallet_id"`
	Amount          float64    `gorm:"type:decimal(15,2);not null" json:"amount"`
	TransactionType string     `gorm:"type:enum('deposit', 'withdrawal', 'payment', 'refund');not null" json:"transaction_type"`
	Status          string     `gorm:"type:enum('pending', 'completed', 'failed');not null" json:"status"`
	Description     string     `gorm:"type:text" json:"description"`
	CompletedAt     *time.Time `gorm:"default:null" json:"completed_at"`

	Wallet Wallet  `gorm:"foreignKey:WalletID" json:"wallet"`
	Claims []Claim `gorm:"foreignKey:TransactionID" json:"claims"`
}

// Claim Model
type Claim struct {
	gorm.Model
	TransactionID uint       `gorm:"not null" json:"transaction_id"`
	UserID        uint       `gorm:"not null" json:"user_id"`
	Amount        float64    `gorm:"type:decimal(15,2);not null" json:"amount"`
	ClaimType     string     `gorm:"type:enum('hotel', 'flight', 'transport', 'other');not null" json:"claim_type"`
	ClaimDetails  string     `gorm:"type:text" json:"claim_details"`
	ServiceID     uint       `gorm:"not null" json:"service_id"`
	ServiceType   string     `gorm:"type:enum('hotel', 'flight', 'transport');not null" json:"service_type"`
	Status        string     `gorm:"type:enum('pending', 'paid', 'failed');not null" json:"status"`
	CompletedAt   *time.Time `gorm:"default:null" json:"completed_at"`

	Transaction Transaction `gorm:"foreignKey:TransactionID" json:"transaction"`
}

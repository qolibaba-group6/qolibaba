package types

import (
	"github.com/google/uuid"
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
	StatusPaid      = "paid"
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

type Transaction struct {
	gorm.Model
	TrackingID      uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"tracking_id"`
	WalletID        uint       `gorm:"not null" json:"wallet_id"`
	Amount          float64    `gorm:"type:decimal(15,2);not null" json:"amount"`
	TransactionType string     `gorm:"type:enum('deposit', 'withdrawal', 'payment', 'refund');not null" json:"transaction_type"`
	Status          string     `gorm:"type:enum('pending', 'completed', 'failed');not null" json:"status"`
	Description     string     `gorm:"type:text" json:"description"`
	CompletedAt     *time.Time `gorm:"default:null" json:"completed_at"`

	Wallet Wallet  `gorm:"foreignKey:WalletID" json:"wallet"`
	Claims []Claim `gorm:"foreignKey:TransactionID" json:"claims"`
}

type Claim struct {
	gorm.Model
	BuyerUserID  uint       `gorm:"not null" json:"user_id"`
	SellerUserID uint       `gorm:"not null" json:"seller_user_id"`
	Amount       float64    `gorm:"type:decimal(15,2);not null" json:"amount"`
	ClaimType    string     `gorm:"type:enum('hotel', 'flight', 'transport', 'other');not null" json:"claim_type"`
	ClaimDetails string     `gorm:"type:text" json:"claim_details"`
	Status       string     `gorm:"type:enum('pending', 'paid', 'failed');not null" json:"status"`
	CompletedAt  *time.Time `gorm:"default:null" json:"completed_at"`

	Transactions []Transaction `gorm:"foreignKey:ClaimID" json:"transactions"`
}

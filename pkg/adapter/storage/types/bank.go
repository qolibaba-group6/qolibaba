package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type WalletRole string
type TransactionType string
type Status string

const (
	WalletRoleUser WalletRole = "user"
	WalletRoleBank WalletRole = "bank"

	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeWithdrawal TransactionType = "withdrawal"

	StatusPending   Status = "pending"
	StatusCompleted Status = "completed"
	StatusPaid      Status = "paid"
	StatusFailed    Status = "failed"
)

// Wallet Model
type Wallet struct {
	gorm.Model
	UserID       *uint         `gorm:"uniqueIndex;not null" json:"user_id"`
	CardNumber   string        `gorm:"type:varchar(16);not null;unique" json:"card_number"`
	Balance      float64       `gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	Role         string        `gorm:"type:varchar(10);not null" json:"role"`
	Transactions []Transaction `gorm:"foreignKey:WalletID" json:"transactions"`
}

type Transaction struct {
	gorm.Model
	TrackingID      uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"tracking_id"`
	WalletID        uint       `gorm:"not null" json:"wallet_id"`
	Amount          float64    `gorm:"type:decimal(15,2);not null" json:"amount"`
	TransactionType string     `gorm:"type:varchar(15);not null" json:"transaction_type"`
	Status          string     `gorm:"type:varchar(10);not null" json:"status"`
	Description     string     `gorm:"type:text" json:"description"`
	CompletedAt     *time.Time `gorm:"default:null" json:"completed_at"`
	ClaimID         uint       `gorm:"not null" json:"claim_id"`

	Wallet Wallet `gorm:"foreignKey:WalletID" json:"wallet"`
	Claim  Claim  `gorm:"foreignKey:ClaimID" json:"claim"`
}

type Claim struct {
	gorm.Model
	BuyerUserID  uint       `gorm:"not null" json:"user_id"`
	SellerUserID uint       `gorm:"not null" json:"seller_user_id"`
	Amount       float64    `gorm:"type:decimal(15,2);not null" json:"amount"`
	ClaimType    string     `gorm:"type:varchar(15);not null" json:"claim_type"`
	ClaimDetails string     `gorm:"type:text" json:"claim_details"`
	Status       string     `gorm:"type:varchar(10);not null" json:"status"`
	CompletedAt  *time.Time `gorm:"default:null" json:"completed_at"`

	Transactions []Transaction `gorm:"foreignKey:ClaimID" json:"transactions"`
}

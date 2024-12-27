package types

import "time"

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
	ID           uint          `gorm:"primaryKey"`
	UserID       *uint         `gorm:"uniqueIndex;not null"`
	CardNumber   string        `gorm:"type:varchar(16);not null;unique"`
	Balance      float64       `gorm:"type:decimal(15,2);default:0.00"`
	Role         string        `gorm:"type:enum('user', 'bank');not null"`
	Transactions []Transaction `gorm:"foreignKey:WalletID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Transaction Model
type Transaction struct {
	ID              uint    `gorm:"primaryKey"`
	WalletID        uint    `gorm:"not null"`
	Amount          float64 `gorm:"type:decimal(15,2);not null"`
	TransactionType string  `gorm:"type:enum('deposit', 'withdrawal', 'payment', 'refund');not null"`
	Status          string  `gorm:"type:enum('pending', 'completed', 'failed');not null"`
	Description     string  `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CompletedAt     *time.Time `gorm:"default:null"`

	Wallet Wallet  `gorm:"foreignKey:WalletID"`
	Claims []Claim `gorm:"foreignKey:TransactionID"`
}

// Claim Model
type Claim struct {
	ID            uint    `gorm:"primaryKey"`
	TransactionID uint    `gorm:"not null"`
	UserID        uint    `gorm:"not null"`
	Amount        float64 `gorm:"type:decimal(15,2);not null"`
	ClaimType     string  `gorm:"type:enum('hotel', 'flight', 'transport', 'other');not null"`
	ClaimDetails  string  `gorm:"type:text"`
	ServiceID     uint    `gorm:"not null"`
	ServiceType   string  `gorm:"type:enum('hotel', 'flight', 'transport');not null"`
	Status        string  `gorm:"type:enum('pending', 'paid', 'failed');not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CompletedAt   *time.Time `gorm:"default:null"`

	Transaction Transaction `gorm:"foreignKey:TransactionID"`
}

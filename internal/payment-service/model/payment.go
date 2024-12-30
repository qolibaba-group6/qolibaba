package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Amount        float64   `gorm:"not null" json:"amount"`
	Status        string    `gorm:"not null" json:"status"` 
	TransactionID string    `gorm:"unique;not null" json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}

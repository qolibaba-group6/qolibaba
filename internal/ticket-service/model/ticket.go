package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Type         string    `gorm:"not null" json:"type"`
	Price        float64   `gorm:"not null" json:"price"`
	Status       string    `gorm:"not null" json:"status"`
	ReturnPolicy string    `gorm:"type:text" json:"return_policy"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

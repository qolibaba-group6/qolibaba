// internal/travel_agencies/domain/trip.go
package domain

import (
	"time"

	"github.com/google/uuid"
)

type Trip struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;"`
	CompanyID     uuid.UUID `gorm:"type:uuid;not null"`
	Type          string    `gorm:"not null"` // مثال: دریایی، قطار، اتوبوس، هوایی
	Origin        string    `gorm:"not null"`
	Destination   string    `gorm:"not null"`
	DepartureTime time.Time `gorm:"not null"`
	ReleaseDate   time.Time `gorm:"not null"`
	Tariff        float64   `gorm:"not null"`
	Status        string    `gorm:"not null"` // مثال: pending, confirmed, canceled
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// internal/travel_agencies/domain/company.go
package domain

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"unique;not null"`
	Owner     string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Trips     []Trip `gorm:"foreignKey:CompanyID"`
}

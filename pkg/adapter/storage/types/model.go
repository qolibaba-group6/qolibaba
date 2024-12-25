package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// The Model struct is a replacement for gorm.Model. it use uuid.UUID as primary key
type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

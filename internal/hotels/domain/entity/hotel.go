package entity

import (
	"time"
)

// Enums for RoomType and BookingStatus
const (
	RoomTypeSingle = "single"
	RoomTypeDouble = "double"
	RoomTypeSuite  = "suite"

	BookingStatusPending   = "pending"
	BookingStatusConfirmed = "confirmed"
	BookingStatusCompleted = "completed"
)

// Hotel model
type Hotel struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(100);not null" validate:"required"`
	Location    string  `gorm:"type:varchar(255);not null" validate:"required"`
	PhoneNumber string  `gorm:"type:varchar(15);not null" validate:"required,e164"`
	Email       *string `gorm:"type:varchar(100);unique" validate:"omitempty,email"`
	Website     *string `gorm:"type:varchar(255)" validate:"omitempty,url"`
	Rooms       []Room  `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Room model
type Room struct {
	ID                uint      `gorm:"primaryKey"`
	HotelID           uint      `gorm:"not null"`
	Type              string    `gorm:"type:enum('single','double','suite');not null" validate:"required,oneof=single double suite"`
	Price             float64   `gorm:"type:decimal(10,2);not null" validate:"required,gt=0"`
	Capacity          int       `gorm:"not null" validate:"required,gt=0"`
	Features          string    `gorm:"type:text"`
	Duration          string    `gorm:"type:enum('12 hours','24 hours');not null" validate:"required,oneof=12 hours 24 hours"`
	PublicReleaseDate time.Time `gorm:"not null" validate:"required"`
	AgencyReleaseDate time.Time `gorm:"not null" validate:"required"`
	Bookings          []Booking `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// Booking model
type Booking struct {
	ID                 uint       `gorm:"primaryKey"`
	RoomID             uint       `gorm:"not null"`
	UserID             uint       `gorm:"not null"`
	StartTime          time.Time  `gorm:"not null" validate:"required"`
	EndTime            time.Time  `gorm:"not null" validate:"required,gtfield=StartTime"`
	TotalPrice         float64    `gorm:"type:decimal(10,2);not null" validate:"required,gt=0"`
	Status             string     `gorm:"type:enum('pending','confirmed','completed');not null" validate:"required,oneof=pending confirmed completed"`
	Confirmed          bool       `gorm:"not null" validate:"required"`
	DateOfConfirmation *time.Time `gorm:"default:null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

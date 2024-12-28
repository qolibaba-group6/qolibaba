package types

import (
	"gorm.io/gorm"
	"time"
)

type RoomType string

const (
	RoomTypeSingle RoomType = "single"
	RoomTypeDouble RoomType = "double"
	RoomTypeSuite  RoomType = "suite"
)

type DurationType string

const (
	Duration12Hours DurationType = "12 hours"
	Duration24Hours DurationType = "24 hours"
)

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCompleted BookingStatus = "completed"
)

type Hotel struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null" validate:"required"`
	Location    string  `gorm:"type:varchar(255);not null" validate:"required"`
	PhoneNumber string  `gorm:"type:varchar(15);not null" validate:"required,e164"`
	Email       *string `gorm:"type:varchar(100);unique" validate:"omitempty,email"`
	Website     *string `gorm:"type:varchar(255)" validate:"omitempty,url"`
	Rooms       []Room  `gorm:"constraint:OnDelete:CASCADE"`
}

type Room struct {
	gorm.Model
	HotelID           uint         `gorm:"not null"`
	Type              RoomType     `gorm:"type:room_type;not null" validate:"required,oneof=single double suite"`
	Price             float64      `gorm:"type:decimal(10,2);not null" validate:"required,gt=0"`
	Capacity          int          `gorm:"not null" validate:"required,gt=0"`
	Features          string       `gorm:"type:text"`
	Duration          DurationType `gorm:"type:duration_type;not null" validate:"required,oneof=12 hours 24 hours"`
	PublicReleaseDate time.Time    `gorm:"not null" validate:"required"`
	AgencyReleaseDate time.Time    `gorm:"not null" validate:"required"`
	Bookings          []Booking    `gorm:"constraint:OnDelete:CASCADE"`
}

type Booking struct {
	gorm.Model
	RoomID             uint          `gorm:"not null"`
	UserID             uint          `gorm:"not null"`
	StartTime          time.Time     `gorm:"not null" validate:"required"`
	EndTime            time.Time     `gorm:"not null" validate:"required,gtfield=StartTime"`
	TotalPrice         float64       `gorm:"type:decimal(10,2);not null" validate:"required,gt=0"`
	Status             BookingStatus `gorm:"type:booking_status;not null" validate:"required,oneof=pending confirmed completed"`
	Confirmed          bool          `gorm:"not null" validate:"required"`
	DateOfConfirmation *time.Time    `gorm:"default:null"`
	IsReferred         *uint         `gorm:"type:bigint;default:null"`
	DeletedAt          *time.Time    `gorm:"index"`
}

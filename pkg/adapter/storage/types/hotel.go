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

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCompleted BookingStatus = "completed"
)

type RoomStatus string

const (
	RoomStatusFree   RoomStatus = "free"
	RoomStatusBooked RoomStatus = "booked"
)

type Hotel struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null" validate:"required" json:"name"`
	Location    string  `gorm:"type:varchar(255);not null" validate:"required" json:"location"`
	PhoneNumber string  `gorm:"type:varchar(15);not null" validate:"required,e164" json:"phone_number"`
	Email       *string `gorm:"type:varchar(100);unique" validate:"omitempty,email" json:"email"`
	Website     *string `gorm:"type:varchar(255)" validate:"omitempty,url" json:"website"`
	Rooms       []Room  `gorm:"constraint:OnDelete:CASCADE" json:"rooms"`
}

type Room struct {
	gorm.Model
	HotelID           uint       `gorm:"not null" json:"hotel_id"`
	Type              RoomType   `gorm:"type:room_type;not null" validate:"required,oneof=single double suite" json:"type"`
	GeneralPrice      float64    `gorm:"type:decimal(10,2);not null" json:"general_price"`
	TourPrice         float64    `gorm:"type:decimal(10,2);not null" json:"tour_price"`
	Capacity          int        `gorm:"not null" validate:"required,gt=0" json:"capacity"`
	Features          string     `gorm:"type:text" json:"features"`
	Duration          int        `gorm:"not null" validate:"required,oneof=12 24" json:"duration"`
	PublicReleaseDate time.Time  `gorm:"not null" validate:"required" json:"public_release_date"`
	AgencyReleaseDate time.Time  `gorm:"not null" validate:"required" json:"agency_release_date"`
	Bookings          []Booking  `gorm:"constraint:OnDelete:CASCADE" json:"bookings"`
	Status            RoomStatus `gorm:"type:room_status;not null;default:'free'" validate:"required,oneof=free booked" json:"status"`
}

type Booking struct {
	gorm.Model
	RoomID             uint          `gorm:"not null" json:"room_id"`
	UserID             uint          `gorm:"not null" json:"user_id"`
	StartTime          time.Time     `gorm:"not null" validate:"required" json:"start_time"`
	EndTime            time.Time     `gorm:"not null" validate:"required,gtfield=StartTime" json:"end_time"`
	TotalPrice         *float64      `gorm:"type:decimal(10,2);not null" validate:"required,gt=0" json:"total_price"`
	Status             BookingStatus `gorm:"type:booking_status;not null" validate:"required,oneof=pending confirmed completed" json:"status"`
	Confirmed          bool          `gorm:"not null;default:false" json:"confirmed"`
	DateOfConfirmation *time.Time    `gorm:"default:null" json:"date_of_confirmation"`
	IsReferred         *uint         `gorm:"type:bigint;default:null" json:"is_referred"`
	ClaimID            *uint         `gorm:"default:null" json:"claim_id"`
	DeletedAt          *time.Time    `gorm:"index" json:"deleted_at"`
}

package types

import (
	"gorm.io/gorm"
	"time"
)

type Tour struct {
	gorm.Model
	AgencyID               uint       `gorm:"not null" json:"agency_id" validate:"required"`
	Title                  string     `gorm:"type:varchar(255);not null" json:"title" validate:"required"`
	Description            string     `gorm:"type:text" json:"description" validate:"required"`
	HotelID                uint       `gorm:"not null" json:"hotel_id" validate:"required"`
	GoingVehicleID         uint       `gorm:"not null" json:"going_vehicle_id" validate:"required"`
	ReturnVehicleID        uint       `gorm:"not null" json:"return_vehicle_id" validate:"required"`
	StartDate              time.Time  `gorm:"not null" json:"start_date" validate:"required"`
	EndDate                time.Time  `gorm:"not null" json:"end_date" validate:"required"`
	Price                  float64    `gorm:"type:decimal(10,2);not null" json:"price" validate:"required,gt=0"`
	PublishedDate          *time.Time `gorm:"default:null" json:"published_date"`
	ReleaseTime            *time.Time `gorm:"default:null" json:"release_time"`
	HotelPrice             float64    `gorm:"type:decimal(10,2);not null" json:"hotel_price" validate:"required,gt=0"`
	GoingTransferPrice     float64    `gorm:"type:decimal(10,2);not null" json:"going_transfer_price" validate:"required,gt=0"`
	ReturningTransferPrice float64    `gorm:"type:decimal(10,2);not null" json:"returning_transfer_price" validate:"required,gt=0"`
	TotalPrice             float64    `gorm:"type:decimal(10,2);not null" json:"total_price" validate:"required,gt=0"`
}

type TravelAgency struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	Email       string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Phone       string `gorm:"type:varchar(15)" json:"phone"`
}

type TourBooking struct {
	gorm.Model
	TourID                  uint       `gorm:"not null;foreignKey:ID;references:ID" json:"tour_id"`
	UserID                  uint       `gorm:"not null" json:"user_id"`
	BookingDate             time.Time  `gorm:"not null" json:"booking_date"`
	BookingStatus           string     `gorm:"type:varchar(50);not null" json:"status"`
	Confirmed               bool       `gorm:"not null;default:false" json:"confirmed"`
	DateConfirmed           *time.Time `gorm:"default:null" json:"date_confirmed"`
	ClaimID                 *uint      `gorm:"type:varchar(255);default:null" json:"claim_id"`
	GoingTransferVehicleID  uint       `gorm:"not null" json:"going_transfer_vehicle_id"`
	ReturnTransferVehicleID uint       `gorm:"not null" json:"return_transfer_vehicle_id"`
	HotelID                 uint       `gorm:"not null" json:"hotel_id"`
	RoomID                  uint       `gorm:"not null" json:"room_id"`
}

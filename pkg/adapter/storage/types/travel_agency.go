package types

import (
	"gorm.io/gorm"
	"time"
)

type Tour struct {
	gorm.Model
	AgencyID        uint       `gorm:"not null" json:"agency_id"`
	Title           string     `gorm:"type:varchar(255);not null" json:"title"`
	Description     string     `gorm:"type:text" json:"description"`
	HotelID         uint       `gorm:"not null" json:"hotel_id"`
	GoingVehicleID  uint       `gorm:"not null" json:"going_vehicle_id"`
	ReturnVehicleID uint       `gorm:"not null" json:"return_vehicle_id"`
	StartDate       time.Time  `gorm:"not null" json:"start_date"`
	EndDate         time.Time  `gorm:"not null" json:"end_date"`
	Price           float64    `gorm:"type:decimal(10,2);not null" json:"price"`
	PublishedDate   *time.Time `gorm:"default:null" json:"published_date"`
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
	TourID                  uint       `gorm:"not null" json:"tour_id"`
	UserID                  uint       `gorm:"not null" json:"user_id"`
	BookingDate             time.Time  `gorm:"not null" json:"booking_date"`
	TotalPrice              float64    `gorm:"type:decimal(10,2);not null" json:"total_price"`
	BookingStatus           string     `gorm:"type:varchar(50);not null" json:"status"`
	Confirmed               bool       `gorm:"not null;default:false" json:"confirmed"`
	DateConfirmed           *time.Time `gorm:"default:null" json:"date_confirmed"`
	ClaimID                 *string    `gorm:"type:varchar(255);default:null" json:"claim_id"`
	GoingTransferVehicleID  uint       `gorm:"not null" json:"going_transfer_vehicle_id"`
	ReturnTransferVehicleID uint       `gorm:"not null" json:"return_transfer_vehicle_id"`
	HotelID                 uint       `gorm:"not null" json:"hotel_id"`
	RoomID                  uint       `gorm:"not null" json:"room_id"`
}

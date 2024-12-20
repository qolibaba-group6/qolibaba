package models

import (
	"time"
)

// TravelAgency model
// Represents a travel agency that can offer tours and earn revenue from ticket and hotel sales.
type TravelAgency struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(100);not null" validate:"required"`
	Address     string  `gorm:"type:varchar(255);not null" validate:"required"`
	PhoneNumber string  `gorm:"type:varchar(15);not null" validate:"required,e164"`
	Email       *string `gorm:"type:varchar(100);unique" validate:"omitempty,email"`
	Website     *string `gorm:"type:varchar(255)" validate:"omitempty,url"`
	Tours       []Tour  `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Tour model
// Represents a tour offered by a travel agency, including tickets and hotel stays.
type Tour struct {
	ID                  uint      `gorm:"primaryKey"`
	TravelAgencyID      uint      `gorm:"not null"`
	Name                string    `gorm:"type:varchar(100);not null" validate:"required"`
	Description         string    `gorm:"type:text"`
	GoingTravelID       uint      `gorm:"not null"`  // References a travel option (e.g., flight/train)
	ReturningTravelID   uint      `gorm:"not null"`  // References a travel option (e.g., flight/train)
	HotelID             uint      `gorm:"not null"`  // References the hotel for the tour
	GoingTravelData     JSON      `gorm:"type:json"` // Stores JSON data for the going travel
	ReturningTravelData JSON      `gorm:"type:json"` // Stores JSON data for the returning travel
	HotelData           JSON      `gorm:"type:json"` // Stores JSON data for the hotel
	Price               float64   `gorm:"type:decimal(10,2);not null" validate:"required,gt=0"`
	ReleaseDate         time.Time `gorm:"not null" validate:"required"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

// TravelClaim model
// Represents a claim for revenue distribution related to a tour (e.g., travel tickets, hotels).
type TravelClaim struct {
	ID        uint    `gorm:"primaryKey"`
	TourID    uint    `gorm:"not null"`
	Type      string  `gorm:"type:enum('ticket','hotel');not null" validate:"required,oneof=ticket hotel"`
	Amount    float64 `gorm:"type:decimal(10,2);not null" validate:"required,gt=0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

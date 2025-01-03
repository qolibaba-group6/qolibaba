package storage

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"qolibaba/internal/travel_agencies/port"
	"qolibaba/pkg/adapter/storage/types"
	"time"
)

type TravelAgencyRepository struct {
	db *gorm.DB
}

func NewTravelAgencyRepository(db *gorm.DB) port.Repo {
	return &TravelAgencyRepository{
		db: db,
	}
}

// Create inserts a new travel agency into the database
func (r *TravelAgencyRepository) Create(agency *types.TravelAgency) error {
	if err := r.db.Create(agency).Error; err != nil {
		return err
	}
	return nil
}

// FindByEmail retrieves a travel agency by email
func (r *TravelAgencyRepository) FindByEmail(email string) (*types.TravelAgency, error) {
	var agency types.TravelAgency
	if err := r.db.Where("email = ?", email).First(&agency).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &agency, nil
}

func (r *TravelAgencyRepository) SaveTour(tour *types.Tour) (*types.Tour, error) {
	if err := r.db.Create(tour).Error; err != nil {
		log.Printf("Error saving tour: %v", err)
		return nil, err
	}
	return tour, nil
}

func (r *TravelAgencyRepository) GetTourByID(tourID uint) (*types.Tour, error) {
	var tour types.Tour
	err := r.db.First(&tour, tourID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("tour with ID %d not found", tourID)
		}
		return nil, fmt.Errorf("error retrieving tour by ID: %v", err)
	}
	return &tour, nil
}

func (r *TravelAgencyRepository) CreateTourBooking(tourBooking *types.TourBooking) (*types.TourBooking, error) {
	err := r.db.Create(tourBooking).Error
	if err != nil {
		return nil, fmt.Errorf("error creating tour booking: %v", err)
	}
	return tourBooking, nil
}

func (r *TravelAgencyRepository) UpdateTourBooking(tourBooking *types.TourBooking) (*types.TourBooking, error) {
	err := r.db.Save(tourBooking).Error
	if err != nil {
		return nil, fmt.Errorf("error updating tour booking: %v", err)
	}
	return tourBooking, nil
}

func (r *TravelAgencyRepository) ConfirmBooking(bookingID uint) (*types.TourBooking, error) {
	var booking types.TourBooking
	if err := r.db.First(&booking, bookingID).Error; err != nil {
		return nil, fmt.Errorf("error fetching booking with ID %d: %v", bookingID, err)
	}

	var tour types.Tour
	if err := r.db.First(&tour, booking.TourID).Error; err != nil {
		return nil, fmt.Errorf("error fetching tour for booking %d: %v", bookingID, err)
	}

	if tour.EndDate.After(time.Now()) {
		return nil, fmt.Errorf("the tour has not ended yet, cannot confirm the booking")
	}

	booking.Confirmed = true
	booking.BookingStatus = string(types.BookingStatusCompleted)
	confirmedAt := time.Now()
	booking.DateConfirmed = &confirmedAt

	if err := r.db.Save(&booking).Error; err != nil {
		return nil, fmt.Errorf("error confirming booking with ID %d: %v", bookingID, err)
	}

	return &booking, nil
}

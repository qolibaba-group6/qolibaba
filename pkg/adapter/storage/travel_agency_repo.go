package storage

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"qolibaba/pkg/adapter/storage/types"
)

type TravelAgencyRepository struct {
	db *gorm.DB
}

func NewTravelAgencyRepository(db *gorm.DB) *TravelAgencyRepository {
	return &TravelAgencyRepository{db: db}
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

func (r *TravelAgencyRepository) SaveTour(tour *types.TourBooking) (*types.TourBooking, error) {
	if err := r.db.Create(tour).Error; err != nil {
		log.Printf("Error saving tour: %v", err)
		return nil, err
	}
	return tour, nil
}

// internal/travel_agencies/service/trip_service.go
package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/domain"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/port"
)

var (
	ErrTripNotFound = errors.New("trip not found")
	ErrInvalidInput = errors.New("invalid input")
)

type TripService struct {
	Repo           port.TripRepository
	CompanyService *CompanyService
}

func NewTripService(repo port.TripRepository, companyService *CompanyService) *TripService {
	return &TripService{
		Repo:           repo,
		CompanyService: companyService,
	}
}

// CreateTrip creates a new trip. It ensures the associated company exists.
func (s *TripService) CreateTrip(companyIDStr, tripType, origin, destination, departureTimeStr, releaseDateStr string, tariff float64, status string) (*domain.Trip, error) {
	// Parse and validate company ID
	companyID, err := uuid.Parse(companyIDStr)
	if err != nil {
		return nil, ErrInvalidInput
	}

	// Check if the company exists
	_, err = s.CompanyService.GetCompany(companyID)
	if err != nil {
		return nil, err
	}

	// Parse and validate departure time
	departureTime, err := time.Parse(time.RFC3339, departureTimeStr)
	if err != nil {
		return nil, ErrInvalidInput
	}

	// Parse and validate release date
	releaseDate, err := time.Parse(time.RFC3339, releaseDateStr)
	if err != nil {
		return nil, ErrInvalidInput
	}

	// Validate tariff
	if tariff < 0 {
		return nil, errors.New("tariff cannot be negative")
	}

	// Create the trip
	trip := &domain.Trip{
		ID:            uuid.New(),
		CompanyID:     companyID,
		Type:          tripType,
		Origin:        origin,
		Destination:   destination,
		DepartureTime: departureTime,
		ReleaseDate:   releaseDate,
		Tariff:        tariff,
		Status:        status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = s.Repo.Create(trip)
	if err != nil {
		return nil, err
	}

	return trip, nil
}

// GetTrip retrieves a trip by its ID.
func (s *TripService) GetTrip(id uuid.UUID) (*domain.Trip, error) {
	trip, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, ErrTripNotFound
	}
	return trip, nil
}

// UpdateTrip updates an existing trip. It ensures the associated company exists.
func (s *TripService) UpdateTrip(idStr, companyIDStr, tripType, origin, destination, departureTimeStr, releaseDateStr string, tariff float64, status string) (*domain.Trip, error) {
	// Parse and validate trip ID
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, ErrInvalidInput
	}

	// Parse and validate company ID
	companyID, err := uuid.Parse(companyIDStr)
	if err != nil {
		return nil, ErrInvalidInput
	}

	// Check if the company exists
	_, err = s.CompanyService.GetCompany(companyID)
	if err != nil {
		return nil, err
	}

	// Retrieve the existing trip
	trip, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, ErrTripNotFound
	}

	// Parse and validate departure time
	departureTime, err := time.Parse(time.RFC3339, departureTimeStr)
	if err != nil {
		return nil, ErrInvalidInput
	}

	// Parse and validate release date
	releaseDate, err := time.Parse(time.RFC3339, releaseDateStr)
	if err != nil {
		return nil, ErrInvalidInput
	}

	// Validate tariff
	if tariff < 0 {
		return nil, errors.New("tariff cannot be negative")
	}

	// Update the trip fields
	trip.CompanyID = companyID
	trip.Type = tripType
	trip.Origin = origin
	trip.Destination = destination
	trip.DepartureTime = departureTime
	trip.ReleaseDate = releaseDate
	trip.Tariff = tariff
	trip.Status = status
	trip.UpdatedAt = time.Now()

	// Save the updated trip
	err = s.Repo.Update(trip)
	if err != nil {
		return nil, err
	}

	return trip, nil
}

// DeleteTrip deletes a trip by its ID.
func (s *TripService) DeleteTrip(id uuid.UUID) error {
	err := s.Repo.Delete(id)
	if err != nil {
		return ErrTripNotFound
	}
	return nil
}

// ListTrips lists all trips for a given company ID.
func (s *TripService) ListTrips(companyIDStr string) ([]domain.Trip, error) {
	if companyIDStr == "" {
		return nil, errors.New("company_id is required")
	}

	// Parse and validate company ID
	companyID, err := uuid.Parse(companyIDStr)
	if err != nil {
		return nil, ErrInvalidInput
	}

	// Check if the company exists
	_, err = s.CompanyService.GetCompany(companyID)
	if err != nil {
		return nil, err
	}

	// Retrieve trips for the company
	trips, err := s.Repo.GetAllByCompany(companyID)
	if err != nil {
		return nil, err
	}

	return trips, nil
}

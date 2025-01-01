// internal/travel_agencies/port/trip.go
package port

import (
	"github.com/google/uuid"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/domain"
)

// TripRepository defines the interface for trip data operations
type TripRepository interface {
	Create(trip *domain.Trip) error
	GetByID(id uuid.UUID) (*domain.Trip, error)
	Update(trip *domain.Trip) error
	Delete(id uuid.UUID) error
	GetAllByCompany(companyID uuid.UUID) ([]domain.Trip, error)
}

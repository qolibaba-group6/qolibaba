// app/travel_agencies/contract.go
package travel_agencies

import (
	"github.com/google/uuid"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/domain"
)

// CompanyRepository defines the interface for company data operations
type CompanyRepository interface {
	Create(company *domain.Company) error
	GetByID(id uuid.UUID) (*domain.Company, error)
	Update(company *domain.Company) error
	Delete(id uuid.UUID) error
	GetAll() ([]domain.Company, error)
}

// TripRepository defines the interface for trip data operations
type TripRepository interface {
	Create(trip *domain.Trip) error
	GetByID(id uuid.UUID) (*domain.Trip, error)
	Update(trip *domain.Trip) error
	Delete(id uuid.UUID) error
	GetAllByCompany(companyID uuid.UUID) ([]domain.Trip, error)
}

// CompanyService defines the interface for company-related business logic
type CompanyService interface {
	CreateCompany(name, owner string) (*domain.Company, error)
	GetCompany(id uuid.UUID) (*domain.Company, error)
	UpdateCompany(id uuid.UUID, name, owner string) (*domain.Company, error)
	DeleteCompany(id uuid.UUID) error
	ListCompanies() ([]domain.Company, error)
}

// TripService defines the interface for trip-related business logic
type TripService interface {
	CreateTrip(companyIDStr, tripType, origin, destination, departureTimeStr, releaseDateStr string, tariff float64, status string) (*domain.Trip, error)
	GetTrip(id uuid.UUID) (*domain.Trip, error)
	UpdateTrip(idStr, companyIDStr, tripType, origin, destination, departureTimeStr, releaseDateStr string, tariff float64, status string) (*domain.Trip, error)
	DeleteTrip(id uuid.UUID) error
	ListTrips(companyIDStr string) ([]domain.Trip, error)
}

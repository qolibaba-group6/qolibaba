// pkg/adapter/storage/travel_agencies_repo.go
package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/domain"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/port"
)

// TravelAgenciesRepo implements both CompanyRepository and TripRepository interfaces
type TravelAgenciesRepo struct {
	DB *gorm.DB
}

// NewTravelAgenciesRepo creates a new instance of TravelAgenciesRepo
func NewTravelAgenciesRepo(db *gorm.DB) (port.CompanyRepository, port.TripRepository) {
	return &TravelAgenciesRepo{DB: db}, &TravelAgenciesRepo{DB: db}
}

// CompanyRepository Methods

// Create inserts a new company into the database
func (r *TravelAgenciesRepo) Create(company *domain.Company) error {
	return r.DB.Create(company).Error
}

// GetByID retrieves a company by its ID
func (r *TravelAgenciesRepo) GetByID(id uuid.UUID) (*domain.Company, error) {
	var company domain.Company
	if err := r.DB.Preload("Trips").First(&company, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

// Update modifies an existing company in the database
func (r *TravelAgenciesRepo) Update(company *domain.Company) error {
	return r.DB.Save(company).Error
}

// Delete removes a company from the database by its ID
func (r *TravelAgenciesRepo) Delete(id uuid.UUID) error {
	return r.DB.Delete(&domain.Company{}, "id = ?", id).Error
}

// GetAll retrieves all companies from the database
func (r *TravelAgenciesRepo) GetAll() ([]domain.Company, error) {
	var companies []domain.Company
	if err := r.DB.Preload("Trips").Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

// TripRepository Methods

// Create inserts a new trip into the database
func (r *TravelAgenciesRepo) Create(trip *domain.Trip) error {
	return r.DB.Create(trip).Error
}

// GetByID retrieves a trip by its ID
func (r *TravelAgenciesRepo) GetByID(id uuid.UUID) (*domain.Trip, error) {
	var trip domain.Trip
	if err := r.DB.First(&trip, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &trip, nil
}

// Update modifies an existing trip in the database
func (r *TravelAgenciesRepo) Update(trip *domain.Trip) error {
	return r.DB.Save(trip).Error
}

// Delete removes a trip from the database by its ID
func (r *TravelAgenciesRepo) Delete(id uuid.UUID) error {
	return r.DB.Delete(&domain.Trip{}, "id = ?", id).Error
}

// GetAllByCompany retrieves all trips associated with a specific company
func (r *TravelAgenciesRepo) GetAllByCompany(companyID uuid.UUID) ([]domain.Trip, error) {
	var trips []domain.Trip
	if err := r.DB.Where("company_id = ?", companyID).Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}

// internal/travel_agencies/port/company.go
package port

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

// internal/travel_agencies/service/company_service.go
package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/domain"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/port"
)

var (
	ErrCompanyNotFound = errors.New("company not found")
)

type CompanyService struct {
	Repo port.CompanyRepository
}

func NewCompanyService(repo port.CompanyRepository) *CompanyService {
	return &CompanyService{Repo: repo}
}

func (s *CompanyService) CreateCompany(name, owner string) (*domain.Company, error) {
	company := &domain.Company{
		ID:        uuid.New(),
		Name:      name,
		Owner:     owner,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.Repo.Create(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (s *CompanyService) GetCompany(id uuid.UUID) (*domain.Company, error) {
	company, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, ErrCompanyNotFound
	}
	return company, nil
}

func (s *CompanyService) UpdateCompany(id uuid.UUID, name, owner string) (*domain.Company, error) {
	company, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, ErrCompanyNotFound
	}
	company.Name = name
	company.Owner = owner
	company.UpdatedAt = time.Now()

	err = s.Repo.Update(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (s *CompanyService) DeleteCompany(id uuid.UUID) error {
	err := s.Repo.Delete(id)
	if err != nil {
		return ErrCompanyNotFound
	}
	return nil
}

func (s *CompanyService) ListCompanies() ([]domain.Company, error) {
	companies, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return companies, nil
}

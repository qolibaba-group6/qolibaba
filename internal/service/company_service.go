package service

import (
	"companies-service/internal/models"
	"companies-service/internal/repository"
	"errors"

	"github.com/google/uuid"
)

type CompanyService interface {
	GetAllCompanies() ([]models.Company, error)
	GetCompanyByID(id uuid.UUID) (*models.Company, error)
	CreateCompany(name, owner string) (*models.Company, error)
	UpdateCompany(id uuid.UUID, name, owner string) (*models.Company, error)
	DeleteCompany(id uuid.UUID) error
}

type companyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) CompanyService {
	return &companyService{repo}
}

func (s *companyService) GetAllCompanies() ([]models.Company, error) {
	return s.repo.GetAll()
}

func (s *companyService) GetCompanyByID(id uuid.UUID) (*models.Company, error) {
	return s.repo.GetByID(id)
}

func (s *companyService) CreateCompany(name, owner string) (*models.Company, error) {
	if name == "" || owner == "" {
		return nil, errors.New("name and owner are required")
	}

	company := &models.Company{
		ID:    uuid.New(),
		Name:  name,
		Owner: owner,
	}

	if err := s.repo.Create(company); err != nil {
		return nil, err
	}

	return company, nil
}

func (s *companyService) UpdateCompany(id uuid.UUID, name, owner string) (*models.Company, error) {
	company, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		company.Name = name
	}

	if owner != "" {
		company.Owner = owner
	}

	if err := s.repo.Update(company); err != nil {
		return nil, err
	}

	return company, nil
}

func (s *companyService) DeleteCompany(id uuid.UUID) error {
	return s.repo.Delete(id)
}

// internal/travel_agencies/service/company_service_test.go
package service

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/domain"
)

// MockCompanyRepository is a mock implementation of CompanyRepository
type MockCompanyRepository struct {
	mock.Mock
}

func (m *MockCompanyRepository) Create(company *domain.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func (m *MockCompanyRepository) GetByID(id uuid.UUID) (*domain.Company, error) {
	args := m.Called(id)
	if company, ok := args.Get(0).(*domain.Company); ok {
		return company, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCompanyRepository) Update(company *domain.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func (m *MockCompanyRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCompanyRepository) GetAll() ([]domain.Company, error) {
	args := m.Called()
	if companies, ok := args.Get(0).([]domain.Company); ok {
		return companies, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateCompany(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	name := "Test Company"
	owner := "John Doe"

	mockRepo.On("Create", mock.AnythingOfType("*domain.Company")).Return(nil)

	company, err := service.CreateCompany(name, owner)

	assert.NoError(t, err)
	assert.Equal(t, name, company.Name)
	assert.Equal(t, owner, company.Owner)
	assert.NotEqual(t, uuid.Nil, company.ID)
	assert.WithinDuration(t, time.Now(), company.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), company.UpdatedAt, time.Second)

	mockRepo.AssertExpectations(t)
}

func TestGetCompany_Success(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()
	company := &domain.Company{
		ID:        id,
		Name:      "Test Company",
		Owner:     "John Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetByID", id).Return(company, nil)

	result, err := service.GetCompany(id)

	assert.NoError(t, err)
	assert.Equal(t, company, result)

	mockRepo.AssertExpectations(t)
}

func TestGetCompany_NotFound(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()

	mockRepo.On("GetByID", id).Return(nil, errors.New("not found"))

	result, err := service.GetCompany(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, ErrCompanyNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateCompany_Success(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()
	existingCompany := &domain.Company{
		ID:        id,
		Name:      "Old Name",
		Owner:     "Old Owner",
		CreatedAt: time.Now().Add(-time.Hour),
		UpdatedAt: time.Now().Add(-time.Hour),
	}

	mockRepo.On("GetByID", id).Return(existingCompany, nil)
	mockRepo.On("Update", existingCompany).Return(nil)

	newName := "New Name"
	newOwner := "New Owner"

	updatedCompany, err := service.UpdateCompany(id, newName, newOwner)

	assert.NoError(t, err)
	assert.Equal(t, newName, updatedCompany.Name)
	assert.Equal(t, newOwner, updatedCompany.Owner)
	assert.WithinDuration(t, time.Now(), updatedCompany.UpdatedAt, time.Second)

	mockRepo.AssertExpectations(t)
}

func TestUpdateCompany_NotFound(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()

	mockRepo.On("GetByID", id).Return(nil, errors.New("not found"))

	updatedCompany, err := service.UpdateCompany(id, "New Name", "New Owner")

	assert.Error(t, err)
	assert.Nil(t, updatedCompany)
	assert.Equal(t, ErrCompanyNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteCompany_Success(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()

	mockRepo.On("Delete", id).Return(nil)

	err := service.DeleteCompany(id)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteCompany_NotFound(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()

	mockRepo.On("Delete", id).Return(errors.New("not found"))

	err := service.DeleteCompany(id)

	assert.Error(t, err)
	assert.Equal(t, ErrCompanyNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestListCompanies(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	companies := []domain.Company{
		{
			ID:        uuid.New(),
			Name:      "Company One",
			Owner:     "Owner One",
			CreatedAt: time.Now().Add(-2 * time.Hour),
			UpdatedAt: time.Now().Add(-time.Hour),
		},
		{
			ID:        uuid.New(),
			Name:      "Company Two",
			Owner:     "Owner Two",
			CreatedAt: time.Now().Add(-3 * time.Hour),
			UpdatedAt: time.Now().Add(-time.Hour),
		},
	}

	mockRepo.On("GetAll").Return(companies, nil)

	result, err := service.ListCompanies()

	assert.NoError(t, err)
	assert.Equal(t, companies, result)

	mockRepo.AssertExpectations(t)
}

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

// MockCompanyRepository is a mock implementation of CompanyRepository interface
type MockCompanyRepository struct {
	mock.Mock
}

// Create mocks the Create method of CompanyRepository
func (m *MockCompanyRepository) Create(company *domain.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

// GetByID mocks the GetByID method of CompanyRepository
func (m *MockCompanyRepository) GetByID(id uuid.UUID) (*domain.Company, error) {
	args := m.Called(id)
	if company, ok := args.Get(0).(*domain.Company); ok {
		return company, args.Error(1)
	}
	return nil, args.Error(1)
}

// Update mocks the Update method of CompanyRepository
func (m *MockCompanyRepository) Update(company *domain.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

// Delete mocks the Delete method of CompanyRepository
func (m *MockCompanyRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetAll mocks the GetAll method of CompanyRepository
func (m *MockCompanyRepository) GetAll() ([]domain.Company, error) {
	args := m.Called()
	if companies, ok := args.Get(0).([]domain.Company); ok {
		return companies, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateCompany(t *testing.T) {
	// Arrange
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	name := "Test Company"
	owner := "John Doe"
	expectedCompany := &domain.Company{
		ID:        uuid.New(),
		Name:      name,
		Owner:     owner,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Expectation: Create method is called with a company having the specified name and owner
	mockRepo.On("Create", mock.AnythingOfType("*domain.Company")).Return(nil).Run(func(args mock.Arguments) {
		company := args.Get(0).(*domain.Company)
		company.ID = expectedCompany.ID
		company.CreatedAt = expectedCompany.CreatedAt
		company.UpdatedAt = expectedCompany.UpdatedAt
	})

	// Act
	company, err := service.CreateCompany(name, owner)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, name, company.Name)
	assert.Equal(t, owner, company.Owner)
	assert.Equal(t, expectedCompany.ID, company.ID)
	assert.WithinDuration(t, expectedCompany.CreatedAt, company.CreatedAt, time.Second)
	assert.WithinDuration(t, expectedCompany.UpdatedAt, company.UpdatedAt, time.Second)

	mockRepo.AssertExpectations(t)
}

func TestGetCompany_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()
	expectedCompany := &domain.Company{
		ID:        id,
		Name:      "Test Company",
		Owner:     "John Doe",
		CreatedAt: time.Now().Add(-time.Hour),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetByID", id).Return(expectedCompany, nil)

	// Act
	company, err := service.GetCompany(id)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, expectedCompany.ID, company.ID)
	assert.Equal(t, expectedCompany.Name, company.Name)
	assert.Equal(t, expectedCompany.Owner, company.Owner)
	assert.Equal(t, expectedCompany.CreatedAt, company.CreatedAt)
	assert.Equal(t, expectedCompany.UpdatedAt, company.UpdatedAt)

	mockRepo.AssertExpectations(t)
}

func TestGetCompany_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()

	mockRepo.On("GetByID", id).Return(nil, errors.New("record not found"))

	// Act
	company, err := service.GetCompany(id)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, company)
	assert.Equal(t, ErrCompanyNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateCompany_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()
	existingCompany := &domain.Company{
		ID:        id,
		Name:      "Old Company",
		Owner:     "Old Owner",
		CreatedAt: time.Now().Add(-2 * time.Hour),
		UpdatedAt: time.Now().Add(-time.Hour),
	}

	newName := "New Company"
	newOwner := "New Owner"
	updatedAt := time.Now()

	mockRepo.On("GetByID", id).Return(existingCompany, nil)
	mockRepo.On("Update", existingCompany).Return(nil).Run(func(args mock.Arguments) {
		company := args.Get(0).(*domain.Company)
		company.Name = newName
		company.Owner = newOwner
		company.UpdatedAt = updatedAt
	})

	// Act
	updatedCompany, err := service.UpdateCompany(id, newName, newOwner)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, updatedCompany)
	assert.Equal(t, newName, updatedCompany.Name)
	assert.Equal(t, newOwner, updatedCompany.Owner)
	assert.Equal(t, existingCompany.ID, updatedCompany.ID)
	assert.Equal(t, existingCompany.CreatedAt, updatedCompany.CreatedAt)
	assert.Equal(t, updatedAt, updatedCompany.UpdatedAt)

	mockRepo.AssertExpectations(t)
}

func TestUpdateCompany_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()
	newName := "New Company"
	newOwner := "New Owner"

	mockRepo.On("GetByID", id).Return(nil, errors.New("record not found"))

	// Act
	updatedCompany, err := service.UpdateCompany(id, newName, newOwner)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, updatedCompany)
	assert.Equal(t, ErrCompanyNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteCompany_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()

	mockRepo.On("Delete", id).Return(nil)

	// Act
	err := service.DeleteCompany(id)

	// Assert
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteCompany_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	id := uuid.New()

	mockRepo.On("Delete", id).Return(errors.New("record not found"))

	// Act
	err := service.DeleteCompany(id)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrCompanyNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestListCompanies_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	companies := []domain.Company{
		{
			ID:        uuid.New(),
			Name:      "Company One",
			Owner:     "Owner One",
			CreatedAt: time.Now().Add(-3 * time.Hour),
			UpdatedAt: time.Now().Add(-2 * time.Hour),
		},
		{
			ID:        uuid.New(),
			Name:      "Company Two",
			Owner:     "Owner Two",
			CreatedAt: time.Now().Add(-4 * time.Hour),
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		},
	}

	mockRepo.On("GetAll").Return(companies, nil)

	// Act
	result, err := service.ListCompanies()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, companies, result)

	mockRepo.AssertExpectations(t)
}

func TestListCompanies_Error(t *testing.T) {
	// Arrange
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	mockRepo.On("GetAll").Return(nil, errors.New("database error"))

	// Act
	result, err := service.ListCompanies()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errors.New("database error"), err)

	mockRepo.AssertExpectations(t)
}

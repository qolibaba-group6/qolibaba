
package service

import (
	"testing"

	"companies-service/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockCompanyRepository struct {
	mock.Mock
}

func (m *MockCompanyRepository) Create(company *models.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func (m *MockCompanyRepository) GetAll() ([]models.Company, error) {
	args := m.Called()
	return args.Get(0).([]models.Company), args.Error(1)
}

func (m *MockCompanyRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCompanyRepository) GetByID(id uuid.UUID) (*models.Company, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Company), args.Error(1)
}

func (m *MockCompanyRepository) Update(company *models.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func TestCreateCompany(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	service := NewCompanyService(mockRepo)

	mockRepo.On("Create", mock.Anything).Return(nil)

	company, err := service.CreateCompany("Test Company", "Test Owner")
	require.NoError(t, err)
	require.Equal(t, "Test Company", company.Name)
	require.Equal(t, "Test Owner", company.Owner)

	mockRepo.AssertExpectations(t)
}

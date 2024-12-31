package handlers

import (
	"bytes"
	"companies-service/internal/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// تعریف یک Mock برای سرویس شرکت
type MockCompanyService struct {
	mock.Mock
}

func (m *MockCompanyService) GetAllCompanies() ([]models.Company, error) {
	args := m.Called()
	return args.Get(0).([]models.Company), args.Error(1)
}

func (m *MockCompanyService) GetCompanyByID(id uuid.UUID) (*models.Company, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Company), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCompanyService) CreateCompany(name, owner string) (*models.Company, error) {
	args := m.Called(name, owner)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Company), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCompanyService) UpdateCompany(id uuid.UUID, name, owner string) (*models.Company, error) {
	args := m.Called(id, name, owner)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Company), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCompanyService) DeleteCompany(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllCompanies(t *testing.T) {
	// Arrange
	mockService := new(MockCompanyService)
	handler := NewCompanyHandler(mockService)

	companies := []models.Company{
		{
			ID:        uuid.New(),
			Name:      "Company A",
			Owner:     "Owner A",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Name:      "Company B",
			Owner:     "Owner B",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockService.On("GetAllCompanies").Return(companies, nil)

	router := gin.Default()
	router.GET("/api/companies", handler.GetAllCompanies)

	// Act
	req, _ := http.NewRequest("GET", "/api/companies", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusOK, resp.Code)
	var responseCompanies []models.Company
	err := json.Unmarshal(resp.Body.Bytes(), &responseCompanies)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(responseCompanies))
	mockService.AssertExpectations(t)
}

func TestGetCompanyByID_NotFound(t *testing.T) {
	// Arrange
	mockService := new(MockCompanyService)
	handler := NewCompanyHandler(mockService)

	testID := uuid.New()
	mockService.On("GetCompanyByID", testID).Return(nil, errors.New("not found"))

	router := gin.Default()
	router.GET("/api/companies/:id", handler.GetCompanyByID)

	// Act
	req, _ := http.NewRequest("GET", "/api/companies/"+testID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, resp.Code)
	var response map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Company not found", response["error"])
	mockService.AssertExpectations(t)
}

func TestCreateCompany_Success(t *testing.T) {
	// Arrange
	mockService := new(MockCompanyService)
	handler := NewCompanyHandler(mockService)

	input := map[string]string{
		"name":  "Company C",
		"owner": "Owner C",
	}
	inputJSON, _ := json.Marshal(input)

	createdCompany := &models.Company{
		ID:        uuid.New(),
		Name:      "Company C",
		Owner:     "Owner C",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockService.On("CreateCompany", "Company C", "Owner C").Return(createdCompany, nil)

	router := gin.Default()
	router.POST("/api/companies", handler.CreateCompany)

	// Act
	req, _ := http.NewRequest("POST", "/api/companies", bytes.NewBuffer(inputJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusCreated, resp.Code)
	var responseCompany models.Company
	err := json.Unmarshal(resp.Body.Bytes(), &responseCompany)
	assert.NoError(t, err)
	assert.Equal(t, "Company C", responseCompany.Name)
	assert.Equal(t, "Owner C", responseCompany.Owner)
	mockService.AssertExpectations(t)
}

func TestCreateCompany_BadRequest(t *testing.T) {
	// Arrange
	mockService := new(MockCompanyService)
	handler := NewCompanyHandler(mockService)

	input := map[string]string{
		"name":  "", // نام خالی
		"owner": "Owner D",
	}
	inputJSON, _ := json.Marshal(input)

	router := gin.Default()
	router.POST("/api/companies", handler.CreateCompany)

	// Act
	req, _ := http.NewRequest("POST", "/api/companies", bytes.NewBuffer(inputJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	var response map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Key: 'CreateCompanyInput.Name' Error:Field validation for 'Name' failed on the 'required' tag")
}

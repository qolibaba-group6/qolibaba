
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"companies-service/internal/config"
	"companies-service/internal/db"
	"companies-service/internal/handlers"
	"companies-service/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	_, err := config.LoadConfig()
	require.NoError(t, err, "Failed to load config")
}

func TestDatabaseConnection(t *testing.T) {
	cfg := "postgres://user:password@localhost:5432/transport?sslmode=disable" // Example connection string
	conn, err := db.Connect(cfg)
	require.NoError(t, err, "Failed to connect to database")
	defer conn.Close()
}

func TestRoutes(t *testing.T) {
	router := gin.Default()

	// Mock handler
	mockHandler := &handlers.CompanyHandler{}
	routes.RegisterCompanyRoutes(router, mockHandler)

	req, _ := http.NewRequest("GET", "/api/companies", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code, "Expected status OK for /api/companies")
}

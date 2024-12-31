package handlers

import (
	"companies-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CompanyHandler struct {
	service service.CompanyService
}

func NewCompanyHandler(service service.CompanyService) *CompanyHandler {
	return &CompanyHandler{service}
}

func RegisterCompanyRoutes(router *gin.Engine, handler *CompanyHandler) {
	companies := router.Group("/api/companies")
	{
		companies.GET("/", handler.GetAllCompanies)
		companies.GET("/:id", handler.GetCompanyByID)
		companies.POST("/", handler.CreateCompany)
		companies.PUT("/:id", handler.UpdateCompany)
		companies.DELETE("/:id", handler.DeleteCompany)
	}
}

func (h *CompanyHandler) GetAllCompanies(c *gin.Context) {
	companies, err := h.service.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (h *CompanyHandler) GetCompanyByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	company, err := h.service.GetCompanyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, company)
}

type CreateCompanyInput struct {
	Name  string `json:"name" binding:"required"`
	Owner string `json:"owner" binding:"required"`
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var input CreateCompanyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, err := h.service.CreateCompany(input.Name, input.Owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, company)
}

type UpdateCompanyInput struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var input UpdateCompanyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, err := h.service.UpdateCompany(id, input.Name, input.Owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	if err := h.service.DeleteCompany(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

package routes

import (
	"companies-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCompanyRoutes(router *gin.Engine, handler *handlers.CompanyHandler) {
	companies := router.Group("/api/companies")
	{
		companies.GET("/", handler.GetAllCompanies)
		companies.GET("/:id", handler.GetCompanyByID)
		companies.POST("/", handler.CreateCompany)
		companies.PUT("/:id", handler.UpdateCompany)
		companies.DELETE("/:id", handler.DeleteCompany)
	}
}

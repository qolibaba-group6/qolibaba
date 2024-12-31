
package routes

import (
	"companies-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterTravelRoutes(router *gin.Engine, handler *handlers.TravelHandler) {
	travels := router.Group("/api/travels")
	{
		travels.GET("/", handler.GetAllTravels)
		travels.GET("/:id", handler.GetTravelByID)
		travels.POST("/", handler.CreateTravel)
		travels.PUT("/:id", handler.UpdateTravel)
		travels.DELETE("/:id", handler.DeleteTravel)
	}
}

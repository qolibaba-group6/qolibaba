package http

import (
	"qolibaba/internal/ports/handlers"
	"qolibaba/internal/ports/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GinRouter struct {
	router *gin.Engine
}

func NewGinRouter() *GinRouter {
	return &GinRouter{router: gin.Default()}
}

func (g *GinRouter) Start(address string, db *gorm.DB) {
	vehicleRepo := repository.NewGormVehicleRepository(db)
	vehicleService := services.NewVehicleService(vehicleRepo)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)

	v1 := g.router.Group("/api/v1")
	v1.POST("/vehicles", vehicleHandler.RegisterVehicle)
	v1.GET("/vehicles", vehicleHandler.GetVehicles)
	v1.PUT("/vehicles/:id/status", vehicleHandler.UpdateVehicleStatus)

	g.router.Run(address)
}

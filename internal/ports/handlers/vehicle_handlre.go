package handlers

import (
	"net/http"
	"qolibaba/internal/core/models"
	"qolibaba/internal/core/services"

	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
	service *services.VehicleService
}

func NewVehicleHandler(service *services.VehicleService) *VehicleHandler {
	return &VehicleHandler{service: service}
}

func (h *VehicleHandler) RegisterVehicle(c *gin.Context) {
	var vehicle models.Vehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.RegisterVehicle(&vehicle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register vehicle"})
		return
	}
	c.JSON(http.StatusCreated, vehicle)
}

func (h *VehicleHandler) GetVehicles(c *gin.Context) {
	vehicles, err := h.service.GetVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch vehicles"})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func (h *VehicleHandler) UpdateVehicleStatus(c *gin.Context) {
	var input struct {
		Status string `json:"status"`
	}
	id := c.Param("id")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateVehicleStatus(id, input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update status"})
		return
	}
	c.Status(http.StatusOK)
}


package handlers

import (
	"companies-service/internal/models"
	"companies-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TravelHandler struct {
	service service.TravelService
}

func NewTravelHandler(service service.TravelService) *TravelHandler {
	return &TravelHandler{service: service}
}

func (h *TravelHandler) GetAllTravels(c *gin.Context) {
	travels, err := h.service.GetAllTravels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, travels)
}

func (h *TravelHandler) GetTravelByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	travel, err := h.service.GetTravelByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "travel not found"})
		return
	}
	c.JSON(http.StatusOK, travel)
}

func (h *TravelHandler) CreateTravel(c *gin.Context) {
	var travel models.Travel
	if err := c.ShouldBindJSON(&travel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateTravel(&travel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, travel)
}

func (h *TravelHandler) UpdateTravel(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var travel models.Travel
	if err := c.ShouldBindJSON(&travel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	travel.ID = id
	if err := h.service.UpdateTravel(&travel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, travel)
}

func (h *TravelHandler) DeleteTravel(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.service.DeleteTravel(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

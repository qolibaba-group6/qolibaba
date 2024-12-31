package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qolibaba/internal/travel_agencies"
	"qolibaba/pkg/adapter/storage/types"
)

type TravelAgencyHandler struct {
	TravelAgencyService *travel_agencies.TravelAgencyService
}

func NewTravelAgencyHandler(service *travel_agencies.TravelAgencyService) *TravelAgencyHandler {
	return &TravelAgencyHandler{TravelAgencyService: service}
}

func (h *TravelAgencyHandler) RegisterNewAgency(c *fiber.Ctx) error {
	var agency types.TravelAgency

	if err := c.BodyParser(&agency); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}
	createdAgency, err := h.TravelAgencyService.RegisterNewAgency(&agency)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(createdAgency)
}

func (h *TravelAgencyHandler) GetAllHotelsAndVehiclesHandler(c *fiber.Ctx) error {
	data, err := h.TravelAgencyService.GetAllHotelsAndVehicles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error retrieving hotels and vehicles: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qolibaba/internal/travel_agencies"
	"qolibaba/pkg/adapter/storage/types"
	"time"
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

func (h *TravelAgencyHandler) OfferTourHandler(c *fiber.Ctx) error {
	var input struct {
		UserID                  uint      `json:"user_id"`
		RoomID                  uint      `json:"room_id"`
		StartTime               time.Time `json:"start_time"`
		EndTime                 time.Time `json:"end_time"`
		TotalPrice              float64   `json:"total_price"`
		GoingTransferVehicleID  uint      `json:"going_transfer_vehicle_id"`
		ReturnTransferVehicleID uint      `json:"return_transfer_vehicle_id"`
		HotelID                 uint      `json:"hotel_id"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid request body: %v", err),
		})
	}

	tour, err := h.TravelAgencyService.OfferTour(input.UserID, input.RoomID, input.StartTime, input.EndTime, input.TotalPrice, input.GoingTransferVehicleID, input.ReturnTransferVehicleID, input.HotelID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error offering tour: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tour offered successfully",
		"tour":    tour,
	})
}

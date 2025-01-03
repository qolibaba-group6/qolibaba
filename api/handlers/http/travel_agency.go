package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qolibaba/internal/travel_agencies/port"
	"qolibaba/pkg/adapter/storage/types"
	"strconv"
)

type TravelAgencyHandler struct {
	TravelAgencyService port.Service
}

func NewTravelAgencyHandler(service port.Service) *TravelAgencyHandler {
	return &TravelAgencyHandler{TravelAgencyService: service}
}

func (h *TravelAgencyHandler) RegisterNewAgency(c *fiber.Ctx) error {
	var agency types.TravelAgency

	if err := c.BodyParser(&agency); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}

	createdAgency, err := h.TravelAgencyService.RegisterNewAgency(&agency)
	if err != nil {
		if err.Error() == "email already in use" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdAgency)
}

func (h *TravelAgencyHandler) GetAllHotelsAndVehiclesHandler(c *fiber.Ctx) error {
	data, err := h.TravelAgencyService.GetAllHotelsAndVehicles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *TravelAgencyHandler) CreateTour(c *fiber.Ctx) error {
	var tourData types.Tour

	if err := c.BodyParser(&tourData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid request body: %v", err),
		})
	}

	tour, err := h.TravelAgencyService.OfferTour(&tourData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create tour: %v", err),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Tour created successfully",
		"tour":    tour,
	})
}

func (h *TravelAgencyHandler) CreateBooking(c *fiber.Ctx) error {
	var bookingData types.TourBooking

	if err := c.BodyParser(&bookingData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid request body: %v", err),
		})
	}

	booking, err := h.TravelAgencyService.CreateTourBooking(&bookingData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create booking: %v", err),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Booking created successfully",
		"booking": booking,
	})
}

func (h *TravelAgencyHandler) ConfirmTourBooking(c *fiber.Ctx) error {
	bookingID, err := strconv.ParseUint(c.Params("bookingID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid booking ID"})
	}

	booking, err := h.TravelAgencyService.ConfirmTourBooking(uint(bookingID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Error confirming tour booking: %v", err)})
	}

	return c.Status(fiber.StatusOK).JSON(booking)
}

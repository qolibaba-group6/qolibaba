package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qolibaba/internal/hotels/domain/entity"
	"qolibaba/internal/hotels/port"
)

type HotelHandler struct {
	hotelService port.Service
}

func NewHotelHandler(hotelService port.Service) *HotelHandler {
	return &HotelHandler{
		hotelService: hotelService,
	}
}

func (h *HotelHandler) RegisterHotelHandler(c *fiber.Ctx) error {
	var hotel entity.Hotel

	if err := c.BodyParser(&hotel); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if hotel.Name == "" || hotel.Location == "" || hotel.PhoneNumber == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, Location, and PhoneNumber are required",
		})
	}

	createdOrUpdatedHotel, err := h.hotelService.RegisterHotel(&hotel)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to register hotel: %v", err),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Hotel registered successfully",
		"hotel":   createdOrUpdatedHotel,
	})
}

// GetHotelByIDHandler retrieves a hotel by its ID
func (h *HotelHandler) GetHotelByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	hotel, err := h.hotelService.GetHotelByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to fetch hotel: %v", err),
		})
	}

	if hotel == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Hotel not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"hotel": hotel,
	})
}

// GetAllHotelsHandler retrieves all hotels
func (h *HotelHandler) GetAllHotelsHandler(c *fiber.Ctx) error {
	hotels, err := h.hotelService.GetAllHotels()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to fetch hotels: %v", err),
		})
	}

	if len(hotels) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "No hotels found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"hotels": hotels,
	})
}

func (h *HotelHandler) DeleteHotelHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Hotel ID is required",
		})
	}

	if err := h.hotelService.DeleteHotel(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to delete hotel: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Hotel deleted successfully",
	})
}

func (h *HotelHandler) CreateOrUpdateRoom(c *fiber.Ctx) error {
	var room entity.Room

	if err := c.BodyParser(&room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid request body: %v", err),
		})
	}

	createdOrUpdatedRoom, err := h.hotelService.CreateOrUpdateRoom(&room)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create or update room: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(createdOrUpdatedRoom)
}

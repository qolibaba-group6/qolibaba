package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"qolibaba/internal/hotels/port"
	"qolibaba/pkg/adapter/storage/types"
	"strconv"
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
	var hotel types.Hotel

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
	var room types.Room

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

// GetRoomByID handles the request to get a room by its ID.
func (h *HotelHandler) GetRoomByID(c *fiber.Ctx) error {
	id := c.Params("id")
	roomID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("invalid room ID: %v", err),
		})
	}

	room, err := h.hotelService.GetRoomByID(uint(roomID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("room not found: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(room)
}

// GetRoomsByHotelID handles the request to get all rooms by hotel ID.
func (h *HotelHandler) GetRoomsByHotelID(c *fiber.Ctx) error {
	hotelID := c.Params("hotel_id")

	hotelIDUint, err := strconv.Atoi(hotelID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("invalid hotel ID: %v", err),
		})
	}

	rooms, err := h.hotelService.GetRoomsByHotelID(uint(hotelIDUint))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("rooms not found for hotel %d: %v", hotelIDUint, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(rooms)
}

// DeleteRoom handles the request to delete a room by its ID.
func (h *HotelHandler) DeleteRoom(c *fiber.Ctx) error {
	id := c.Params("id")

	roomID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("invalid room ID: %v", err),
		})
	}

	err = h.hotelService.DeleteRoom(uint(roomID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to delete room: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "room deleted successfully",
	})
}

// CreateBooking handles booking creation
func (h *HotelHandler) CreateBooking(c *fiber.Ctx) error {
	var booking types.Booking
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid request body: %v", err),
		})
	}

	createdBooking, err := h.hotelService.CreateBooking(&types.Booking{
		RoomID:    booking.RoomID,
		UserID:    booking.UserID,
		StartTime: booking.StartTime,
		EndTime:   booking.EndTime,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error creating booking: %v", err),
		})
	}

	return c.Status(http.StatusCreated).JSON(createdBooking)
}

// GetBookingByID handles fetching a booking by ID
func (h *HotelHandler) GetBookingByID(c *fiber.Ctx) error {
	id := c.Params("id")
	bookingID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid booking ID format",
		})
	}

	booking, err := h.hotelService.GetBookingByID(uint(bookingID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Booking not found: %v", err),
		})
	}
	return c.JSON(booking)
}

// GetBookingsByUserID handles fetching all bookings for a given user
func (h *HotelHandler) GetBookingsByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	bookings, err := h.hotelService.GetBookingsByUserID(uint(userIDUint))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error fetching bookings for user: %v", err),
		})
	}

	return c.JSON(bookings)
}

// SoftDeleteBooking handles soft deleting a booking
func (h *HotelHandler) SoftDeleteBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	bookingID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid booking ID format",
		})
	}
	err = h.hotelService.SoftDeleteBooking(uint(bookingID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error deleting booking: %v", err),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

// ConfirmBooking handles confirming a completed booking
func (h *HotelHandler) ConfirmBooking(c *fiber.Ctx) error {
	// Get booking ID from the URL
	id := c.Params("id")

	// Convert id to uint
	bookingID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid booking ID format",
		})
	}

	// Call the booking service to confirm the booking
	confirmedBooking, err := h.hotelService.ConfirmBooking(uint(bookingID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error confirming booking: %v", err),
		})
	}

	return c.JSON(confirmedBooking)
}

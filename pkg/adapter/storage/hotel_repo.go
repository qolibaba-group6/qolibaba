package storage

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"qolibaba/internal/hotels/port"
	"qolibaba/pkg/adapter/storage/types"
	"time"
)

type hotelRepo struct {
	db *gorm.DB
}

func NewHotelRepo(db *gorm.DB) port.Repo {
	return &hotelRepo{
		db: db,
	}
}

func (r *hotelRepo) RegisterHotel(hotel *types.Hotel) (*types.Hotel, error) {
	var existingHotel types.Hotel
	if err := r.db.Where("name = ?", hotel.Name).First(&existingHotel).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("error checking hotel existence: %v", err)
		}
		if err := r.db.Create(hotel).Error; err != nil {
			return nil, fmt.Errorf("error creating hotel: %v", err)
		}
		return hotel, nil
	}

	existingHotel.Location = hotel.Location
	existingHotel.PhoneNumber = hotel.PhoneNumber
	existingHotel.Email = hotel.Email
	existingHotel.Website = hotel.Website
	existingHotel.UpdatedAt = time.Now()
	if err := r.db.Save(&existingHotel).Error; err != nil {
		return nil, fmt.Errorf("error updating hotel: %v", err)
	}
	return &existingHotel, nil
}

// GetHotelByID fetches a hotel by its ID.
func (r *hotelRepo) GetHotelByID(id string) (*types.Hotel, error) {
	var hotel types.Hotel
	if err := r.db.Where("id = ?", id).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Hotel not found
		}
		return nil, fmt.Errorf("error fetching hotel by ID: %v", err)
	}
	return &hotel, nil
}

// GetAllHotels fetches all hotels. If no hotels are found, it returns a custom error.
func (r *hotelRepo) GetAllHotels() ([]types.Hotel, error) {
	var hotels []types.Hotel
	if err := r.db.Find(&hotels).Error; err != nil {
		return nil, fmt.Errorf("error fetching all hotels: %v", err)
	}

	if len(hotels) == 0 {
		return nil, fmt.Errorf("no hotels found in the database")
	}

	return hotels, nil
}

func (r *hotelRepo) DeleteHotel(id string) error {
	var hotel types.Hotel
	if err := r.db.Where("id = ?", id).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("hotel with ID %s not found", id)
		}
		return fmt.Errorf("error fetching hotel by ID: %v", err)
	}

	if err := r.db.Delete(&hotel).Error; err != nil {
		return fmt.Errorf("error deleting hotel: %v", err)
	}
	return nil
}

// CreateOrUpdateRoom creates a new room or updates an existing one.
func (r *hotelRepo) CreateOrUpdateRoom(room *types.Room) (*types.Room, error) {
	var hotel types.Hotel
	if err := r.db.Where("id = ?", room.HotelID).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("hotel with ID %d does not exist", room.HotelID)
		}
		return nil, fmt.Errorf("error checking hotel existence: %v", err)
	}

	var existingRoom types.Room
	if err := r.db.Where("id = ?", room.ID).First(&existingRoom).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if room.Status == "" {
				room.Status = types.RoomStatusFree
			}

			if err := r.db.Create(room).Error; err != nil {
				return nil, fmt.Errorf("error creating room: %v", err)
			}
			return room, nil
		}
		return nil, fmt.Errorf("error checking room existence: %v", err)
	}

	existingRoom.Type = room.Type
	existingRoom.Price = room.Price
	existingRoom.Capacity = room.Capacity
	existingRoom.Features = room.Features
	existingRoom.Duration = room.Duration
	existingRoom.PublicReleaseDate = room.PublicReleaseDate
	existingRoom.AgencyReleaseDate = room.AgencyReleaseDate
	existingRoom.Status = room.Status
	existingRoom.UpdatedAt = time.Now()

	if err := r.db.Save(&existingRoom).Error; err != nil {
		return nil, fmt.Errorf("error updating room: %v", err)
	}
	return &existingRoom, nil
}

func (r *hotelRepo) GetRoomByID(id uint) (*types.Room, error) {
	var room types.Room
	if err := r.db.First(&room, id).Error; err != nil {
		return nil, fmt.Errorf("error fetching room by ID: %v", err)
	}
	return &room, nil
}

func (r *hotelRepo) GetRoomsByHotelID(hotelID uint) ([]types.Room, error) {
	var rooms []types.Room
	if err := r.db.Where("hotel_id = ?", hotelID).Find(&rooms).Error; err != nil {
		return nil, fmt.Errorf("error fetching rooms by hotel ID: %v", err)
	}
	return rooms, nil
}

func (r *hotelRepo) DeleteRoom(id uint) error {
	if err := r.db.Delete(&types.Room{}, id).Error; err != nil {
		return fmt.Errorf("error deleting room: %v", err)
	}
	return nil
}

// CreateBooking creates a new booking in the system.
func (r *hotelRepo) CreateBooking(booking *types.Booking) (*types.Booking, error) {
	if booking.Confirmed == false {
		booking.Confirmed = false
	}
	if booking.StartTime.After(booking.EndTime) {
		return nil, fmt.Errorf("start time must be before end time")
	}
	var existingBooking types.Booking
	if err := r.db.Where("room_id = ? AND user_id = ? AND start_time = ? AND end_time = ?",
		booking.RoomID, booking.UserID, booking.StartTime, booking.EndTime).First(&existingBooking).Error; err == nil {
		existingBooking.TotalPrice = booking.TotalPrice
		existingBooking.Status = booking.Status
		existingBooking.Confirmed = booking.Confirmed
		existingBooking.DateOfConfirmation = booking.DateOfConfirmation
		existingBooking.IsReferred = booking.IsReferred

		if err := r.db.Save(&existingBooking).Error; err != nil {
			return nil, fmt.Errorf("error updating booking: %v", err)
		}

		return &existingBooking, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := r.db.Create(booking).Error; err != nil {
			return nil, fmt.Errorf("error creating booking: %v", err)
		}

		return booking, nil
	} else {
		return nil, fmt.Errorf("error checking booking existence: %v", err)
	}
}

// UpdateBooking updates an existing booking.
func (r *hotelRepo) UpdateBooking(booking *types.Booking) (*types.Booking, error) {
	var existingBooking types.Booking
	if err := r.db.First(&existingBooking, booking.ID).Error; err != nil {
		return nil, fmt.Errorf("error finding booking: %v", err)
	}

	if err := r.db.Save(booking).Error; err != nil {
		return nil, fmt.Errorf("error updating booking: %v", err)
	}

	return booking, nil
}

// SoftDeleteBooking soft deletes a booking by setting the DeletedAt field.
func (r *hotelRepo) SoftDeleteBooking(id uint) error {
	var booking types.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		return fmt.Errorf("booking not found: %v", err)
	}

	now := time.Now()
	booking.DeletedAt = &now
	if err := r.db.Save(&booking).Error; err != nil {
		return fmt.Errorf("error deleting booking: %v", err)
	}

	return nil
}

// GetBookingByID retrieves a booking by its ID.
func (r *hotelRepo) GetBookingByID(id uint) (*types.Booking, error) {
	var booking types.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		return nil, fmt.Errorf("error finding booking: %v", err)
	}
	return &booking, nil
}

// GetBookingsByUserID retrieves all bookings for a given user ID.
func (r *hotelRepo) GetBookingsByUserID(userID uint) ([]types.Booking, error) {
	var bookings []types.Booking
	if err := r.db.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, fmt.Errorf("error fetching bookings for user %d: %v", userID, err)
	}
	return bookings, nil
}

func (r *hotelRepo) ConfirmBooking(bookingID uint) (*types.Booking, error) {
	var booking types.Booking
	if err := r.db.First(&booking, bookingID).Error; err != nil {
		return nil, fmt.Errorf("error fetching booking with ID %d: %v", bookingID, err)
	}
	if booking.Status != types.BookingStatusCompleted {
		return nil, fmt.Errorf("booking is not completed, cannot confirm")
	}
	booking.Confirmed = true
	confirmedAt := time.Now()
	booking.DateOfConfirmation = &confirmedAt

	if err := r.db.Save(&booking).Error; err != nil {
		return nil, fmt.Errorf("error confirming booking with ID %d: %v", bookingID, err)
	}

	return &booking, nil
}

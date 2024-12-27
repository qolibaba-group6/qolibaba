package hotels

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"qolibaba/internal/hotels/port"
	"qolibaba/pkg/adapter/storage/types"
	"regexp"
)

type service struct {
	hotelRepo port.Repo
	validate  *validator.Validate
}

func NewService(repo port.Repo) port.Service {
	return &service{
		hotelRepo: repo,
		validate:  validator.New(),
	}
}

func (s *service) RegisterHotel(hotel *types.Hotel) (*types.Hotel, error) {
	if err := s.validate.Struct(hotel); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	if err := validatePhoneNumber(hotel.PhoneNumber); err != nil {
		return nil, err
	}

	if err := validateEmail(hotel.Email); err != nil {
		return nil, err
	}

	createdOrUpdatedHotel, err := s.hotelRepo.RegisterHotel(hotel)
	if err != nil {
		return nil, fmt.Errorf("failed to register or update hotel: %v", err)
	}

	return createdOrUpdatedHotel, nil
}

func (s *service) GetHotelByID(hotelID string) (*types.Hotel, error) {
	hotel, err := s.hotelRepo.GetHotelByID(hotelID)
	if err != nil {
		return nil, fmt.Errorf("error fetching hotel by ID: %v", err)
	}

	if hotel == nil {
		return nil, fmt.Errorf("hotel with ID %s not found", hotelID)
	}

	return hotel, nil
}

func (s *service) GetAllHotels() ([]types.Hotel, error) {
	hotels, err := s.hotelRepo.GetAllHotels()
	if err != nil {
		return nil, fmt.Errorf("error fetching all hotels: %v", err)
	}

	return hotels, nil
}

func (s *service) DeleteHotel(id string) error {
	if err := s.hotelRepo.DeleteHotel(id); err != nil {
		return fmt.Errorf("failed to delete hotel: %v", err)
	}
	return nil
}

// CreateOrUpdateRoom creates a new room or updates an existing one.
func (s *service) CreateOrUpdateRoom(room *types.Room) (*types.Room, error) {
	if err := s.validate.Struct(room); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	if room.Price <= 0 {
		return nil, fmt.Errorf("price must be greater than zero")
	}
	if room.Capacity <= 0 {
		return nil, fmt.Errorf("capacity must be greater than zero")
	}

	createdOrUpdatedRoom, err := s.hotelRepo.CreateOrUpdateRoom(room)
	if err != nil {
		return nil, fmt.Errorf("failed to create or update room: %v", err)
	}

	return createdOrUpdatedRoom, nil
}

// GetRoomByID fetches a room by its ID.
func (s *service) GetRoomByID(id uint) (*types.Room, error) {
	room, err := s.hotelRepo.GetRoomByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get room: %v", err)
	}
	return room, nil
}

// GetRoomsByHotelID fetches all rooms for a given hotel by its ID.
func (s *service) GetRoomsByHotelID(hotelID uint) ([]types.Room, error) {
	rooms, err := s.hotelRepo.GetRoomsByHotelID(hotelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms for hotel %d: %v", hotelID, err)
	}
	return rooms, nil
}

// DeleteRoom deletes a room by its ID.
func (s *service) DeleteRoom(id uint) error {
	err := s.hotelRepo.DeleteRoom(id)
	if err != nil {
		return fmt.Errorf("failed to delete room: %v", err)
	}
	return nil
}

func validatePhoneNumber(phone string) error {
	re := regexp.MustCompile(`^\+98\d{10}$`)
	if !re.MatchString(phone) {
		return fmt.Errorf("invalid phone number format, must be E164 (e.g., +1234567890)")
	}
	return nil
}

func validateEmail(email *string) error {
	if email != nil && *email != "" {
		re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !re.MatchString(*email) {
			return fmt.Errorf("invalid email format")
		}
	}
	return nil
}

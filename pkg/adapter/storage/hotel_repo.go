package storage

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"qolibaba/internal/hotels/domain/entity"
	"qolibaba/internal/hotels/port"
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

func (r *hotelRepo) RegisterHotel(hotel *entity.Hotel) (*entity.Hotel, error) {
	var existingHotel entity.Hotel
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
func (r *hotelRepo) GetHotelByID(id string) (*entity.Hotel, error) {
	var hotel entity.Hotel
	if err := r.db.Where("id = ?", id).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Hotel not found
		}
		return nil, fmt.Errorf("error fetching hotel by ID: %v", err)
	}
	return &hotel, nil
}

// GetAllHotels fetches all hotels. If no hotels are found, it returns a custom error.
func (r *hotelRepo) GetAllHotels() ([]entity.Hotel, error) {
	var hotels []entity.Hotel
	if err := r.db.Find(&hotels).Error; err != nil {
		return nil, fmt.Errorf("error fetching all hotels: %v", err)
	}

	if len(hotels) == 0 {
		return nil, fmt.Errorf("no hotels found in the database")
	}

	return hotels, nil
}

func (r *hotelRepo) DeleteHotel(id string) error {
	var hotel entity.Hotel
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
func (r *hotelRepo) CreateOrUpdateRoom(room *entity.Room) (*entity.Room, error) {
	var existingRoom entity.Room
	if err := r.db.Where("id = ?", room.ID).First(&existingRoom).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
	existingRoom.UpdatedAt = time.Now()

	if err := r.db.Save(&existingRoom).Error; err != nil {
		return nil, fmt.Errorf("error updating room: %v", err)
	}
	return &existingRoom, nil
}

func (r *hotelRepo) GetRoomByID(id uint) (*entity.Room, error) {
	var room entity.Room
	if err := r.db.First(&room, id).Error; err != nil {
		return nil, fmt.Errorf("error fetching room by ID: %v", err)
	}
	return &room, nil
}

func (r *hotelRepo) GetRoomsByHotelID(hotelID uint) ([]entity.Room, error) {
	var rooms []entity.Room
	if err := r.db.Where("hotel_id = ?", hotelID).Find(&rooms).Error; err != nil {
		return nil, fmt.Errorf("error fetching rooms by hotel ID: %v", err)
	}
	return rooms, nil
}

func (r *hotelRepo) DeleteRoom(id uint) error {
	if err := r.db.Delete(&entity.Room{}, id).Error; err != nil {
		return fmt.Errorf("error deleting room: %v", err)
	}
	return nil
}

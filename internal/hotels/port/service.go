package port

import (
	_ "qolibaba/internal/hotels"
	"qolibaba/pkg/adapter/storage/types"
)

type Service interface {
	//hotel services.
	RegisterHotel(hotel *types.Hotel) (*types.Hotel, error)
	GetHotelByID(id string) (*types.Hotel, error)
	GetAllHotels() ([]types.Hotel, error)
	DeleteHotel(id string) error

	// room services.
	CreateOrUpdateRoom(room *types.Room) (*types.Room, error)
	GetRoomByID(id uint) (*types.Room, error)
	GetRoomsByHotelID(hotelID uint) ([]types.Room, error)
	DeleteRoom(id uint) error
}

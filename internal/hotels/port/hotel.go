package port

import (
	"qolibaba/pkg/adapter/storage/types"
)

type Repo interface {
	//hotels repos interfaces
	RegisterHotel(hotel *types.Hotel) (*types.Hotel, error)
	GetHotelByID(hotelID string) (*types.Hotel, error)
	GetAllHotels() ([]types.Hotel, error)
	DeleteHotel(id string) error

	// rooms repos interfaces
	CreateOrUpdateRoom(room *types.Room) (*types.Room, error)
	GetRoomByID(id uint) (*types.Room, error)
	GetRoomsByHotelID(hotelID uint) ([]types.Room, error)
	DeleteRoom(id uint) error
}

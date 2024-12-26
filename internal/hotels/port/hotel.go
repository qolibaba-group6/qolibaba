package port

import "qolibaba/internal/hotels/domain/entity"

type Repo interface {
	//hotels repos interfaces
	RegisterHotel(hotel *entity.Hotel) (*entity.Hotel, error)
	GetHotelByID(hotelID string) (*entity.Hotel, error)
	GetAllHotels() ([]entity.Hotel, error)
	DeleteHotel(id string) error

	// rooms repos interfaces
	CreateOrUpdateRoom(room *entity.Room) (*entity.Room, error)
	GetRoomByID(id uint) (*entity.Room, error)
	GetRoomsByHotelID(hotelID uint) ([]entity.Room, error)
	DeleteRoom(id uint) error
}

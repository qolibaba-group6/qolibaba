package port

import (
	_ "qolibaba/internal/hotels"
	"qolibaba/internal/hotels/domain/entity"
)

type Service interface {
	//hotel services.
	RegisterHotel(hotel *entity.Hotel) (*entity.Hotel, error)
	GetHotelByID(id string) (*entity.Hotel, error)
	GetAllHotels() ([]entity.Hotel, error)
	DeleteHotel(id string) error

	// room services.
	CreateOrUpdateRoom(room *entity.Room) (*entity.Room, error)
}

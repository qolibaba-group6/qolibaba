package hotel

import (
	"qolibaba/config"
	"qolibaba/internal/hotels/port"

	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	HotelService() port.Service
}

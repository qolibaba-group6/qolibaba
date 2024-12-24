package repository

import (
	"gorm.io/gorm"
	"qolibaba/internal/core/models"
)

type VehicleRepository interface {
	CreateVehicle(vehicle *models.Vehicle) error
	FetchVehicles() ([]models.Vehicle, error)
	GetVehicleByID(id uint) (*models.Vehicle, error)
	UpdateVehicle(vehicle *models.Vehicle) error
}

type GormVehicleRepository struct {
	db *gorm.DB
}

func NewGormVehicleRepository(db *gorm.DB) *GormVehicleRepository {
	return &GormVehicleRepository{db: db}
}

func (r *GormVehicleRepository) CreateVehicle(vehicle *models.Vehicle) error {
	return r.db.Create(vehicle).Error
}

func (r *GormVehicleRepository) FetchVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	err := r.db.Find(&vehicles).Error
	return vehicles, err
}

func (r *GormVehicleRepository) GetVehicleByID(id uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	err := r.db.First(&vehicle, id).Error
	return &vehicle, err
}

func (r *GormVehicleRepository) UpdateVehicle(vehicle *models.Vehicle) error {
	return r.db.Save(vehicle).Error
}

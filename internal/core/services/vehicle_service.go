package services

import (
	"errors"
	"qolibaba/internal/core/models"
	"qolibaba/internal/ports/repository"
)

type VehicleService struct {
	repo repository.VehicleRepository
}

func NewVehicleService(repo repository.VehicleRepository) *VehicleService {
	return &VehicleService{repo: repo}
}

func (s *VehicleService) RegisterVehicle(vehicle *models.Vehicle) error {
	return s.repo.CreateVehicle(vehicle)
}

func (s *VehicleService) GetVehicles() ([]models.Vehicle, error) {
	return s.repo.FetchVehicles()
}

func (s *VehicleService) UpdateVehicleStatus(id uint, status string) error {
	vehicle, err := s.repo.GetVehicleByID(id)
	if err != nil {
		return err
	}

	vehicle.Status = status
	return s.repo.UpdateVehicle(vehicle)
}

func (s *VehicleService) MatchVehicle(passengers int) (*models.Vehicle, error) {
	vehicles, err := s.repo.FetchVehicles()
	if err != nil {
		return nil, err
	}

	for _, v := range vehicles {
		if v.Status == "active" && v.Capacity >= passengers {
			return &v, nil
		}
	}
	return nil, errors.New("no suitable vehicle found")
}

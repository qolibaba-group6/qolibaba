package service

import (
	"companies-service/internal/models"
	"companies-service/internal/repository"

	"github.com/google/uuid"
)

type TravelService interface {
	GetAllTravels() ([]models.Travel, error)
	GetTravelByID(id uuid.UUID) (*models.Travel, error)
	CreateTravel(travel *models.Travel) error
	UpdateTravel(travel *models.Travel) error
	DeleteTravel(id uuid.UUID) error
}

type travelService struct {
	repo repository.TravelRepository
}

func NewTravelService(repo repository.TravelRepository) TravelService {
	return &travelService{repo: repo}
}

func (s *travelService) GetAllTravels() ([]models.Travel, error) {
	return s.repo.GetAll()
}

func (s *travelService) GetTravelByID(id uuid.UUID) (*models.Travel, error) {
	return s.repo.GetByID(id)
}

func (s *travelService) CreateTravel(travel *models.Travel) error {
	travel.ID = uuid.New()
	if err := s.repo.Create(travel); err != nil {
		return err
	}
	return nil
}

func (s *travelService) UpdateTravel(travel *models.Travel) error {
	return s.repo.Update(travel)
}

func (s *travelService) DeleteTravel(id uuid.UUID) error {
	return s.repo.Delete(id)
}

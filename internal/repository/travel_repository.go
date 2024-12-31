
package repository

import (
	"companies-service/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TravelRepository interface {
	GetAll() ([]models.Travel, error)
	GetByID(id uuid.UUID) (*models.Travel, error)
	Create(travel *models.Travel) error
	Update(travel *models.Travel) error
	Delete(id uuid.UUID) error
}

type travelRepository struct {
	db *sqlx.DB
}

func NewTravelRepository(db *sqlx.DB) TravelRepository {
	return &travelRepository{db: db}
}

func (r *travelRepository) GetAll() ([]models.Travel, error) {
	var travels []models.Travel
	err := r.db.Select(&travels, "SELECT * FROM travels")
	return travels, err
}

func (r *travelRepository) GetByID(id uuid.UUID) (*models.Travel, error) {
	var travel models.Travel
	err := r.db.Get(&travel, "SELECT * FROM travels WHERE id = $1", id)
	return &travel, err
}

func (r *travelRepository) Create(travel *models.Travel) error {
	_, err := r.db.Exec(
		"INSERT INTO travels (id, type, origin, destination, price, release_date) VALUES ($1, $2, $3, $4, $5, $6)",
		travel.ID, travel.Type, travel.Origin, travel.Destination, travel.Price, travel.ReleaseDate,
	)
	return err
}

func (r *travelRepository) Update(travel *models.Travel) error {
	_, err := r.db.Exec(
		"UPDATE travels SET type = $1, origin = $2, destination = $3, price = $4, release_date = $5 WHERE id = $6",
		travel.Type, travel.Origin, travel.Destination, travel.Price, travel.ReleaseDate, travel.ID,
	)
	return err
}

func (r *travelRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM travels WHERE id = $1", id)
	return err
}

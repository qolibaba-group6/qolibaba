package repository

import (
	"companies-service/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CompanyRepository interface {
	GetAll() ([]models.Company, error)
	GetByID(id uuid.UUID) (*models.Company, error)
	Create(company *models.Company) error
	Update(company *models.Company) error
	Delete(id uuid.UUID) error
}

type companyRepository struct {
	db *sqlx.DB
}

func NewCompanyRepository(db *sqlx.DB) CompanyRepository {
	return &companyRepository{db}
}

func (r *companyRepository) GetAll() ([]models.Company, error) {
	var companies []models.Company
	err := r.db.Select(&companies, "SELECT * FROM companies ORDER BY created_at DESC")
	return companies, err
}

func (r *companyRepository) GetByID(id uuid.UUID) (*models.Company, error) {
	var company models.Company
	err := r.db.Get(&company, "SELECT * FROM companies WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) Create(company *models.Company) error {
	query := `INSERT INTO companies (id, name, owner, created_at, updated_at)
	          VALUES (:id, :name, :owner, NOW(), NOW())`
	_, err := r.db.NamedExec(query, company)
	return err
}

func (r *companyRepository) Update(company *models.Company) error {
	query := `UPDATE companies SET name = :name, owner = :owner, updated_at = NOW() WHERE id = :id`
	_, err := r.db.NamedExec(query, company)
	return err
}

func (r *companyRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM companies WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

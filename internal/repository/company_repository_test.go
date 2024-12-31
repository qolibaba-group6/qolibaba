
package repository

import (
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
    "github.com/jmoiron/sqlx"
)

func TestGetAllCompanies(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("failed to open mock sql database: %s", err)
    }
    defer db.Close()

    mock.ExpectQuery("SELECT (.+) FROM companies").WillReturnRows(
        sqlmock.NewRows([]string{"id", "name", "owner"}).AddRow("c5111dec-03fc-45e4-95ab-5f8745905443", "Company A", "Owner A"),
    )

    repo := NewCompanyRepository(sqlx.NewDb(db, "sqlmock"))
    companies, err := repo.GetAll()
    if err != nil {
        t.Fatalf("unexpected error: %s", err)
    }

    if len(companies) != 1 || companies[0].Name != "Company A" {
        t.Errorf("unexpected companies: %+v", companies)
    }
}

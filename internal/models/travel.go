
package models

import (
	"time"
	"github.com/google/uuid"
)

type Travel struct {
	ID          uuid.UUID `db:"id"`
	Type        string    `db:"type"` // نوع سفر (هوایی، دریایی و ...)
	Origin      string    `db:"origin"`
	Destination string    `db:"destination"`
	Price       float64   `db:"price"`
	ReleaseDate time.Time `db:"release_date"`
}

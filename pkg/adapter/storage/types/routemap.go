package types

import "github.com/google/uuid"

type Terminal struct {
	Model
	Name    string  `gorm:"uniqueIndex:idx_terminal"`
	Type    uint8   `gorm:"uniqueIndex:idx_terminal"`
	Country string  `gorm:"uniqueIndex:idx_terminal"`
	State   string  `gorm:"uniqueIndex:idx_terminal"`
	City    string  `gorm:"uniqueIndex:idx_terminal"`
	Routs   []Route `gorm:"foreignKey:SourceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Route struct {
	Model
	SourceID      uuid.UUID `gorm:"not null"`
	Source        Terminal  `gorm:"foreignKey:SourceID"`
	DestinationID uuid.UUID `gorm:"not null"`
	Destination   Terminal  `gorm:"foreignKey:DestinationID"`
	RouteNumber   uint
	TransportType uint8
	Distance      float64
}

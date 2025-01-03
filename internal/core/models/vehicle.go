package models

type Vehicle struct {
	ID         uint
	UniqueCode string
	Type       string
	Speed      float64
	Capacity   int
	HourlyRate float64
	Status     string
}

type Assignment struct {
	ID        uint
	VehicleID uint
	CompanyID uint
	StartTime string
	EndTime   string
}

package port

import "qolibaba/pkg/adapter/storage/types"

type Repo interface {
	Create(agency *types.TravelAgency) error
	FindByEmail(email string) (*types.TravelAgency, error)
	SaveTour(tour *types.Tour) (*types.Tour, error)
	GetTourByID(tourID uint) (*types.Tour, error)
	CreateTourBooking(tourBooking *types.TourBooking) (*types.TourBooking, error)
	UpdateTourBooking(tourBooking *types.TourBooking) (*types.TourBooking, error)
	ConfirmBooking(bookingID uint) (*types.TourBooking, error)
}

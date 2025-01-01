package port

import "qolibaba/pkg/adapter/storage/types"

type Service interface {
	RegisterNewAgency(agency *types.TravelAgency) (*types.TravelAgency, error)
	GetAllHotelsAndVehicles() (map[string]interface{}, error)
	OfferTour(tour *types.Tour) (*types.Tour, error)
	CreateTourBooking(booking *types.TourBooking) (*types.TourBooking, error)
	ConfirmTourBooking(bookingID uint) (*types.TourBooking, error)
}

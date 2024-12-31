package port

import "qolibaba/pkg/adapter/storage/types"

type Repo interface {
	Create(agency *types.TravelAgency) error
	FindByEmail(email string) (*types.TravelAgency, error)
	SaveTour(tour *types.TourBooking) (*types.TourBooking, error)
}

package port

import "qolibaba/pkg/adapter/storage/types"

type Service interface {
	RegisterNewAgency(agency *types.TravelAgency) (*types.TravelAgency, error)
}

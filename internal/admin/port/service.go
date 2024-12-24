package port

import "context"

type Service interface {
	SayHello(ctx context.Context, name string) (string, error)
	// UpdateUserPermissions()
	// GetUserHistory()
	// GetCompanyHistory()
	// GetTransportsHistory()
	// BlockUserAndAssetsOperator()
	// BlockTransportCompanyOperator()
	// BlockHotelOperator()
	// BlockTravelAgencyOperator()
	// BlockVehicleOperator()
	// AssignTransportCompanyOperator()
	// AssignHotelOperator()
	// AssignTravelAgencyOperator()
	// DismissTransportCompanyOperator()
	// DismissHotelOperator()
	// DismissTravelAgencyOperator()
}

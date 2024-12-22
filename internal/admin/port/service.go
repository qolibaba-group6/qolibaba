package port

type Service interface {
	UpdateUserPermissions()
	GetUserHistory()
	GetCompanyHistory()
	GetTransportsHistory()
	BlockUserAndAssetsOperator()
	BlockTransportCompanyOperator()
	BlockHotelOperator()
	BlockTravelAgencyOperator()
	BlockVehicleOperator()
	AssignTransportCompanyOperator()
	AssignHotelOperator()
	AssignTravelAgencyOperator()
	DismissTransportCompanyOperator()
	DismissHotelOperator()
	DismissTravelAgencyOperator()
}

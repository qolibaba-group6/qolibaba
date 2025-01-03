package http

import (
	"context"
	"fmt"
	"qolibaba/api/service"
	"qolibaba/app"
	"qolibaba/app/bank"
	"qolibaba/app/hotel"
	"qolibaba/config"
	userDomain "qolibaba/internal/user/domain"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.Config) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerAuthAPI(appContainer, cfg.Server, api)
	registerAdminAPI(appContainer, api, cfg)
	registerRoutemapAPI(api, cfg)

	return router.Listen(fmt.Sprintf(":%d", cfg.Server.HttpPort))
}

func RunHotel(appContainer hotel.App, serverCfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerHotelAPI(appContainer, api)

	return router.Listen(fmt.Sprintf(":%d", serverCfg.HttpPort))
}

func RunBank(appContainer bank.App, serverCfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerBankAPI(appContainer, api)

	return router.Listen(fmt.Sprintf(":%d", serverCfg.HttpPort))
}

func registerAuthAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	userPortService := appContainer.UserService(context.Background())
	userService := service.NewUserService(userPortService, cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)
	router.Post("/sign-up", SignUp(userService))
	router.Post("/sign-in", SingIn(userService))
}

func registerAdminAPI(appContainer app.App, router fiber.Router, cfg config.Config) {
	adminRouter := router.Group("/admin")
	userService := service.NewUserService(
		appContainer.UserService(context.Background()),
		cfg.Server.Secret,
		cfg.Server.AuthExpMinute,
		cfg.Server.AuthRefreshMinute)

	authMiddleware := newAuthMiddleware([]byte(cfg.Server.Secret))

	adminRouter.Post("/terminal",
		authMiddleware,
		rolesAccessMiddleware([]string{userDomain.RoleAdmin}), 
		CreateTerminal(cfg.RoutemapService),
	)
	adminRouter.Post("/route",
		authMiddleware,
		rolesAccessMiddleware([]string{userDomain.RoleAdmin}), 
		CreateRoute(cfg.RoutemapService),
	)
	adminRouter.Put("/users/:id/role", 
		authMiddleware,
		rolesAccessMiddleware([]string{userDomain.RoleAdmin}), 
		UpdateRole(userService),
	)
}

func registerRoutemapAPI(router fiber.Router, cfg config.Config) {
	routemapRouter := router.Group("/routemap")

	routemapRouter.Get("/terminal",
		newAuthMiddleware([]byte(cfg.Server.Secret)),
		GetTerminal(cfg.RoutemapService),
	)
	routemapRouter.Get("/route",
		newAuthMiddleware([]byte(cfg.Server.Secret)),
		GetRouteByID(cfg.RoutemapService),
	)
}

func registerHotelAPI(appContainer hotel.App, router fiber.Router) {
	hotelService := appContainer.HotelService()

	hotelHandler := NewHotelHandler(hotelService)
	router.Post("/hotels/upsert", hotelHandler.RegisterHotelHandler)
	router.Get("/hotels/get-all", hotelHandler.GetAllHotelsHandler)
	router.Get("/hotels/get-one/:id", hotelHandler.GetHotelByIDHandler)
	router.Delete("/hotels/delete/:id", hotelHandler.DeleteHotelHandler)
	router.Post("/rooms/upsert", hotelHandler.CreateOrUpdateRoom)
	router.Get("/rooms/get-one/:id", hotelHandler.GetRoomByID)
	router.Get("/rooms/get-one-by-hotelId/:hotel_id", hotelHandler.GetRoomsByHotelID)
	router.Delete("/rooms/delete/:id", hotelHandler.DeleteRoom)
	router.Post("/rooms/book-hotel", hotelHandler.CreateBooking)
	router.Get("/rooms/booking-detail/:id", hotelHandler.GetBookingByID)
	router.Get("/rooms/booking-detail-by-userId/:user_id", hotelHandler.GetBookingsByUserID)
	router.Post("/rooms/cancel-booking/:id", hotelHandler.SoftDeleteBooking)
	router.Post("/rooms/confirm-booking/:id", hotelHandler.ConfirmBooking)
}

func registerBankAPI(appContainer bank.App, router fiber.Router) {
	bankService := appContainer.BankService()

	bankHandler := NewBankHandler(bankService)
	router.Post("/bank/wallet", bankHandler.CreateWallet)
	router.Post("/bank/charge-wallet", bankHandler.ChargeWalletHandler)
	router.Post("/bank/process-unconfirmed-claim", bankHandler.ProcessUnconfirmedClaim)
	router.Post("/bank/process-confirmed-claim/:claim_id", bankHandler.ProcessConfirmedClaimHandler)
}

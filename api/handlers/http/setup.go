package http

import (
	"context"
	"fmt"
	"qolibaba/api/service"
	"qolibaba/app"
	"qolibaba/app/hotel"

	// "qolibaba/app/admin"
	"qolibaba/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, serverCfg config.ServerConfig, adminCfg config.AdminServiceConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerAuthAPI(appContainer, serverCfg, api)
	registerAdminAPI(api, adminCfg)

	return router.Listen(fmt.Sprintf(":%d", serverCfg.HttpPort))
}

func RunHotel(appContainer hotel.App, serverCfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerHotelAPI(appContainer, api)

	return router.Listen(fmt.Sprintf(":%d", serverCfg.HttpPort))
}

func registerAuthAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	userPortService := appContainer.UserService(context.Background())
	userService := service.NewUserService(userPortService, cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)
	router.Post("/sign-up", SignUp(userService))
	router.Post("/sign-in", SingIn(userService))
	router.Get("/test", newAuthMiddleware([]byte(cfg.Secret)), TestHandler)
	// userSvcGetter := userServiceGetter(appContainer, cfg)
	// router.Post("/sign-up", setTransaction(appContainer.DB()), SignUp(userSvcGetter))
	// router.Get("/send-otp", setTransaction(appContainer.DB()), SendSignInOTP(userSvcGetter))
	// router.Post("/sign-in", setTransaction(appContainer.DB()), SignIn(userSvcGetter))
	// router.Get("/test", newAuthMiddleware([]byte(cfg.Secret)), TestHandler)
}

func registerHotelAPI(appContainer hotel.App, router fiber.Router) {
	//TODO fix it. add the other routes.
	hotelService := appContainer.HotelService()

	hotelHandler := NewHotelHandler(hotelService)
	router.Post("/hotels/upsert", hotelHandler.RegisterHotelHandler)
	router.Get("/hotels/get-all", hotelHandler.GetAllHotelsHandler)
	router.Get("/hotels/get-one/:id", hotelHandler.GetHotelByIDHandler)
	router.Post("/rooms/upsert", hotelHandler.CreateOrUpdateRoom)
}

func registerAdminAPI(router fiber.Router, cfg config.AdminServiceConfig) {
	adminRouter := router.Group("/admin")

	adminRouter.Post("/say-hello", SayHello(cfg))
}

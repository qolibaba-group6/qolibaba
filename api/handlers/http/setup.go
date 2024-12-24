package http

import (
	"context"
	"fmt"
	"qolibaba/api/service"
	"qolibaba/app"
	// "qolibaba/app/admin"
	"qolibaba/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerAuthAPI(appContainer, cfg, api)
	registerAdminAPI(api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerAuthAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	userService := appContainer.UserService(context.Background())
	router.Post("/sign-up", SignUp(service.NewUserService(userService,
		cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)))
	// userSvcGetter := userServiceGetter(appContainer, cfg)
	// router.Post("/sign-up", setTransaction(appContainer.DB()), SignUp(userSvcGetter))
	// router.Get("/send-otp", setTransaction(appContainer.DB()), SendSignInOTP(userSvcGetter))
	// router.Post("/sign-in", setTransaction(appContainer.DB()), SignIn(userSvcGetter))
	// router.Get("/test", newAuthMiddleware([]byte(cfg.Secret)), TestHandler)
}


func registerAdminAPI(router fiber.Router) {
	adminRouter := router.Group("/admin")

	adminRouter.Post("/say-hello", SayHello())
}
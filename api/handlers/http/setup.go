package http

import (
	"context"
	"fmt"
	"qolibaba/api/service"
	"qolibaba/app"

	"qolibaba/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.Config) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerAuthAPI(appContainer, cfg.Server, api)
	registerAdminAPI(api, cfg)

	return router.Listen(fmt.Sprintf(":%d", cfg.Server.HttpPort))
}

func registerAuthAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	userPortService := appContainer.UserService(context.Background())
	userService := service.NewUserService(userPortService, cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)
	router.Post("/sign-up", SignUp(userService))
	router.Post("/sign-in", SingIn(userService))
	router.Get("/test", newAuthMiddleware([]byte(cfg.Secret)), TestHandler)
}


func registerAdminAPI(router fiber.Router, cfg config.Config) {
	adminRouter := router.Group("/admin")

	adminRouter.Post("/say-hello", SayHello(cfg.AdminService))
	adminRouter.Post("/terminal", 
		newAuthMiddleware([]byte(cfg.Server.Secret)), 
		adminAccessMiddleware, 
		CreateTerminal(cfg.RoutemapService),
	)
}
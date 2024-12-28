package main

import (
	"log"

	"user-panel/internal/middleware"
	"user-panel/pkg/services"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	// Apply Metrics Middleware
	r.Use(MetricsMiddleware())

	// Expose metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router := gin.Default()

	// Middleware
	router.Use(middleware.AuditMiddleware())

	// Routes
	router.POST("/auth/login", services.Login)
	router.GET("/user/activities", middleware.AuthMiddleware("USER"), services.GetUserActivities)
	router.GET("/user/wallet", middleware.AuthMiddleware("USER"), services.GetWalletBalance)
	router.POST("/user/wallet/charge", middleware.AuthMiddleware("USER"), services.ChargeWallet)
	router.GET("/user/notifications", middleware.AuthMiddleware("USER"), services.GetNotifications)

	// Add other routes here
	log.Println("User Panel Service is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

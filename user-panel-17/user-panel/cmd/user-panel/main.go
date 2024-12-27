package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "user-panel/internal/middleware"
    "user-panel/pkg/services"
)

func main() {
    router := gin.Default()

    // Middleware
    router.Use(middleware.AuditMiddleware())

    // Routes
    router.POST("/auth/login", services.Login)
    router.GET("/user/activities", middleware.AuthMiddleware("USER"), services.GetUserActivities)
    router.GET("/user/wallet", middleware.AuthMiddleware("USER"), services.GetWalletBalance)
    router.POST("/user/wallet/charge", middleware.AuthMiddleware("USER"), services.ChargeWallet)
    router.GET("/user/notifications", middleware.AuthMiddleware("USER"), services.GetNotifications)

    log.Println("User Panel Service is running on port 8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}


package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"companies-service/internal/config"
	"companies-service/internal/db"
	"companies-service/internal/handlers"
	"companies-service/internal/repository"
	"companies-service/internal/routes"
	"companies-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from .env
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to the database
	dbConn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	// Set up Gin
	router := gin.Default()

	// Add CORS support
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Initialize repositories, services, and handlers
	companyRepo := repository.NewCompanyRepository(dbConn)
	companyService := service.NewCompanyService(companyRepo)
	companyHandler := handlers.NewCompanyHandler(companyService)

	// Define routes
	routes.RegisterCompanyRoutes(router, companyHandler)

	// Create a server with graceful shutdown support
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}

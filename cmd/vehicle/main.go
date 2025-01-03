package main

import (
	"log"
	"qolibaba/internal/adapters/database"
	"qolibaba/internal/adapters/http"
	"qolibaba/internal/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.InitDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize router and start server
	router := http.NewGinRouter()
	router.Start(cfg.ServerAddress, db)
}

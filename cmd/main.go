package main

import (
	"flag"
	"log"
	"os"
	"qolibaba/api/handlers/http"
	"qolibaba/app"
	"qolibaba/config"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.NewMustApp(c)

	log.Fatal(http.Run(appContainer, c.Server))
}


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

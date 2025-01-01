package main

import (
	"flag"
	"log"
	"os"
	"qolibaba/api/handlers/http"
	"qolibaba/app/hotel"
	"qolibaba/app/travel_agency"
	"qolibaba/config"
)

var (
	configPath = flag.String("config", "config.json", "service configuration file")
)

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	cfg := config.MustReadConfig(*configPath)

	// Initialize the Hotel service
	hotelApp, err := hotel.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize hotel app: %v", err)
	}

	// Initialize the Travel Agency service
	travelAgencyApp, err := travel_agency.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize travel agency app: %v", err)
	}

	// Run HTTP server for both services (Hotel and Travel Agency)
	go func() {
		err := http.RunHotel(hotelApp, cfg.Server)
		if err != nil {
			log.Fatalf("failed to start HTTP server for hotel: %v", err)
		}
	}()

	go func() {
		// err := http.RunTravelAgency(travelAgencyApp, cfg.Server)
		if err != nil {
			log.Fatalf("failed to start HTTP server for travel agency: %v", err)
		}
	}()

	// Block indefinitely to keep the main function running
	select {}
}

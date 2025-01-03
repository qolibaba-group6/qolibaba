package main

import (
	"flag"
	"log"
	"os"
	"qolibaba/api/handlers/http"
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

	travelAgencyApp, err := travel_agency.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize travel agency app: %v", err)
	}

	go func() {
		err := http.RunAgencies(travelAgencyApp, cfg.Server)
		if err != nil {
			log.Fatalf("failed to start HTTP server for travel agency: %v", err)
		}
	}()

	select {}
}

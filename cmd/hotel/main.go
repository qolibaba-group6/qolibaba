package main

import (
	"flag"
	"log"
	"os"
	"qolibaba/api/handlers/http"
	"qolibaba/app/hotel"
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

	hotelApp, err := hotel.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize hotel app: %v", err)
	}

	err = http.Run(hotelApp, cfg.ServerConfig)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

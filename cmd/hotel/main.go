package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"qolibaba/api/handlers/http"
	"qolibaba/app/hotel"
	"qolibaba/config"

	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()

	hotelService := hotelApp.HotelService()
	hotelHandler := http.NewHotelHandler(hotelService)

	app.Post("/api/upsertHotel", hotelHandler.RegisterHotelHandler)
	app.Get("/api/getAllHotels", hotelHandler.GetAllHotelsHandler)
	app.Get("/api/getHotel/:id", hotelHandler.GetHotelByIDHandler)
	app.Post("/upsertRoom", hotelHandler.CreateOrUpdateRoom)

	log.Printf("Hotel service listening on port %d", cfg.HotelService.Port)
	if err := app.Listen(fmt.Sprintf(":%d", cfg.HotelService.Port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

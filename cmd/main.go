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

	cfg := config.MustReadConfig(*configPath)

	appContainer := app.NewMustApp(cfg)

	log.Fatal(http.Run(appContainer, cfg))
}

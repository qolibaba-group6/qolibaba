package main

import (
	"flag"
	"log"
	"os"
	"qolibaba/api/handlers/http"
	"qolibaba/app/bank"
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

	bankApp, err := bank.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize bank app: %v", err)
	}

	err = http.RunBank(bankApp, cfg.Server)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/your_project/api/handlers/grpc"
	"github.com/your_project/api/handlers/http"
	"github.com/your_project/config"
	"github.com/your_project/internal/travel"
	"github.com/your_project/pkg/logger"
	"github.com/your_project/pkg/messaging"
	"github.com/your_project/pkg/postgres"
	"google.golang.org/grpc"
)

func main() {
	// بارگذاری تنظیمات از فایل
	cfg, err := config.LoadConfig("config/config-sample.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// راه‌اندازی لاگ
	log := logger.New(cfg.Logging.Level)

	// اتصال به پایگاه داده PostgreSQL
	db, err := postgres.NewPostgresDB(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// راه‌اندازی سرویس‌ها
	travelRepo := travel.NewGormTravelRepo(db)
	travelService := travel.NewTravelService(travelRepo, log)

	// راه‌اندازی گیت‌وی HTTP
	httpHandler := http.NewTravelHandler(travelService)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: httpHandler,
	}

	// راه‌اندازی سرویس gRPC
	grpcServer := grpc.NewServer()
	grpcHandler := grpc.NewTravelHandler(travelService)
	grpc.RegisterTravelServiceServer(grpcServer, grpcHandler)

	// راه‌اندازی کانال پیام‌رسانی RabbitMQ (اختیاری)
	rabbitMQ := messaging.NewRabbitMQ(cfg.AppName)
	go rabbitMQ.Start()

	// اجرای سرور HTTP به صورت جداگانه
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}()

	// اجرای سرور gRPC
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Error starting gRPC server: %v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error running gRPC server: %v", err)
		}
	}()

	// نگه داشتن برنامه در حال اجرا
	select {}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	httpHandler "github.com/ehsansobhani/project_structure-3/api/handlers/http" // استفاده از alias
	"github.com/ehsansobhani/project_structure-3/app/user"
	"github.com/ehsansobhani/project_structure-3/config"
	"github.com/ehsansobhani/project_structure-3/pkg/adapter/storage"
	"github.com/ehsansobhani/project_structure-3/pkg/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// بارگذاری تنظیمات
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("[Error] Loading config: %v", err)
	}

	// ایجاد رشته اتصال به دیتابیس
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	// اتصال به دیتابیس با استفاده از Gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[Error] Connecting to database: %v", err)
	}
	log.Println("[Info] Database connected successfully.")

	// ایجاد repository
	repo := storage.NewUserRepository(db)

	// تنظیم کلید JWT از فایل پیکربندی
	jwtSecret := []byte(cfg.JWTSecret)
	if len(jwtSecret) == 0 {
		log.Fatalf("[Error] JWT secret key is not set in configuration")
	}

	// ایجاد سرویس کاربر
	userService := user.NewUserService(repo, jwtSecret)

	// ایجاد هاندلرها
	handler := httpHandler.NewUserHandler(userService)

	// استفاده از میدلورها و تنظیم مسیرها
	mux := http.NewServeMux()
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)
	mux.Handle("/profile", middleware.AuthMiddleware(handler.GetProfile, jwtSecret))
	mux.Handle("/update", middleware.AuthMiddleware(handler.UpdateProfile, jwtSecret))
	mux.Handle("/delete", middleware.AuthMiddleware(handler.DeleteUser, jwtSecret))

	// شروع سرور HTTP
	serverAddress := fmt.Sprintf(":%s", cfg.HTTPPort)
	log.Printf("[Info] Server is running on port %s...", cfg.HTTPPort)
	if err := http.ListenAndServe(serverAddress, mux); err != nil {
		log.Fatalf("[Error] Starting server: %v", err)
	}
}

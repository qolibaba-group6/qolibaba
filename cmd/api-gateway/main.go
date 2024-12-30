// cmd/api-gateway/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"travel-booking-app/internal/config"
	"travel-booking-app/internal/database"
	"travel-booking-app/internal/middleware"
	notificationHandler "travel-booking-app/internal/notification-service/handler"
	notificationModel "travel-booking-app/internal/notification-service/model"
	paymentHandler "travel-booking-app/internal/payment-service/handler"
	paymentModel "travel-booking-app/internal/payment-service/model"
	ticketHandler "travel-booking-app/internal/ticket-service/handler"
	ticketModel "travel-booking-app/internal/ticket-service/model"
	userHandler "travel-booking-app/internal/user-service/handler"
	userModel "travel-booking-app/internal/user-service/model"
	"travel-booking-app/internal/user-service/service"

	"github.com/gorilla/mux"
)

func main() {

	go service.StartGRPCServer()

	// بارگذاری پیکربندی
	config.LoadConfig()

	// اتصال به پایگاه داده
	database.ConnectDatabase()

	// خودکارسازی مهاجرت مدل‌ها
	err := database.DB.AutoMigrate(
		&userModel.User{},
		&ticketModel.Ticket{},
		&paymentModel.Payment{},
		&notificationModel.Notification{},
	)
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	// ایجاد روت‌ها
	router := mux.NewRouter()

	// افزودن میان‌افزارها
	router.Use(middleware.LoggingMiddleware)

	// مسیریابی برای سرویس کاربران
	userH := userHandler.NewUserHandler()
	router.HandleFunc("/users/register", userH.Register).Methods("POST")
	router.HandleFunc("/users/login", userH.Login).Methods("POST")

	// مسیریابی برای سرویس بلیط
	ticketH := ticketHandler.NewTicketHandler()
	router.HandleFunc("/tickets", ticketH.CreateTicket).Methods("POST")
	router.HandleFunc("/tickets/{id}", ticketH.GetTicket).Methods("GET")
	router.HandleFunc("/tickets", ticketH.GetTicketsByUser).Methods("GET")
	router.HandleFunc("/tickets/{id}", ticketH.UpdateTicket).Methods("PUT")
	router.HandleFunc("/tickets/{id}", ticketH.DeleteTicket).Methods("DELETE")

	// مسیریابی برای سرویس پرداخت
	paymentH := paymentHandler.NewPaymentHandler()
	router.HandleFunc("/payments", paymentH.CreatePayment).Methods("POST")
	router.HandleFunc("/payments/{id}", paymentH.GetPayment).Methods("GET")
	router.HandleFunc("/payments", paymentH.GetPaymentsByUser).Methods("GET")
	router.HandleFunc("/payments/{id}/status", paymentH.UpdatePaymentStatus).Methods("PUT")
	router.HandleFunc("/payments/{id}", paymentH.DeletePayment).Methods("DELETE")

	// مسیریابی برای سرویس نوتیفیکیشن
	notificationH := notificationHandler.NewNotificationHandler()
	router.HandleFunc("/notifications", notificationH.CreateNotification).Methods("POST")
	router.HandleFunc("/notifications/{id}", notificationH.GetNotification).Methods("GET")
	router.HandleFunc("/notifications", notificationH.GetNotificationsByUser).Methods("GET")
	router.HandleFunc("/notifications/{id}", notificationH.UpdateNotification).Methods("PUT")
	router.HandleFunc("/notifications/{id}", notificationH.DeleteNotification).Methods("DELETE")

	// شروع سرور
	serverPort := config.AppConfig.Server.Port
	log.Printf("API Gateway is running on port %d", serverPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

// Import the gRPC service package

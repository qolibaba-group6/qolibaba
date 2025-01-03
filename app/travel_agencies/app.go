// app/travel_agencies/app.go
package travel_agencies

import (
	"context"
	"fmt"
	"time"

	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ehsansobhani/travel_agencies/api/handlers/grpc"
	"github.com/ehsansobhani/travel_agencies/api/handlers/http"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/adapter/storage"
	"github.com/ehsansobhani/travel_agencies/internal/travel_agencies/service"
	"github.com/ehsansobhani/travel_agencies/pkg/cache"
	"github.com/ehsansobhani/travel_agencies/pkg/logger"
	"github.com/ehsansobhani/travel_agencies/pkg/messaging"
	"github.com/ehsansobhani/travel_agencies/pkg/middleware"
	"github.com/ehsansobhani/travel_agencies/pkg/postgres"
	"google.golang.org/grpc"
)

type App struct {
	Config             Config
	Logger             *logger.Logger
	DB                 *postgres.GormDB
	RedisCache         *cache.RedisCache
	RabbitMQ           *messaging.RabbitMQ
	CompanyRepo        service.CompanyRepository
	TripRepo           service.TripRepository
	CompanyService     *service.CompanyService
	TripService        *service.TripService
	HTTPHandler        *http.TravelAgenciesHandler
	GRPCServer         *grpc.TravelAgenciesGRPCServer
	Router             *mux.Router
	HTTPServer         *http.Server
	GRPCListener       net.Listener
	GRPCServerInstance *grpc.Server
}

// Config holds the configuration for the application
type Config struct {
	DBDSN         string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	RabbitMQURI   string
	RabbitMQQueue string
	HTTPPort      string
	GRPCPort      string
}

// NewApp initializes the application with all dependencies
func NewApp(config Config) (*App, error) {
	// Initialize Logger
	logger := logger.NewLogger()

	// Connect to Database
	db, err := postgres.NewGormDB(config.DBDSN)
	if err != nil {
		logger.Fatal("Failed to connect to database:", err)
		return nil, err
	}

	// Initialize Cache
	redisCache := cache.NewRedisCache(config.RedisAddr, config.RedisPassword, config.RedisDB)

	// Initialize Messaging
	rabbitMQ := messaging.NewRabbitMQ(config.RabbitMQURI, config.RabbitMQQueue)

	// Initialize Repositories
	companyRepo, tripRepo := storage.NewTravelAgenciesRepo(db)

	// Initialize Services
	companyService := service.NewCompanyService(companyRepo)
	tripService := service.NewTripService(tripRepo, companyService)

	// Initialize Handlers
	httpHandler := http.NewTravelAgenciesHandler(companyService, tripService, logger)
	grpcServer := grpc.NewTravelAgenciesGRPCServer(companyService, tripService)

	// Initialize Router
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware(logger))
	router.Use(middleware.RateLimitMiddleware(100, time.Minute)) // مثال: حداکثر 100 درخواست در دقیقه

	// Define HTTP Routes
	// Company Routes
	router.HandleFunc("/companies", httpHandler.CreateCompany).Methods("POST")
	router.HandleFunc("/companies/{id}", httpHandler.GetCompany).Methods("GET")
	router.HandleFunc("/companies/{id}", httpHandler.UpdateCompany).Methods("PUT")
	router.HandleFunc("/companies/{id}", httpHandler.DeleteCompany).Methods("DELETE")
	router.HandleFunc("/companies", httpHandler.ListCompanies).Methods("GET")

	// Trip Routes
	router.HandleFunc("/trips", httpHandler.CreateTrip).Methods("POST")
	router.HandleFunc("/trips/{id}", httpHandler.GetTrip).Methods("GET")
	router.HandleFunc("/trips/{id}", httpHandler.UpdateTrip).Methods("PUT")
	router.HandleFunc("/trips/{id}", httpHandler.DeleteTrip).Methods("DELETE")
	router.HandleFunc("/trips", httpHandler.ListTrips).Methods("GET")

	// Initialize HTTP Server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.HTTPPort),
		Handler: router,
	}

	// Initialize gRPC Server
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCPort))
	if err != nil {
		logger.Fatal("Failed to listen on gRPC port:", err)
		return nil, err
	}
	grpcSrv := grpc.NewServer()
	pb.RegisterTravelAgenciesServiceServer(grpcSrv, grpcServer)

	return &App{
		Config:             config,
		Logger:             logger,
		DB:                 db,
		RedisCache:         redisCache,
		RabbitMQ:           rabbitMQ,
		CompanyRepo:        companyRepo,
		TripRepo:           tripRepo,
		CompanyService:     companyService,
		TripService:        tripService,
		HTTPHandler:        httpHandler,
		GRPCServer:         grpcServer,
		Router:             router,
		HTTPServer:         httpServer,
		GRPCListener:       grpcListener,
		GRPCServerInstance: grpcSrv,
	}, nil
}

// Run starts the HTTP and gRPC servers
func (app *App) Run() error {
	// Run HTTP server in a separate goroutine
	go func() {
		app.Logger.Info("HTTP server listening on", app.Config.HTTPPort)
		if err := app.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("Failed to start HTTP server:", err)
		}
	}()

	// Run gRPC server
	app.Logger.Info("gRPC server listening on", app.Config.GRPCPort)
	if err := app.GRPCServerInstance.Serve(app.GRPCListener); err != nil {
		app.Logger.Fatal("Failed to start gRPC server:", err)
		return err
	}

	return nil
}

// Shutdown gracefully shuts down the application
func (app *App) Shutdown(ctx context.Context) error {
	// Shutdown HTTP server
	if err := app.HTTPServer.Shutdown(ctx); err != nil {
		app.Logger.Error("HTTP server Shutdown:", err)
		return err
	}

	// Stop gRPC server
	app.GRPCServerInstance.GracefulStop()

	// Close Database connection
	sqlDB, err := app.DB.DB.DB()
	if err != nil {
		app.Logger.Error("Failed to get raw DB connection:", err)
		return err
	}
	if err := sqlDB.Close(); err != nil {
		app.Logger.Error("Failed to close database connection:", err)
		return err
	}

	// Close Redis connection
	if err := app.RedisCache.Client.Close(); err != nil {
		app.Logger.Error("Failed to close Redis connection:", err)
		return err
	}

	// Close RabbitMQ connection
	if err := app.RabbitMQ.Conn.Close(); err != nil {
		app.Logger.Error("Failed to close RabbitMQ connection:", err)
		return err
	}

	app.Logger.Info("Application shutdown successfully")
	return nil
}

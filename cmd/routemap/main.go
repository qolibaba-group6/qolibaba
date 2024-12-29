package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	grpcAPI "qolibaba/api/handlers/grpc"
	"qolibaba/api/pb"
	"qolibaba/api/service"
	"qolibaba/app/routemap"
	"qolibaba/config"

	"google.golang.org/grpc"
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

	app := routemap.NewMustApp(cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.RoutemapService.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	
	svc := grpcAPI.NewRoutemapGRPCServer(*service.NewRoutemapService(app.RoutemapService()))

	pb.RegisterRoutemapServiceServer(grpcServer, svc)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
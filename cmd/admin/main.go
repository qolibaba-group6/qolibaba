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
	"qolibaba/app/admin"
	"qolibaba/config"

	"google.golang.org/grpc"
)

var (
	configPath = flag.String("config", "config.json", "service configuration file")
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	adminApp := admin.NewMustApp(c)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	
	svc := grpcAPI.NewAdminGRPCServer(*service.NewAdminService(adminApp.AdminService()))

	pb.RegisterAdminServiceServer(grpcServer, svc)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
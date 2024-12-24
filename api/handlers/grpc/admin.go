package grpc

import (
	"context"
	"log"
	"qolibaba/api/pb"
	"qolibaba/api/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type adminGRPCApi struct {
	pb.UnimplementedAdminServiceServer
	svc service.AdminService
}

func NewAdminGRPCServer(svc service.AdminService) pb.AdminServiceServer {
	return &adminGRPCApi{
		svc: svc,
	}
}

func (s *adminGRPCApi) SayHello(ctx context.Context, req *pb.AdminSayHelloRequest) (*pb.AdminSayHelloResponse, error) {
	return s.svc.SayHello(ctx, req)
}


func SayHelloClient(ctx context.Context, req *pb.AdminSayHelloRequest) (*pb.AdminSayHelloResponse, error) {
	client, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	asClient := pb.NewAdminServiceClient(client)

	// ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("api-key", "123456"))

	// var (
	// 	trailer metadata.MD
	// )

	res, err := asClient.SayHello(ctx, req)

	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Fatalf("error code : %d, msg : %s", status.Code(), status.Message())
		}
		log.Fatal(err)
	}

	// fmt.Println(trailer.Get("x-user-id"))

	return  res, nil
}
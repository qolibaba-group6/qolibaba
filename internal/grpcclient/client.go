package grpcclient

import (
	"context"
	"log"
	"time"

	pb "final12/proto"

	"google.golang.org/grpc"
)

type CompanyServiceClient struct {
	client pb.CompanyServiceClient
}

func NewCompanyServiceClient(address string) (*CompanyServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewCompanyServiceClient(conn)
	return &CompanyServiceClient{client: client}, nil
}

func (c *CompanyServiceClient) GetCompany(id string) (*pb.CompanyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := c.client.GetCompany(ctx, &pb.CompanyRequest{Id: id})
	if err != nil {
		return nil, err
	}
	log.Printf("Received company: %+v", response)
	return response, nil
}

func (c *CompanyServiceClient) CreateCompany(name, owner string) (*pb.CompanyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := c.client.CreateCompany(ctx, &pb.NewCompany{Name: name, Owner: owner})
	if err != nil {
		return nil, err
	}
	log.Printf("Created company: %+v", response)
	return response, nil
}

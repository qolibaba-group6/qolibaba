package main


        import (
            "google.golang.org/grpc"
            "log"
        )

        

        package grpc

        func ConnectToService(address string) (*grpc.ClientConn, error) {
            conn, err := grpc.Dial(address, grpc.WithInsecure())
            if err != nil {
                log.Printf("Failed to connect to service at %s: %v", address, err)
                return nil, err
            }
            return conn, nil
        }
    
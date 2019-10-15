package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"golang-ddd-boilerplate/application"
	grpc2 "golang-ddd-boilerplate/infrastructure/grpc"
	userPb "golang-ddd-boilerplate/protobuf-files/user"
)

func main() {
	startGrpcServer(":8080")
}

func startGrpcServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	userService := application.NewUserService()
	userPb.RegisterUserServiceServer(grpcServer, grpc2.NewUserServer(userService))

	reflection.Register(grpcServer)

	fmt.Printf("gRPC server started at %v", address)

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

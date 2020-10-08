package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"golang-boilerplate/application"
	infraGRPC "golang-boilerplate/infrastructure/grpc"
	userPB "golang-boilerplate/protobuf-files/user"
)

func main() {
	err := startGRPCServer(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func startGRPCServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	gRPCServer := grpc.NewServer()

	userService := application.NewUserService()
	userPB.RegisterUserServiceServer(gRPCServer, infraGRPC.NewUserServer(userService))

	reflection.Register(gRPCServer)

	fmt.Printf("gRPC server started at %v", address)

	if err := gRPCServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

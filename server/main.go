package main

import (
	"context"
	"log"
	"net"

	pbHello "github.com/water25234/golang-gRPC/protoc/hello"
	pbUser "github.com/water25234/golang-gRPC/protoc/user"
	"google.golang.org/grpc"
)

type helloService struct {
	pbHello.UnimplementedHelloServiceServer
}

type userService struct {
	pbUser.UnimplementedUserServiceServer
}

func (s *helloService) SayHello(ctx context.Context, in *pbHello.HelloRequest) (*pbHello.HelloResponse, error) {
	log.Printf("Received: %v", in.GetGreeting())
	return &pbHello.HelloResponse{Reply: "Hello, " + in.GetGreeting()}, nil
}

func (s *userService) Login(ctx context.Context, in *pbUser.LoginRequest) (*pbUser.LoginResponse, error) {
	log.Printf("Received: %v", map[string]interface{}{
		"userName": in.GetUsername(),
		"password": in.GetPassword(),
	})
	return &pbUser.LoginResponse{
		UserID:   3000,
		Username: in.GetUsername(),
		Password: in.GetPassword(),
		Name:     "Justin Huang",
		Email:    "water@gmail.com",
		Nickname: "Shun",
	}, nil
}

func main() {
	addr := "127.0.0.1:9999"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server listening on", addr)
	gRPCServer := grpc.NewServer()

	// Hello protoc
	pbHello.RegisterHelloServiceServer(gRPCServer, &helloService{})

	// User protoco
	pbUser.RegisterUserServiceServer(gRPCServer, &userService{})

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("User failed to serve: %v", err)
	}
}

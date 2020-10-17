package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pbHello "github.com/water25234/golang-gRPC/protoc/hello"
	pbUser "github.com/water25234/golang-gRPC/protoc/user"
	"google.golang.org/grpc"
)

func main() {
	addr := "127.0.0.1:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Can not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	c := pbHello.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pbHello.HelloRequest{Greeting: "Moto"})
	if err != nil {
		log.Fatalf("Could not get nonce: %v", err)
	}
	fmt.Println("Hello Response: ", r.GetReply())

	userConn := pbUser.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := userConn.Login(ctx, &pbUser.LoginRequest{
		Username: "zxc@gmail.com",
		Password: "qwaszx",
	})
	defer cancel()

	if err != nil {
		log.Fatalf("Could not get nonce: %v", err)
	}
	fmt.Println("User Response: ", map[string]interface{}{
		"userID":   resp.GetUserID(),
		"userName": resp.GetUsername(),
		"password": resp.GetPassword(),
		"email":    resp.GetEmail(),
		"nickName": resp.GetNickname(),
	})
}

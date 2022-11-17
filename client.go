package main

import (
	"context"
	"log"
	"time"

	"awesomeProject5/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8001"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewUserInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Userinfo(ctx, &proto.UserRequest{
		Username: "admin",
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.Msg, r.Code)
}

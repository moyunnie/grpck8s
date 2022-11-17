package main

import (
	"awesomeProject5/model"
	"awesomeProject5/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	//proto.RegisterUserInfoServer(s, &Server{})
	proto.RegisterUserInfoServer(s, &model.Server{})
	//reflection.Register(s)
	defer func() {
		s.Stop()
		err := listen.Close()
		if err != nil {
			return
		}
	}()
	fmt.Println("Serving 8001...")
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

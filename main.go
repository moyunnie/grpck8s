package main

import (
	"awesomeProject5/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	proto.UnimplementedUserInfoServer
}

func (s *server) Userinfo(context.Context, *proto.UserRequest) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		Msg:  "msg 响应码",
		Code: "200",
	}, nil
}
func main() {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterUserInfoServer(s, &server{})
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

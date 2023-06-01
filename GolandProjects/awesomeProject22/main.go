package main

import (
	"awesomeProject22/global"
	"awesomeProject22/initialize"
	"awesomeProject22/proto"
	"awesomeProject22/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	global.GormDb = initialize.InitGorm()
	global.K8sclient, _ = initialize.InitK8s()
	lis, err := net.Listen("tcp", ":8121")
	if err != nil {
		log.Fatal(err)
	}
	// 注册服务到gRPC
	s := grpc.NewServer()
	proto.RegisterK8SServiceServer(s, &service.K8sService{})
	reflection.Register(s)
	// 启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

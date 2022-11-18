package main

import (
	"awesomeProject5/global"
	"awesomeProject5/initial"
	"awesomeProject5/model"
	"awesomeProject5/proto"
	"awesomeProject5/utils"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db := initial.InitDB()
	err := db.AutoMigrate(&model.UserInfo{})
	if err != nil {
		return
	}
	global.GVA_DB = db
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

	cli, err := utils.NewServiceRegister([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	err = cli.Register("user", "127.0.0.1", 8001, 5)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		cli.UnRegister()
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()
	fmt.Println("Serving 8001...")
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}

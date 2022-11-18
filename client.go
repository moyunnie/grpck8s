package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"

	"awesomeProject5/proto"
	"google.golang.org/grpc"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // etcd节点,因为使用的单节点,所以这里只有一个
		DialTimeout: 5 * time.Second,            //超时时间
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[INFO] connect to etcd success")
	defer func(cli *clientv3.Client) {
		err := cli.Close()
		if err != nil {

		}
	}(cli)
	ctxbak, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctxbak, "user")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	var address = ""
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		address = string(ev.Value)
	}
	fmt.Println(address, "这是address")
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

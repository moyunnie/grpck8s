package main

import (
	"awesomeProject22/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// 创建gRPC连接
	// WithInsecure option 指定不启用认证功能
	conn, err := grpc.Dial(":8121", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	// 创建gRPC client
	client := proto.NewK8SServiceClient(conn)
	//pods, err := client.ListPods(context.Background(), &proto.ListPodsRequest{Page: &proto.Page{
	//	Page:     1,
	//	PageSize: 10,
	//}})
	//if err != nil {
	//	fmt.Println(err.Error(), "异常了")
	//	return
	//}
	//for i, pod := range pods.Pods {
	//	fmt.Println(i, pod)
	//}
	space, err := client.ListNameSpace(context.Background(), &proto.ListNameSpaceRequest{Page: &proto.Page{
		Page:     1,
		PageSize: 10,
	}})
	if err != nil {
		return
	}
	fmt.Println(space)
}

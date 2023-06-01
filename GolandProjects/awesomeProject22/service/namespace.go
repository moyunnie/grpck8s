package service

import (
	"awesomeProject22/global"
	"awesomeProject22/proto"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *K8sService) ListNameSpace(ctx context.Context, in *proto.ListNameSpaceRequest) (*proto.ListNameSpaceResponse, error) {
	list, err := global.K8sclient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	//Kubernetes 中，一些预定义的标签 key 有特殊的含义，
	//例如 app.kubernetes.io/name 用于指定应用程序的名称，
	//app.kubernetes.io/version 用于指定应用程序的版本，
	//app.kubernetes.io/part-of 用于指定应用程序所属的部分等。
	var Namespace []*proto.ListNameSpace
	for _, item := range list.Items {
		//fmt.Println(item.Name)
		for s2, s3 := range item.Labels {
			fmt.Println(s2, s3)
		}
		//Namespace = append(Namespace, proto.ListNameSpace{
		//	NameSpace:  item.Name,
		//	Label:      item.Labels,
		//	Status:     item.Status,
		//	CreateTime: item.CreationTimestamp,
		//})
	}
	return &proto.ListNameSpaceResponse{
		ListNameSpace: Namespace,
	}, nil
}

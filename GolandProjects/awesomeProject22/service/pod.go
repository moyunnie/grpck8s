package service

import (
	"awesomeProject22/global"
	"awesomeProject22/proto"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//func (s *K8sService) GetPodAll(page *request.PageInfo) []v1.Pod {
//	// 计算所需的页码和每页显示的数量
//	pageNum := page.Page - 1
//	pageSize := page.PageSize
//	list, err := global.K8sclient.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
//	if err != nil {
//		return nil
//	}
//	// 计算应该返回的 Pod 列表
//	startIndex := pageNum * pageSize
//	endIndex := startIndex + pageSize
//	if endIndex > len(list.Items) {
//		endIndex = len(list.Items)
//	}
//	return list.Items[startIndex:endIndex]
//}

// 分页
func (s *K8sService) ListPods(ctx context.Context, in *proto.ListPodsRequest) (*proto.ListPodsResponse, error) {
	list, err := global.K8sclient.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var listpod proto.ListPodsResponse
	var pod []*proto.Pod
	for _, i2 := range list.Items {
		//for i, v := range i2.Status.Reason {
		//	fmt.Println(i, v)
		//}
		//for s2, s3 := range i2.Labels {
		//	fmt.Println(s2, s3)
		//}
		//for s2, s3 := range i2.Annotations {
		//	fmt.Println(s2, s3)
		//}
		//for i, field := range i2.Spec {
		//	fmt.Println(i, field.Manager)
		//}
		//fmt.Println(i2.Spec.Hostname)
		//fmt.Println(i2.Spec.Containers)
		for _, container := range i2.Spec.Containers {
			//fmt.Println(i, container.Name, container.Image, container.Ports)
			for i3, port := range container.Ports {
				fmt.Println(i3, port.HostPort, port.ContainerPort)
			}
		}
		pod = append(pod, &proto.Pod{
			Name:      i2.Name,
			Namespace: i2.Namespace,
		})
	}
	// 计算应该返回的 Pod 列表
	startIndex := in.Page.Page * in.Page.PageSize
	endIndex := startIndex + in.Page.PageSize
	if int(endIndex) > len(list.Items) {
		endIndex = uint32(len(list.Items))
	}
	listpod.Pods = pod[startIndex:endIndex]
	return &listpod, nil
}

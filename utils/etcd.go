package utils

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type ServiceRegister struct {
	etcdClient *clientv3.Client
	leaseID    clientv3.LeaseID
}

func NewServiceRegister(endpoints []string) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	serviceRegister := &ServiceRegister{
		etcdClient: cli,
	}
	return serviceRegister, nil
}

func (sr *ServiceRegister) Register(serviceName, serviceHost string, servicePort uint, refreshSeconds int64) error {
	//创建租约
	lease, err := sr.etcdClient.Grant(context.Background(), refreshSeconds*3)
	if err != nil {
		return err
	}
	sr.leaseID = lease.ID
	key := fmt.Sprintf("%s", serviceName)
	value := fmt.Sprintf("%s:%d", serviceHost, servicePort)
	//设置key,value
	_, err = sr.etcdClient.Put(context.Background(), key, value, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	//自动续约
	keepaliveChan, err := sr.etcdClient.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		return err
	}
	str := fmt.Sprintf("注册服务完成，服务名称：%s,服务地址：%s,服务端口：%d", serviceName, serviceHost, servicePort)
	fmt.Println(str)
	go func() {
		for ka := range keepaliveChan {
			fmt.Println("续约:", ka.TTL, time.Now())
		}
	}()

	return nil
}

func (sr *ServiceRegister) UnRegister() error {
	if _, err := sr.etcdClient.Revoke(context.Background(), sr.leaseID); err != nil {
		return err
	}
	fmt.Println("服务注销")
	return sr.etcdClient.Close()
}

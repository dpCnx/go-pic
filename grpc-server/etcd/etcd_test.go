package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"testing"
	"time"
)

/*const GrpcServer = "GrpcServer"

var (
	client *clientv3.Client
	err    error
)*/

func TestMain(m *testing.M) {

	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		zap.L().Error(err.Error())
	}

	m.Run()
}

func TestRegisterServerToEtcd(t *testing.T) {

	kv := clientv3.NewKV(client)

	_, err = kv.Put(context.Background(), "/"+GrpcServer+"/127.0.0.1", "127.0.0.1")
	if err != nil {
		return
	}

	return
}

func TestUnRegisterServerToEtcd(t *testing.T) {

	kv := clientv3.NewKV(client)

	_, err = kv.Delete(context.Background(), "/"+GrpcServer+"/127.0.0.1")
	if err != nil {
		return
	}

	return
}

func TestGetServerToEtcd(t *testing.T) {

	kv := clientv3.NewKV(client)

	getRespose, err := kv.Get(context.Background(), "/"+GrpcServer+"/", clientv3.WithPrefix())
	if err != nil {
		return
	}

	fmt.Println(getRespose.Kvs)

	return
}

/*
	etcd 依赖问题
	go mod edit -replace github.com/coreos/bbolt@v1.3.4=go.etcd.io/bbolt@v1.3.4
	go mod edit -replace google.golang.org/grpc@v1.29.1=google.golang.org/grpc@v1.26.0
	go mod tidy
*/

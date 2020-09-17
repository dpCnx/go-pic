package etcd

import (
	"context"
	"fmt"
	"go-pic/conf"
	"go.etcd.io/etcd/clientv3"
	"time"
)

const GrpcServer = "GrpcServer"

var (
	client *clientv3.Client
	err    error
)

func init() {

	config := clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%s", conf.C.Etcd.IP, conf.C.Etcd.Port)},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		panic(err)
	}

}

func GetServerToEtcd() (getRespose *clientv3.GetResponse, err error) {

	kv := clientv3.NewKV(client)

	getRespose, err = kv.Get(context.Background(), "/"+GrpcServer+"/", clientv3.WithPrefix())
	if err != nil {
		return
	}

	return
}

package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"grpc-server/conf"
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

	if client == nil {
		panic("etcd client is nil")
	}

}

func RegisterServerToEtcd(addr string) (err error) {

	kv := clientv3.NewKV(client)

	_, err = kv.Put(context.Background(), "/"+GrpcServer+"/"+addr, addr)
	if err != nil {
		return
	}


	return
}

func UnRegisterServerToEtcd(addr string) (err error) {

	kv := clientv3.NewKV(client)

	_, err = kv.Delete(context.Background(), "/"+GrpcServer+"/"+addr)
	if err != nil {
		return
	}

	return
}

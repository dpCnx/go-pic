package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.uber.org/zap"
	"time"
	"web/balance"
	"web/conf"
)

const GrpcServer = "GrpcServer"

var (
	client *clientv3.Client
	err    error
	RR     balance.RoundRobinBalance

	grpcServers map[string]string //server集合
)

func init() {

	config := clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%s", conf.C.Etcd.IP, conf.C.Etcd.Port)},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		zap.L().Error(err.Error())
		return
	}

	if client == nil {
		zap.L().Error(err.Error())
		return
	}

	etcdRespose, err := getServerToEtcd()
	if err != nil {
		zap.L().Error(err.Error())
		return
	}

	grpcServers = make(map[string]string)
	RR = balance.RoundRobinBalance{}

	for _, v := range etcdRespose.Kvs {
		grpcServers[string(v.Key)] = string(v.Value)
		if err = RR.Add(string(v.Value)); err != nil {
			zap.L().Error(err.Error())
			continue
		}
	}

	fmt.Println("init server:", RR.GetAll())

	go watchServer()
}

func getServerToEtcd() (getRespose *clientv3.GetResponse, err error) {

	kv := clientv3.NewKV(client)

	if kv == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	getRespose, err = kv.Get(ctx, "/"+GrpcServer+"/", clientv3.WithPrefix())
	if err != nil {
		return
	}

	return
}

func watchServer() {

	watcher := clientv3.NewWatcher(client)

	watchRespChan := watcher.Watch(context.Background(), "/"+GrpcServer+"/", clientv3.WithPrefix())

	// 处理kv变化事件
	for watchResp := range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:

				grpcServers[string(event.Kv.Key)] = string(event.Kv.Value)

			case mvccpb.DELETE:

				delete(grpcServers, string(event.Kv.Key))

			}

			//更新server

			RR.Reset()

			for _, v := range grpcServers {
				if err = RR.Add(v); err != nil {
					zap.L().Error(err.Error())
					continue
				}
			}

			fmt.Println("now server:", RR.GetAll())
		}
	}
}

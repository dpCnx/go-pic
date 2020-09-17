package models

import (
	"context"
	"fmt"
	"go-pic/conf"
	"go-pic/etcd"
	picproto "go-pic/proto"
	"google.golang.org/grpc"
	"time"
)

type Pic struct {
}

func (p *Pic) QuarkPic(page, pagesize int) (pic []string, err error) {

	etcdRespose, err := etcd.GetServerToEtcd()
	if err != nil {
		return
	}

	for _, v := range etcdRespose.Kvs {
		fmt.Println(string(v.Value))
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.C.Grpc.Ip, conf.C.Grpc.Port), grpc.WithInsecure())

	if err != nil {
		return
	}

	defer conn.Close()

	c := picproto.NewPicServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.GetPics(ctx, &picproto.RequestPic{
		Page:     int32(page),
		PageSize: int32(pagesize),
	})

	if err != nil {
		return
	}

	pic = res.Pics

	return
}

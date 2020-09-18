package models

import (
	"context"
	"go-pic/etcd"
	picproto "go-pic/proto"
	"google.golang.org/grpc"
	"time"
)

type Pic struct {
}

func (p *Pic) QuarkPic(page, pagesize int) (pic []string, err error) {

	conn, err := grpc.Dial(etcd.RR.Get(), grpc.WithInsecure())

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

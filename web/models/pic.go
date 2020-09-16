package models

import (
	"context"
	"fmt"
	"go-pic/conf"
	pic_proto "go-pic/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

var (
	conn *grpc.ClientConn
	c    pic_proto.PicServerClient
	err  error
)

func init() {

	conn, err = grpc.Dial(fmt.Sprintf("%s:%s", conf.C.Grpc.Ip, conf.C.Grpc.Port), grpc.WithInsecure())

	if err != nil {
		zap.L().Error(err.Error())
		return
	}

	c = pic_proto.NewPicServerClient(conn)
}

func Colse() {

	if conn != nil {
		conn.Close()
	}

}

type Pic struct {
}

func (p *Pic) QuarkPic(page, pagesize int) (pic []string, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.GetPics(ctx, &pic_proto.RequestPic{
		Page:     int32(page),
		PageSize: int32(pagesize),
	})

	if err != nil {
		return
	}

	pic = res.Pics

	return
}

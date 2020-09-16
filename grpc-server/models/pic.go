package models

import (
	"context"
	proto "grpc-server/proto"
	"time"
)

type PicServer struct{}

func (p *PicServer) GetPics(ctx context.Context, request *proto.RequestPic) (*proto.ResposePic, error) {

	time.Sleep(10 * time.Second)

	return &proto.ResposePic{
		Pics: []string{"A", "B"},
	}, nil
}

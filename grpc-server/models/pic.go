package models

import (
	"context"
	proto "grpc-server/proto"
)

type PicServer struct{}

func (p *PicServer) GetPics(ctx context.Context, request *proto.RequestPic) (*proto.ResposePic, error) {

	return &proto.ResposePic{
		Pics: []string{"A", "B"},
	}, nil
}

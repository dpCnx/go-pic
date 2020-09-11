package models

import (
	"context"
	"errors"
	proto "grpc-server/proto"
)

type PicServer struct{}

func (p *PicServer) GetPics(ctx context.Context, request *proto.RequestPic) (*proto.ResposePic, error) {

	if ctx.Err() == context.Canceled {
		return nil, errors.New("PicServer canceled")
	}

	return &proto.ResposePic{
		Pics: []string{"A", "B"},
	}, nil
}

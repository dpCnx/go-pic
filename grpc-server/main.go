package main

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	_ "grpc-server/log"
	"grpc-server/models"
	proto "grpc-server/proto"
	"net"
)

func main() {

	zap.L().Info("pic-server start...")

	lis, err := net.Listen("tcp", "127.0.0.1:9999")

	if err != nil {
		zap.L().Fatal(err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterPicServerServer(s, &models.PicServer{})
	if err = s.Serve(lis); err != nil {
		zap.L().Fatal(err.Error())
	}
}

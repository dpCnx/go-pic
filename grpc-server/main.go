package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"grpc-server/conf"
	_ "grpc-server/log"
	"grpc-server/models"
	proto "grpc-server/proto"
	"net"
)

func main() {

	zap.L().Info("pic-server start...")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.C.Server.Ip, conf.C.Server.Port))

	if err != nil {
		zap.L().Fatal(err.Error())
	}

	s := grpc.NewServer()
	proto.RegisterPicServerServer(s, &models.PicServer{})
	if err = s.Serve(lis); err != nil {
		zap.L().Fatal(err.Error())
	}
}

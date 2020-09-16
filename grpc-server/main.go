package main

import (
	"fmt"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"grpc-server/conf"
	_ "grpc-server/log"
	"grpc-server/middleware"
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

	opts := []grpc.ServerOption{
		grpcmiddleware.WithUnaryServerChain(
			middleware.Recovery,
			middleware.RequestLog,
			middleware.Hystrix,
		),
	}

	s := grpc.NewServer(opts...)
	proto.RegisterPicServerServer(s, &models.PicServer{})
	if err = s.Serve(lis); err != nil {
		zap.L().Fatal(err.Error())
	}
}

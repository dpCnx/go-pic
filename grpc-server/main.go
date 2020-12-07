package main

import (
	"fmt"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"grpc-server/conf"
	"grpc-server/etcd"
	_ "grpc-server/log"
	"grpc-server/middleware"
	"grpc-server/models"
	proto "grpc-server/proto"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	zap.L().Info("pic-server start...")

	addr := fmt.Sprintf("%s:%s", conf.C.Server.Ip, conf.C.Server.Port)

	zap.L().Info(addr)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		zap.L().Fatal(err.Error())
	}

	opts := []grpc.ServerOption{
		grpcmiddleware.WithUnaryServerChain(
			middleware.Recovery,
			middleware.RequestLog,
		),
	}

	s := grpc.NewServer(opts...)
	proto.RegisterPicServerServer(s, &models.PicServer{})

	//注册到etcd
	if err := etcd.RegisterServerToEtcd(addr); err != nil {
		zap.L().Fatal(err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {

		<-quit

		_ = etcd.UnRegisterServerToEtcd(addr)

		zap.L().Info("pic-server stop...")
		s.Stop()

	}()

	if err = s.Serve(lis); err != nil {
		zap.L().Fatal(err.Error())
	}
}

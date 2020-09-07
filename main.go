package main

import (
	"context"
	"fmt"
	"go-pic/conf"
	_ "go-pic/log"
	"go-pic/models"
	"go-pic/router"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	r := router.LoadRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.C.Server.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("正在关闭go-pic...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error(err.Error())
		return
	}

	models.CloseRdb()

	zap.L().Info("退出go-pic...")
}

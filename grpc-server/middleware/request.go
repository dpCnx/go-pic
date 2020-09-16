package middleware

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func RequestLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	zap.L().Info("request",
		zap.String("method", info.FullMethod),
		zap.String("req", fmt.Sprint(req)),
	)

	resp, err := handler(ctx, req)

	cost := time.Since(start)
	zap.L().Info("respose",
		zap.String("method", info.FullMethod),
		zap.String("resp", fmt.Sprint(resp)),
		zap.Duration("cost", cost),
	)

	return resp, err
}

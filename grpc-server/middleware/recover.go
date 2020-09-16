package middleware

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"runtime/debug"
)

func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {

			zap.L().Error("recovery",
				zap.String("error", fmt.Sprint(e)),
				zap.String("stack", string(debug.Stack())),
			)
			err = errors.New(fmt.Sprintf("err:%v", e))
			return
		}
	}()

	return handler(ctx, req)
}

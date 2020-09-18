package middleware

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	hystrix.ConfigureCommand("grpc", hystrix.CommandConfig{
		Timeout:                1000, // 单次请求 超时时间
		MaxConcurrentRequests:  1,    // 最大并发量
		SleepWindow:            5000, // 熔断后多久去尝试服务是否可用
		RequestVolumeThreshold: 1,    // 验证熔断的 请求数量, 10秒内采样
		ErrorPercentThreshold:  1,    // 验证熔断的 错误百分比
	})

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	/*go func() {
		if err := http.ListenAndServe(net.JoinHostPort("", "2001"), hystrixStreamHandler); err != nil {
			zap.L().Error(err.Error())
		}
	}()*/
}

func Hystrix(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	if err := hystrix.Do("grpc", func() error {
		resp, err = handler(ctx, req)
		if err != nil {
			return err
		}
		return nil
	}, nil); err != nil {
		//加入自动降级处理，如获取缓存数据等

		zap.L().Error(err.Error())

		switch err {
		case hystrix.ErrCircuitOpen:
			zap.L().Error(err.Error())
		case hystrix.ErrMaxConcurrency:
			zap.L().Error(err.Error())
		default:
			zap.L().Error(err.Error())
		}

		return nil, err
	} else {
		return resp, err
	}

}

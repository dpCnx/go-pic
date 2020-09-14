package models

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"grpc-server/conf"
	"time"
)

var (
	redisDb *redis.Client
)

func init() {

	redisDb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.C.Redis.IP, conf.C.Redis.Port),
	})

	_, err := redisDb.Ping().Result()
	if err != nil {
		panic(err)
	}

	zap.L().Info("pic-server redis start...")
}

func RsetString(key string, value interface{}, min int) error {
	return redisDb.Set(key, value, time.Minute*time.Duration(min)).Err()
}

func RgetString(key string) (string, error) {
	return redisDb.Get(key).Result()
}

func RgetInt(key string) (int, error) {
	return redisDb.Get(key).Int()
}

func RdeleteString(key string) (int64, error) {
	return redisDb.Del(key).Result()
}

func RincrString(key string) (int64, error) {
	return redisDb.Incr(key).Result()
}

func Rexpire(key string, min int) error {
	return redisDb.Expire(key, time.Second*time.Duration(min)).Err()
}

func CloseRdb() {
	if redisDb != nil {
		_ = redisDb.Close()
	}
}

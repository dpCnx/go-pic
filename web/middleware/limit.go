package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"time"
)

var rl ratelimit.Limiter

func init() {
	//传入的 10 指的是 1s 内只有能有 10 个请求通过
	rl = ratelimit.New(10)
}

func Limiter() gin.HandlerFunc {
	prev := time.Now()
	return func(c *gin.Context) {
		now := rl.Take()
		fmt.Println(now.Sub(prev)) // 这里不需要, 只是打印下多次请求之间的时间间隔
		//prev = now                 // 这里不需要, 只是打印下多次请求之间的时间间隔
	}
}

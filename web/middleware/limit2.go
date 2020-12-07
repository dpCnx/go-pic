package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
	"web/models"
)

func TokenRateLimiter() gin.HandlerFunc {

	bucket := ratelimit.NewBucket(time.Second, 1)
	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) == 0 {
			models.ResponseErrorWithMsg(c, models.CodeServerBusy, "rate limit")
			c.Abort()
			return
		}
		c.Next()
	}

}

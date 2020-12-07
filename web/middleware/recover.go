package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"runtime/debug"
	"web/models"
)

// GinRecovery recover掉项目可能出现的panic
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("recovery",
					zap.String("error", fmt.Sprint(err)),
					zap.String("stack", string(debug.Stack())),
				)

				models.ResponseError(c, models.CodeServerBusy)
				return
			}
		}()
		c.Next()
	}
}

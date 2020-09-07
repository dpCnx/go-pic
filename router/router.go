package router

import (
	"github.com/gin-gonic/gin"
	"go-pic/conf"
)

func LoadRouter() *gin.Engine {

	gin.SetMode(conf.C.Server.Mode)

	r := gin.Default()

	return r
}

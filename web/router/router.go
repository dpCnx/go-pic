package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"web/conf"
	"web/controller"
	"web/docs"
	"web/middleware"
)

func LoadRouter() *gin.Engine {

	gin.SetMode(conf.C.Server.Mode)

	r := gin.New()
	r.Use(middleware.GinRecovery())
	r.Use(middleware.RequestLog())
	r.Use(middleware.Sentinel())

	//r := gin.Default()

	docs.SwaggerInfo.Title = "Api"
	docs.SwaggerInfo.Description = "api文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:9988"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/getpic", controller.GetPic)
	}

	return r
}

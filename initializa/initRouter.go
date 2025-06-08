package initializa

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/middlewares"
	. "github.com/maolchen/krm-backend/routers"
)

func InitRouter() *gin.Engine {
	// 创建路由
	Router := gin.Default()

	// 使用中间件
	Router.Use(middlewares.LoggerMiddleware())
	Router.Use(middlewares.AuthMiddleware())

	ApiGroup := Router.Group("/api")
	InitAuthRouter(ApiGroup)
	InitUserRouter(ApiGroup)
	InitClusterRouter(ApiGroup)
	return Router

}

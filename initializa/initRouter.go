package initializa

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/project_demo/middlewares"
	. "github.com/maolchen/project_demo/routers"
)

func InitRouter() *gin.Engine {
	// 创建路由
	Router := gin.Default()

	// 使用中间件
	Router.Use(middlewares.RequestTimeMiddleware())
	Router.Use(middlewares.AuthMiddleware())

	ApiGroup := Router.Group("/api")
	InitAuthRouter(ApiGroup)
	InitUserRouter(ApiGroup)
	return Router

}

package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/maolchen/project_demo/config"
	"go.uber.org/zap"
	"log"
	"time"

	"github.com/maolchen/project_demo/initializa"
)

func init() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if err := initializa.InitConfig(configPath); err != nil {
		log.Panicf("初始化配置失败: %v", err)
	}
	// 初始化日志
	initializa.InitZap()
	//打印初始化配置到日志
	initializa.PrintConfig()

	//初始化数据库
	if err := initializa.InitDB(); err != nil {
		zap.L().Error("初始化数据库失败", zap.Error(err))
	}
}

func SlowHandler(c *gin.Context) {
	time.Sleep(1500 * time.Millisecond) // 1.5秒延迟
	c.JSON(200, gin.H{
		"message": "Slow response done.",
	})
}
func main() {

	//初始化router
	//r := gin.Default()
	//
	////使用全局记录请求时间中间件
	//r.Use(middlewares.RequestTimeMiddleware())
	//// 正常接口
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "pong"})
	//})
	//
	//r.GET("/slow", SlowHandler)

	r := initializa.InitRouter()

	zap.S().Infof("启动服务.......,listen: %s", config.Cfg.Address)
	r.Run(config.Cfg.Address)
}

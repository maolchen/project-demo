package main

import (
	"flag"
	"fmt"
	logutil "github.com/maolchen/project_demo/utils/logutils"
	"go.uber.org/zap"
	"log"

	"github.com/maolchen/project_demo/initializa" // 替换为你的模块名
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if err := initializa.Initializa(configPath); err != nil {
		log.Panicf("初始化配置失败: %v", err)
	}
}

func main() {
	cfg := initializa.Cfg
	fmt.Println(*cfg)
	logger := logutil.NewZapLogger(cfg.LogConf)
	defer logger.Sync()
	fmt.Println(logger)
	logger.Info("this is info")
	logger.Error("this is an error")
	logger.Warn("this is a warn")
	logger.Debug("this is a debug")

	zap.S().Infof("Address: %s, 测试地址：%s", cfg.Address, cfg.Address)
	zap.S().Infof("DB Path: %s", cfg.DbPath)
	zap.S().Infof("JWT Secret: %s", cfg.Secret)
	zap.S().Infof("Log File: %s", cfg.LogConf.LogFile)
	zap.S().Infof("Max Age: %d", cfg.LogConf.MaxAge)
	zap.S().Infof("Max Backups: %d", cfg.LogConf.MaxBackups)
	zap.S().Infof("Max Size: %d MB", cfg.LogConf.MaxSize)
	zap.S().Infof("Compress: %t", cfg.LogConf.Compress)

	//fmt.Printf("Local Time: %t\n", cfg.LogConf.LocalTime)
}

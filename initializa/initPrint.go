package initializa

import (
	"github.com/maolchen/project_demo/config"
	"go.uber.org/zap"
)

func PrintConfig() {
	zap.S().Info("当前服务启动配置========================================>")
	zap.S().Infof("Address: %s,", config.Cfg.Address)
	zap.S().Infof("DB Path: %s", config.Cfg.DbPath)
	zap.S().Infof("JWT Secret: %s", config.Cfg.Secret)
	zap.S().Infof("Log File: %s", config.Cfg.LogConf.LogFile)
	zap.S().Infof("Max Age: %d", config.Cfg.LogConf.MaxAge)
	zap.S().Infof("Max Backups: %d", config.Cfg.LogConf.MaxBackups)
	zap.S().Infof("Max Size: %d MB", config.Cfg.LogConf.MaxSize)
	zap.S().Infof("Compress: %t", config.Cfg.LogConf.Compress)
	zap.S().Info("当前服务配置加载完成========================================>")
}

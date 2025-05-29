package initializa

import (
	"github.com/maolchen/project_demo/config"
	logutil "github.com/maolchen/project_demo/utils"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitZap() {
	Logger = logutil.NewZapLogger(config.Cfg.LogConf)
	defer Logger.Sync()
}

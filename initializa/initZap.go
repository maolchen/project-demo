package initializa

import (
	"github.com/maolchen/krm-backend/config"
	logutil "github.com/maolchen/krm-backend/utils"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitZap() {
	Logger = logutil.NewZapLogger(config.Cfg.LogConf)
	defer Logger.Sync()
}

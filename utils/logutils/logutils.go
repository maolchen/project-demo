package logutil

import (
	"fmt"
	"github.com/maolchen/project_demo/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

// 文件切割
func getWriter(cfg *config.LogConf) zapcore.WriteSyncer {

	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})
}
func NewZapLogger(cfg *config.LogConf) *zap.Logger {
	//var cores []zapcore.Core
	var logger *zap.Logger

	// 设置编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	var writer zapcore.WriteSyncer
	// 判断是写控制台还是写日志文件
	if cfg.LogFile != "" {
		fmt.Println("LogFile:", cfg.LogFile)
		//写文件 ，判断是写json格式还是普通格式
		if strings.ToLower(cfg.LogType) == "json" {
			encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 带颜色的大写级别
			writer = getWriter(cfg)
			fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
			fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(writer), parseLevel(cfg.LogLevel))
			logger = zap.New(fileCore, zap.AddCaller())

		} else {
			encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //带颜色的大写级别
			writer = getWriter(cfg)
			fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)
			fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(writer), parseLevel(cfg.LogLevel))
			logger = zap.New(fileCore, zap.AddCaller())
		}
	} else {
		fmt.Println("LogFile:", cfg.LogFile)
		if strings.ToLower(cfg.LogType) == "json" {
			encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
			consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)
			consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), parseLevel(cfg.LogLevel))
			logger = zap.New(consoleCore, zap.AddCaller())
		} else {
			encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
			consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), parseLevel(cfg.LogLevel))
			logger = zap.New(consoleCore, zap.AddCaller())
		}
	}

	zap.ReplaceGlobals(logger)
	return logger
}

// 解析日志级别
func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

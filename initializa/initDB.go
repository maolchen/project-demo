package initializa

import (
	"github.com/maolchen/krm-backend/config"
	"github.com/maolchen/krm-backend/database"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
)

func InitDB() error {

	if err := database.InitORM(config.Cfg); err != nil {
		zap.S().Fatalf("初始化数据库失败: %v", err)
	}

	db := database.GetDB()
	if err := db.AutoMigrate(&models.User{}); err != nil {
		zap.L().Error("初始化用户表结构失败", zap.Error(err))
		return err
	} else {
		zap.L().Info("初始化表结构完成")
	}
	if err := db.AutoMigrate(&models.ClusterInfo{}); err != nil {
		zap.L().Error("初始化集群表结构失败", zap.Error(err))
		return err
	} else {
		zap.L().Info("初始化集群表结构完成")
	}

	zap.L().Info("初始化表结构完成")
	return nil
}

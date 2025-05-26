package initializa

import (
	"github.com/maolchen/project_demo/config"
	"github.com/maolchen/project_demo/controllers/database"
	"github.com/maolchen/project_demo/models"
	"go.uber.org/zap"
)

func InitDB() error {

	if err := database.InitORM(config.Cfg); err != nil {
		zap.S().Fatalf("初始化数据库失败: %v", err)
	}

	db := database.GetDB()
	if err := db.AutoMigrate(&models.User{}); err != nil {
		zap.L().Error("初始化表结构失败", zap.Error(err))
		return err
	}

	zap.L().Info("初始化表结构完成")
	return nil
}

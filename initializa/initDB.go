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

	// 添加一个默认用户
	// 检查默认用户是否存在，如果不存在则创建
	var count int64
	defaultUser := models.User{
		Username: "admin",
		HashPass: "$2a$04$O0824mhYk3hvzSxpLwGp5./Ipb6YSvIif/fju.Ki2yuKBKHsuXtia",
	}
	if db.Model(&models.User{}).Where("username = ?", defaultUser.Username).Count(&count); count == 0 {
		if err := db.Create(&defaultUser).Error; err != nil {
			zap.L().Error("创建默认用户失败", zap.Error(err))
			return err
		}
		zap.L().Info("默认用户创建成功")
	} else {
		zap.L().Info("默认用户已存在")
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

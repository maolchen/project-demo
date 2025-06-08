package database

import (
	"github.com/glebarez/sqlite"
	"github.com/maolchen/krm-backend/config"
	"github.com/maolchen/krm-backend/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitORM 初始化 ORM 并连接数据库，应该在应用启动时调用一次
func InitORM(cfg *config.Conf) error {
	if err := utils.EnsureDirExists(cfg.DbPath); err != nil {
		return err
	}
	var err error
	db, err = gorm.Open(sqlite.Open(cfg.DbPath), &gorm.Config{})
	if err != nil {
		zap.L().Error("数据库连接失败", zap.String("db_path", cfg.DbPath), zap.Error(err))
		return err
	}

	// 可选: 设置数据库连接池参数等
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("获取底层 DB 失败", zap.Error(err))
		return err
	}
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)

	zap.L().Info("数据库连接成功", zap.String("db_path", cfg.DbPath))
	return nil
}

// GetDB 返回已初始化的数据库实例，各模块可以通过此方法获取 db 进行操作
func GetDB() *gorm.DB {
	if db == nil {
		zap.L().Error("数据库尚未初始化，请先调用 InitORM")
	}
	return db
}

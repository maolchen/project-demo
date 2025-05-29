package utils

import (
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

// EnsureDirExists 检查目录是否存在，若不存在则创建之
func EnsureDirExists(filePath string) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		zap.S().Infof("%s目录不存在，正在创建", dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			zap.S().Errorf("%s创建目录失败", dir)
			return err
		}
	}
	return nil
}

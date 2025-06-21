package initializa

import (
	"encoding/base64"
	. "github.com/maolchen/krm-backend/config"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
)

// 从数据库初始化ClusterKubeconfig
func InitClusterKubeconfig() {

	clusters, err := models.GetAllClusters()
	if err != nil {
		zap.L().Error("数据库查询失败", zap.Error(err))
		return
	}
	for _, c := range clusters {
		config, _ := base64.StdEncoding.DecodeString(c.Kubeconfig)
		ClusterKubeconfig[c.Name] = config
	}
}

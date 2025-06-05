package service

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
)

func AddCluster(ctx *gin.Context) error {
	zap.S().Info("添加集群逻辑")
	// 接收数据
	clusterInfo := models.ClusterInfo{}

	if err := ctx.ShouldBindJSON(&clusterInfo); err != nil {
		return err
	}

	// 校验集群状态

	//数据处理
	return clusterInfo.Insert()
	
}

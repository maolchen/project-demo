package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddCluster(ctx *gin.Context) {
	zap.S().Info("添加集群逻辑")
}

package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ClusterAdd(c *gin.Context) {
	zap.S().Info("添加集群操作")
}
func ClusterList(c *gin.Context) {
	zap.S().Info("列出集群操作")
}
func ClusterUpdate(c *gin.Context) {
	zap.S().Info("更新集群操作")
}

func ClusterDelete(c *gin.Context) {
	zap.S().Info("删除集群操作")
}

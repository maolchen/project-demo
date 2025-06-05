package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/api"
)

func InitClusterRouter(r *gin.RouterGroup) {
	r.GET("/cluster/delete", api.ClusterDelete)
	r.GET("/cluster/list", api.ClusterList)
	r.POST("/cluster/update", api.ClusterUpdate)
	r.POST("/cluster/add", api.ClusterAdd)
}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/api"
)

func InitClusterRouter(r *gin.RouterGroup) {
	r.POST("/cluster/delete/:name", api.ClusterDelete)
	r.GET("/cluster/list", api.ClusterList)
	r.POST("/cluster/edit/:name", api.ClusterEditByName)
	r.POST("/cluster/update/:name", api.ClusterUpdate)
	r.POST("/cluster/add", api.ClusterAdd)
}

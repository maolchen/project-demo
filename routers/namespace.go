package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/api"
)

func InitNamespaceRouter(r *gin.RouterGroup) {
	r.GET("/namespace/delete", api.Delete)
	r.GET("/namespace/list", api.List)
	r.GET("/namespace/get", api.Get)
	r.POST("/namespace/update", api.Update)
	r.POST("/namespace/create", api.Create)
}

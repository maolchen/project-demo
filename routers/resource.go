package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/api/resource"
)

func InitResourceRouter(r *gin.RouterGroup) {
	r.GET("/:resource/delete", resource.Delete)
	r.GET("/:resource/list", resource.List)
	r.GET("/:resource/get", resource.Get)
	r.POST("/:resource/update", resource.Update)
	r.POST("/:resource/create", resource.Create)
}

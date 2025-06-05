package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/api"
)

func InitAuthRouter(r *gin.RouterGroup) {
	r.GET("/auth/logout")
	r.POST("/auth/login", api.Login)
}

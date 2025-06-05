package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/api"
)

func InitUserRouter(r *gin.RouterGroup) {
	r.POST("/user/create", api.UserCreate)

}

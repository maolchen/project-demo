package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/project_demo/api"
)

func InitUserRouter(r *gin.RouterGroup) {
	r.POST("/user/create", api.UserCreate)

}

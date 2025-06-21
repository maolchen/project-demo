package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/namespace"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
)

func List(ctx *gin.Context) {
	var basicInfo models.BasicInfo
	//fmt.Println(ctx.Query("clusterName"))
	if !utils.BindQuery(ctx, &basicInfo) {
		return
	}
	//fmt.Println("basicInfo=================", basicInfo)
	nsList, err := namespace.ListNamespace(&basicInfo)
	if err != nil {
		utils.Error(http.StatusInternalServerError, "获取命名空间失败"+":"+err.Error()).Send(ctx)
		return
	}
	utils.SuccessWithItems("获取命名空间成功", nsList).Send(ctx)
}

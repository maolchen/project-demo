package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/namespace"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
)

func Get(ctx *gin.Context) {
	basicInfo := models.BasicInfo{}
	if !utils.BindQuery(ctx, &basicInfo) {
		return
	}

	ns, err := namespace.GetNamespace(&basicInfo)
	if err != nil {
		utils.Error(http.StatusInternalServerError, "获取命名空间详情失败"+":"+err.Error()).Send(ctx)
		return
	}
	utils.SuccessWithItem("获取命名空间详情成功", ns).Send(ctx)
}

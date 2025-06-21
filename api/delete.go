package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/namespace"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
)

func Delete(ctx *gin.Context) {
	basicInfo := models.BasicInfo{}
	if !utils.BindQuery(ctx, &basicInfo) {
		return
	}

	if err := namespace.DeleteNamespace(&basicInfo); err != nil {
		utils.Error(http.StatusInternalServerError, "删除命名空间失败"+":"+err.Error()).Send(ctx)
		return
	}
	utils.SuccessNoData("删除命名空间成功").Send(ctx)
}

package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/factory"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
)

func Get(ctx *gin.Context) {
	basicInfo := models.BasicInfo{}
	if !utils.BindQuery(ctx, &basicInfo) {
		return
	}

	clientset, err := models.NewClientSetForBasicInfo(&basicInfo)
	if err != nil {
		utils.Error(http.StatusInternalServerError, "连接集群失败："+err.Error()).Send(ctx)
		return
	}

	resourceType := ctx.Param("resource") // /api/:resource/create
	svc := (&factory.ResourceFactory{}).GetService(resourceType, clientset)

	result, err := svc.Get(&basicInfo)
	if err != nil {
		utils.Error(http.StatusInternalServerError, "查询"+resourceType+"资源失败："+err.Error()).Send(ctx)
		return
	}
	utils.SuccessWithItem("查询"+resourceType+"命名空间详情成功", result).Send(ctx)
}

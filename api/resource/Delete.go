package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/factory"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
)

func Delete(ctx *gin.Context) {
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

	if err := svc.Delete(&basicInfo); err != nil {
		utils.Error(http.StatusInternalServerError, "删除"+resourceType+"资源失败："+err.Error()).Send(ctx)
		return
	}
	utils.SuccessNoData("删除" + resourceType + "资源成功").Send(ctx)
}

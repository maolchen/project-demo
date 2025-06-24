package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/factory"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
)

func BatchRestart(ctx *gin.Context) {
	var req models.BatchRequest
	if !utils.BindJSON(ctx, &req) {
		return
	}

	clientset, err := models.NewClientSetForBasicInfo(&req.BasicInfo)
	if err != nil {
		utils.Error(http.StatusInternalServerError, "连接集群失败："+err.Error()).Send(ctx)
		return
	}

	resourceType := ctx.Param("resource") // /api/:resource/batchrestart
	svc, ok := (&factory.ResourceFactory{}).GetService(resourceType, clientset).(interface {
		Restart(basicInfo *models.BasicInfo) error
	})
	if !ok {
		utils.Error(http.StatusInternalServerError, "不支持重启操作的资源类型："+resourceType).Send(ctx)
		return
	}

	for _, name := range req.Names {
		req.BasicInfo.Name = name
		if err := svc.Restart(&req.BasicInfo); err != nil {
			utils.Error(http.StatusInternalServerError, "重启资源 "+name+" 失败："+err.Error()).Send(ctx)
			return
		}
	}
	utils.SuccessNoData("批量重启" + resourceType + "成功").Send(ctx)
}

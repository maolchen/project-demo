package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/factory"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
)

func Restart(ctx *gin.Context) {
	basicInfo := models.BasicInfo{}
	if !utils.BindQuery(ctx, &basicInfo) {
		return
	}

	clientset, err := models.NewClientSetForBasicInfo(&basicInfo)
	if err != nil {
		utils.Error(http.StatusInternalServerError, "连接集群失败："+err.Error()).Send(ctx)
		return
	}

	resourceType := ctx.Param("resource")                                                    // /api/:resource/restart
	svc, ok := (&factory.ResourceFactory{}).GetService(resourceType, clientset).(interface { //断言svc是否实现了interface { Restart(basicInfo *models.BasicInfo) error} 接口
		//部分资源是不支持重启操作，只有实现了 Restart接口的才会返回true
		Restart(basicInfo *models.BasicInfo) error
	})
	if !ok {
		utils.Error(http.StatusInternalServerError, "不支持重启操作的资源类型："+resourceType).Send(ctx)
		return
	}

	if err := svc.Restart(&basicInfo); err != nil {
		utils.Error(http.StatusInternalServerError, "重启资源失败："+err.Error()).Send(ctx)
		return
	}
	utils.SuccessNoData("重启资源" + resourceType + ":" + basicInfo.Name + "成功").Send(ctx)
}

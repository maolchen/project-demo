package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/factory"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
)

func BatchDelete(ctx *gin.Context) {
	var req models.BatchRequest
	if !utils.BindJSON(ctx, &req) {
		return
	}
	//fmt.Println(req)
	clientset, err := models.NewClientSetForBasicInfo(&req.BasicInfo)
	if err != nil {
		utils.Error(http.StatusInternalServerError, "连接集群失败："+err.Error()).Send(ctx)
		return
	}

	resourceType := ctx.Param("resource") // /api/:resource/batchdelete
	svc := (&factory.ResourceFactory{}).GetService(resourceType, clientset)

	for _, name := range req.Names {
		req.BasicInfo.Name = name
		if err := svc.Delete(&req.BasicInfo); err != nil {
			utils.Error(http.StatusInternalServerError, "删除资源 "+resourceType+":"+name+" 失败："+err.Error()).Send(ctx)
			return
		}
	}
	utils.SuccessNoData("批量删除资源" + resourceType + "成功").Send(ctx)
}

package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/factory"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
)

func Rollback(ctx *gin.Context) {
	var req models.RollbackRequest
	if !utils.BindQuery(ctx, &req) {
		return
	}
	basicInfo := models.BasicInfo{}
	basicInfo.ClusterName = req.ClusterName

	clientset, err := models.NewClientSetForBasicInfo(&basicInfo)
	if err != nil {
		utils.Error(http.StatusInternalServerError, "连接集群失败："+err.Error()).Send(ctx)
		return
	}

	resourceType := ctx.Param("resource") // /api/:resource/rollback
	svc, ok := (&factory.ResourceFactory{}).GetService(resourceType, clientset).(interface {
		Rollback(req *models.RollbackRequest) error
	})
	if !ok {
		utils.Error(http.StatusInternalServerError, "不支持回滚操作的资源类型："+resourceType).Send(ctx)
		return
	}

	if err := svc.Rollback(&req); err != nil {
		utils.Error(http.StatusInternalServerError, "回滚"+resourceType+"资源失败："+err.Error()).Send(ctx)
		return
	}
	utils.SuccessNoData("回滚" + resourceType + "资源成功").Send(ctx)
}

// ListRevisionsHandler 查询 Deployment 的历史版本
func ListRevisionsHandler(ctx *gin.Context) {
	basicInfo := models.BasicInfo{}
	if !utils.BindQuery(ctx, &basicInfo) {
		return
	}

	clientset, err := models.NewClientSetForBasicInfo(&basicInfo)
	if err != nil {
		utils.Error(http.StatusInternalServerError, "连接集群失败："+err.Error()).Send(ctx)
		return
	}

	resourceType := ctx.Param("resource") // /api/:resource/list-revisions
	svc, ok := (&factory.ResourceFactory{}).GetService(resourceType, clientset).(interface {
		ListRevisions(basicInfo *models.BasicInfo) (interface{}, error)
	})

	if !ok {
		utils.Error(http.StatusInternalServerError, "不支持操作的资源类型："+resourceType).Send(ctx)
		return
	}
	revisions, err := svc.ListRevisions(&basicInfo)

	if err != nil {
		utils.Error(http.StatusInternalServerError, "查询"+resourceType+"资源历史版本失败："+err.Error()).Send(ctx)
		return
	}
	utils.SuccessWithItems("查询"+resourceType+"资源历史版本成功", revisions).Send(ctx)
}

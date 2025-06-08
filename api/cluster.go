package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/constants"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service"
	"github.com/maolchen/krm-backend/utils"
	"go.uber.org/zap"
	"net/http"
)

func ClusterAdd(ctx *gin.Context) {
	zap.S().Info("添加集群操作")
	var cluster models.ClusterInfo
	if err := ctx.ShouldBindJSON(&cluster); err != nil {
		utils.Error(http.StatusBadRequest, constants.RequestParmsError).Send(ctx)
		zap.S().Errorf("获取参数失败---->%s", err.Error())
		return
	}

	if err := service.AddCluster(&cluster); err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail).Send(ctx)
		zap.S().Errorf("数据插入失败----->%s", err.Error())
		return
	}

	utils.SuccessNoData(constants.ClusterSuccess).Send(ctx)
	zap.S().Info("添加集群成功")
}
func ClusterList(ctx *gin.Context) {
	zap.S().Info("列出集群操作")
	var clusterArray []models.ClusterListResponse
	clusterArray, err := service.ListCluster()
	if err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail+err.Error()).Send(ctx)
		zap.S().Errorf("数据查询失败----->%s", err.Error())
		return
	}

	utils.SuccessWithData(constants.ClusterSuccess, clusterArray).Send(ctx)
	zap.S().Info("查询集群成功")
}
func ClusterUpdate(ctx *gin.Context) {
	zap.S().Info("更新集群操作")
	name := ctx.Param("name")
	if name == "" {
		utils.Error(http.StatusBadRequest, "集群名称不能为空").Send(ctx)
		zap.S().Error("获取参数失败---->集群名称不能为空")
		return
	}

	var updates map[string]interface{}
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		utils.Error(http.StatusBadRequest, constants.RequestParmsError+err.Error()).Send(ctx)
		zap.S().Errorf("获取参数失败---->%s", err.Error())
		return
	}

	if err := service.UpdateClusterByName(name, updates); err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail+err.Error()).Send(ctx)
		zap.S().Errorf("数据查询失败----->%s", err.Error())
		return
	}
	utils.SuccessNoData(constants.ClusterSuccess).Send(ctx)
	zap.S().Info("更新集群成功")
}

func ClusterDelete(ctx *gin.Context) {
	zap.S().Info("删除集群操作")
	name := ctx.Param("name")
	if name == "" {
		utils.Error(http.StatusBadRequest, "集群名称不能为空").Send(ctx)
		zap.S().Error("获取参数失败---->集群名称不能为空")
		return
	}

	if err := service.DeleteCluster(name); err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail+err.Error()).Send(ctx)
		zap.S().Errorf("获取参数失败---->%s", err.Error())
		return
	}
	utils.SuccessNoData(constants.ClusterSuccess).Send(ctx)
	zap.S().Info("删除集群成功")
}

func ClusterEditByName(ctx *gin.Context) {
	zap.S().Info("编辑集群操作")
	name := ctx.Param("name")
	if name == "" {
		utils.Error(http.StatusBadRequest, "集群名称不能为空").Send(ctx)
		zap.S().Error("获取参数失败---->集群名称不能为空")
		return
	}

	clusterEditRes, err := service.GetClusterEditByName(name)
	if err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail+err.Error()).Send(ctx)
		zap.S().Errorf("获取参数失败---->%s", err.Error())
		return
	}
	utils.SuccessWithData(constants.ClusterSuccess, clusterEditRes).Send(ctx)
	zap.S().Infof("查询%s集群成功", name)
}

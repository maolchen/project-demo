package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/constants"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/clusters"
	"github.com/maolchen/krm-backend/utils"
	"go.uber.org/zap"
	"net/http"
)

func ClusterAdd(ctx *gin.Context) {
	zap.S().Info("正在添加集群......")
	var cluster models.ClusterInfo
	//ShouldBindJSON
	if !utils.BindJSON(ctx, &cluster) {
		return
	}

	if err := clusters.AddCluster(&cluster); err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail).Send(ctx)
		zap.S().Errorf("数据插入失败----->%s", err.Error())
		return
	}

	utils.SuccessNoData(constants.ClusterSuccess).Send(ctx)
	zap.S().Info("添加集群成功!!!")
}

func ClusterDelete(ctx *gin.Context) {
	zap.S().Info("正在删除集群......")
	var cluster models.ClusterInfo
	//ShouldBindQuery
	if !utils.BindQuery(ctx, &cluster) {
		return
	}

	if err := clusters.DeleteClusterByName(cluster.Name); err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail+err.Error()).Send(ctx)
		zap.S().Errorf("数据删除失败---->%s", err.Error())
		return
	}
	utils.SuccessNoData(constants.ClusterSuccess).Send(ctx)
	zap.S().Info("删除集群成功！！！")
}

func ClusterUpdate(ctx *gin.Context) {
	zap.S().Info("正在更新集群信息......")

	var cluster models.ClusterInfo

	//ShouldBindJSON
	if !utils.BindJSON(ctx, &cluster) {
		return
	}

	if cluster.Name == "" {
		utils.Error(http.StatusBadRequest, "集群名称不能为空").Send(ctx)
		zap.S().Error("获取参数失败---->集群名称不能为空")
		return
	}

	if err := clusters.UpdateClusterByName(&cluster); err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail+err.Error()).Send(ctx)
		zap.S().Errorf("数据更新失败----->%s", err.Error())
		return
	}
	utils.SuccessNoData(constants.ClusterSuccess).Send(ctx)
	zap.S().Info("更新集群成功！！！")
}

func ClusterList(ctx *gin.Context) {
	zap.S().Info("正在查询所有集群......")
	var clusterStatuses []models.ClusterStatus
	clusterStatuses, err := clusters.ListClusters()
	if err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail+err.Error()).Send(ctx)
		zap.S().Errorf("数据查询失败----->%s", err.Error())
		return
	}

	utils.SuccessWithItems(constants.ClusterSuccess, clusterStatuses).Send(ctx)
	zap.S().Info("查询所有集群成功！！！")
}

func ClusterGet(ctx *gin.Context) {
	zap.S().Info("正在获取集群详情......")
	var cluster models.ClusterInfo
	//前端只传一个name
	//ShouldBindQuery
	if !utils.BindQuery(ctx, &cluster) {
		return
	}

	clusterEditRes, err := clusters.GetClusterEditByName(cluster.Name)
	if err != nil {
		utils.Error(http.StatusInternalServerError, constants.ClusterFail+err.Error()).Send(ctx)
		zap.S().Errorf("获取集群详情失败---->%s", err.Error())
		return
	}
	utils.SuccessWithItem(constants.ClusterSuccess, clusterEditRes).Send(ctx)
	zap.S().Infof("获取集群%s详情成功！！！", cluster.Name)
}

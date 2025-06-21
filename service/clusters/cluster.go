package clusters

import (
	"encoding/base64"
	"errors"
	"github.com/maolchen/krm-backend/config"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/utils"
	"go.uber.org/zap"
)

func AddCluster(cluster *models.ClusterInfo) error {
	zap.S().Info("添加集群...")

	// base64解码kubeconfig
	zap.S().Debug("cluster.Kubeconfig::::", cluster.Kubeconfig)
	kubeconfig, err := base64.StdEncoding.DecodeString(cluster.Kubeconfig)
	if err != nil {
		zap.L().Error("kubeconfig 解码失败", zap.Error(err))
		return err
	}
	zap.S().Debug("kubeconfig:::::", string(kubeconfig))

	//校验kubeconfig

	if err := utils.ValidateKubeconfig(kubeconfig); err != nil {
		zap.L().Error("kubeconfig 校验失败", zap.Error(err))
		return err
	}

	//存储kubeconfig到内存中
	config.ClusterKubeconfig[cluster.Name] = kubeconfig

	zap.S().Debugf("当前存储的集群有：%s", utils.PrintClusterKubeconfig(config.ClusterKubeconfig))

	//数据库插入
	return cluster.Insert()
}

func DeleteCluster(cluster *models.ClusterInfo) error {
	zap.S().Info("删除集群...")

	if cluster.Name == "" {
		zap.S().Error("获取参数失败---->集群名称不能为空")
		return errors.New("集群名称不能为空")
	}
	//删除ClusterKubeconfig中的数据
	delete(config.ClusterKubeconfig, cluster.Name)
	zap.S().Debugf("当前存储的集群有：%s", utils.PrintClusterKubeconfig(config.ClusterKubeconfig))
	//删除数据库中的数据
	return cluster.Delete(cluster.Name)
}

func UpdateClusterByName(cluster *models.ClusterInfo) error {
	zap.S().Info("更新集群...")
	updates := map[string]interface{}{}
	// 如果提供了 Label，加入更新列表
	if cluster.Label != nil {
		updates["label"] = *cluster.Label
	}

	if cluster.Kubeconfig != "" {
		newKubeconfig, err := base64.StdEncoding.DecodeString(cluster.Kubeconfig)
		if err != nil {
			zap.L().Error("kubeconfig 解码失败", zap.Error(err))
			return err
		}
		//校验kubeconfig

		if err := utils.ValidateKubeconfig(newKubeconfig); err != nil {
			zap.L().Error("kubeconfig 校验失败", zap.Error(err))
			return err
		}
		//更新ClusterKubeconfig
		config.ClusterKubeconfig[cluster.Name] = newKubeconfig

		zap.S().Debugf("当前存储的集群有：%s", utils.PrintClusterKubeconfig(config.ClusterKubeconfig))

		updates["kubeconfig"] = cluster.Kubeconfig
	}

	//数据库更新操作
	return cluster.Update(cluster.Name, updates)
}

func ListCluster() (res []models.ClusterStatus, err error) {
	zap.S().Info("列出所有集群...")
	clusters, err := models.GetAllClusters()

	if err != nil {
		zap.L().Error("数据库查询失败", zap.Error(err))
		return res, err
	}

	for _, c := range clusters {
		clusterStatus, err := c.GetClusterStatus(config.ClusterKubeconfig[c.Name])
		if err != nil {
			zap.S().Errorf("获取集群列表失败%s", err.Error())
			return res, err
		}
		res = append(res, clusterStatus)
	}
	return res, nil
}

func GetClusterEditByName(cluster *models.ClusterInfo) (res models.ClusterStatus, err error) {
	zap.S().Info("获取集群详情...")
	if cluster.Name == "" {
		zap.S().Error("获取参数失败---->集群名称不能为空")
		return res, errors.New("集群名称不能为空")
	}
	clusterStatus, err := cluster.GetClusterStatus(config.ClusterKubeconfig[cluster.Name])
	if err != nil {
		zap.S().Errorf("获取集群详情失败%s", err.Error())
		return res, err
	}

	return clusterStatus, nil
}

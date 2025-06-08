package service

import (
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
)

func AddCluster(cluster *models.ClusterInfo) error {
	zap.S().Info("添加集群...")
	return cluster.Insert()
}

func ListCluster() ([]models.ClusterListResponse, error) {
	zap.S().Info("查看所有集群...")
	clusters, err := models.GetAllClusters()
	if err != nil {
		return nil, err
	}
	var res []models.ClusterListResponse
	for _, c := range clusters {
		status := map[string]string{"reachable": "false", "nodes": "0", "pods": "0"}
		//暂时不考虑k8s状态的逻辑
		//kubeconfigBytes := []byte(c.Kubeconfig)
		//clusterStatus, err := kubernetes.GetClusterStatus(kubeconfigBytes)
		//if err == nil {
		//	status = clusterStatus
		//}

		res = append(res, models.ClusterListResponse{
			Name:   c.Name,
			Label:  c.Label,
			Status: status,
		})
	}
	return res, nil
}

func UpdateClusterByName(name string, updates map[string]interface{}) error {
	zap.S().Info("更新集群...")
	clusterInfo := models.ClusterInfo{}
	return clusterInfo.Update(name, updates)
}

func DeleteCluster(name string) error {
	zap.S().Info("删除集群...")
	clusterInfo := models.ClusterInfo{}
	return clusterInfo.Delete(name)
}

func GetClusterEditByName(name string) (models.ClusterEditResponse, error) {
	zap.S().Info("编辑集群...")
	return (&models.ClusterInfo{}).GetByName(name)
}

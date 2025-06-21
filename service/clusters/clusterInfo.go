package clusters

import (
	"fmt"
	"github.com/maolchen/krm-backend/config"
	. "github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service/common"
	"go.uber.org/zap"
)

// 获取集群状态
func GetClusterStatuses(clusters []ClusterInfo) ([]ClusterStatus, error) {
	var res []ClusterStatus
	for _, c := range clusters {
		status, err := GetClusterStatusByName(c.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, status)
	}
	return res, nil
}

func GetClusterStatusByName(name string) (ClusterStatus, error) {
	var cluster ClusterInfo
	resp, err := cluster.GetByName(name)
	if err != nil {
		return ClusterStatus{}, err
	}

	status := ClusterStatus{
		ClusterResponse: resp,
		Version:         "",
		Status:          "inactive",
	}

	clientset, err := common.NewClientSet(config.ClusterKubeconfig[name])
	if err != nil {
		return status, err
	}

	serverVersion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return status, fmt.Errorf("无法访问集群，请检查权限或网络: %v", err)
	}

	status.Version = serverVersion.String()
	status.Status = "active"
	zap.S().Infof("Successfully connected to clusters,当前连接集群：%s------>,当前集群状态：%s------>当前集群版本：%s", name, status.Status, serverVersion.String())
	return status, nil
}

package models

import (
	"github.com/maolchen/krm-backend/config"
	"github.com/maolchen/krm-backend/service/common"
	"k8s.io/client-go/kubernetes"
)

// 定义一个全局的数据结构
type BasicInfo struct {
	ClusterName string      `json:"clusterName" form:"clusterName"`
	Namespace   string      `json:"namespace" form:"namespace"`
	Name        string      `json:"name" form:"name"`
	Item        interface{} `json:"item"`
}

// BatchRequest 定义批量操作结构体
type BatchRequest struct {
	BasicInfo BasicInfo `json:"basicInfo"`
	Names     []string  `json:"names"` // 要删除的资源名称列表
}

// RollbackRequest 定义回滚请求的结构体
type RollbackRequest struct {
	ClusterName string `json:"clusterName" form:"clusterName"`
	Namespace   string `json:"namespace" form:"namespace"`
	Name        string `json:"name" form:"name"`
	Revision    string `json:"revision" form:"revision"`
}

func NewClientSetForBasicInfo(b *BasicInfo) (*kubernetes.Clientset, error) {

	kubeconfig := config.ClusterKubeconfig[b.ClusterName]
	clientSet, err := common.NewClientSet(kubeconfig)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

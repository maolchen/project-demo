package models

import (
	"errors"
	"fmt"
	"github.com/maolchen/krm-backend/database"
	"github.com/maolchen/krm-backend/service/common"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ClusterInfo struct {
	gorm.Model
	Name       string  `gorm:"unique;not null" json:"name" form:"name"`
	Label      *string `gorm:"default:null" json:"label" form:"label"` //集群标签
	Kubeconfig string  `gorm:"not null" json:"kubeconfig" form:"kubeconfig"`
}

type ClusterResponse struct {
	Name  string  `json:"name"`
	Label *string `json:"label"`
}

type ClusterStatus struct {
	ClusterResponse
	Version string `json:"version"`
	Status  string `json:"status"`
}

func (c *ClusterInfo) TableName() string {
	return "cluster_info"
}

// 添加集群
func (c *ClusterInfo) Insert() error {
	return database.GetDB().Create(c).Error
}

// 更新集群
func (c *ClusterInfo) Update(name string, updates map[string]interface{}) error {
	//delete(updates, "name") // 禁止修改 Name
	return database.GetDB().
		Model(c).
		Where("name = ?", name).
		Updates(updates).Error
}

// 删除集群
func (c *ClusterInfo) Delete(name string) error {
	db := database.GetDB().Unscoped().Where("name = ?", name).Delete(&ClusterInfo{})
	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("集群不存在或已被删除")
	}

	return nil
}

// 获取集群byName
func (c *ClusterInfo) GetByName(name string) (ClusterResponse, error) {

	db := database.GetDB()
	if err := db.Where("name = ?", name).First(&c).Error; err != nil {
		return ClusterResponse{}, err
	}

	return ClusterResponse{
		Name:  c.Name,
		Label: c.Label,
	}, nil
}

// 获取所有集群
func GetAllClusters() ([]ClusterInfo, error) {
	var clusters []ClusterInfo
	db := database.GetDB()
	if err := db.Find(&clusters).Error; err != nil {
		return nil, err
	}
	return clusters, nil
}

// 获取集群状态
func (c *ClusterInfo) GetClusterStatus(kubconfig []byte) (ClusterStatus, error) {
	clusterStatus := ClusterStatus{
		ClusterResponse: ClusterResponse{
			Name:  c.Name,
			Label: c.Label,
		},
		Status: "inactive",
	}

	clientset, err := common.NewClientSet(kubconfig)
	if err != nil {
		return clusterStatus, err
	}
	// 获取集群版本
	serverVersion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return clusterStatus, fmt.Errorf("无法访问集群，请检查权限或网络: %v", err)
	}

	//
	zap.S().Infof("Successfully connected to clusters,当前集群版本%s", serverVersion.String())
	clusterStatus.Version = serverVersion.String()
	clusterStatus.Status = "active"

	zap.S().Debugf("当前集群状态：%s", clusterStatus)
	return clusterStatus, nil
}

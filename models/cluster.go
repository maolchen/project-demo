package models

import (
	"github.com/maolchen/krm-backend/database"
	"gorm.io/gorm"
)

type ClusterInfo struct {
	gorm.Model
	Name       string  `gorm:"unique;not null" json:"name"`
	Label      *string `gorm:"default:null" json:"label"` //集群标签
	Kubeconfig string  `gorm:"not null" json:"kubeconfig"`
}

type ClusterEditResponse struct {
	Name  string  `json:"name"`
	Label *string `json:"label"`
}

type ClusterListResponse struct {
	Name   string            `json:"name"`
	Label  *string           `json:"label"`
	Status map[string]string `json:"status"`
}

func (c *ClusterInfo) TableName() string {
	return "cluster_info"
}

// 添加集群
func (c *ClusterInfo) Insert() error {
	return database.GetDB().Create(c).Error
}

func (c *ClusterInfo) Update(name string, updates map[string]interface{}) error {
	delete(updates, "name") // 禁止修改 Name
	return database.GetDB().
		Model(c).
		Where("name = ?", name).
		Updates(updates).Error
}

func (c *ClusterInfo) Delete(name string) error {
	return database.GetDB().
		Where("name = ?", name).
		Delete(&ClusterInfo{}).Error
}

func (c *ClusterInfo) GetByName(name string) (ClusterEditResponse, error) {
	var cluster ClusterInfo
	db := database.GetDB()
	if err := db.Where("name = ?", name).First(&cluster).Error; err != nil {
		return ClusterEditResponse{}, err
	}

	return ClusterEditResponse{
		Name:  cluster.Name,
		Label: cluster.Label,
	}, nil
}

func GetAllClusters() ([]ClusterInfo, error) {
	var clusters []ClusterInfo
	db := database.GetDB()
	if err := db.Find(&clusters).Error; err != nil {
		return nil, err
	}
	return clusters, nil
}

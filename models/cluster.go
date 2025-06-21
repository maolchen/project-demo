package models

import (
	"github.com/maolchen/krm-backend/database"
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

// 删除集群
func (c *ClusterInfo) Delete(name string) error {
	return DeleteByName(c, name)
}

// 更新集群
func (c *ClusterInfo) Update(name string, updates map[string]interface{}) error {
	return UpdateByName(c, name, updates)
}

//func (c *ClusterInfo) Update(name string, updates map[string]interface{}) error {
//	//delete(updates, "name") // 禁止修改 Name
//	return database.GetDB().
//		Model(c).
//		Where("name = ?", name).
//		Updates(updates).Error
//}

//func (c *ClusterInfo) Delete(name string) error {
//	db := database.GetDB().Unscoped().Where("name = ?", name).Delete(&ClusterInfo{})
//	if db.Error != nil {
//		return db.Error
//	}
//
//	if db.RowsAffected == 0 {
//		return errors.New("集群不存在或已被删除")
//	}
//
//	return nil
//}

// 获取所有集群
func GetAllClusters() ([]ClusterInfo, error) {
	var clusters []ClusterInfo
	db := database.GetDB()
	if err := db.Find(&clusters).Error; err != nil {
		return nil, err
	}
	return clusters, nil
}

// 查询单个集群信息
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

package models

import (
	"github.com/maolchen/krm-backend/database"
	"gorm.io/gorm"
)

type ClusterInfo struct {
	gorm.Model
	Name       string  `gorm:"unique;not null" json:"name"`
	Label      *string `gorm:"default:null" json:"label"` //集群标签
	Kubeconfig string  `gorm:"not null" json:"-`
}

type ClusterRespoense struct {
	Name   string            `json:"name"`
	Label  *string           `json:"label"`
	Status map[string]string `json:"status"`
}

var db = database.GetDB()

func (c *ClusterInfo) TableName() string {
	return "cluster_info"
}

// 添加集群
func (c *ClusterInfo) Insert() error {

	res := db.Create(&c)
	return res.Error

}

func (c *ClusterInfo) Update(name string, updates map[string]interface{}) error {
	var cluster ClusterInfo
	if _, err := c.GetByName(name); err != nil {
		return err
	}
	if newKubeconfig, ok := updates["kubeconfig"]; ok {
		cluster.Kubeconfig = newKubeconfig.(string)
	}
	if newLabel, ok := updates["label"]; ok {
		labelStr := newLabel.(string)
		cluster.Label = &labelStr
	}
	return db.Save(&cluster).Error

}

func (c *ClusterInfo) Delete(name string) error {
	var cluster ClusterInfo
	if _, err := c.GetByName(name); err != nil {
		return err
	}
	return db.Delete(&cluster).Error
}

func (c *ClusterInfo) GetAll() []ClusterInfo {
	var ClusterArray []ClusterInfo
	db.Find(&ClusterArray)
	return ClusterArray
}

func (c *ClusterInfo) GetByName(name string) (ClusterInfo, error) {
	var cluster ClusterInfo
	err := db.Where("name = ?", name).First(&cluster).Error
	return cluster, err
}

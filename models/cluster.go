package models

import "gorm.io/gorm"

type ClusterInfo struct {
	gorm.Model
	Name       string `gorm:"unique;not null" json:"name"`
	Label      string `gorm:"default null" json:"label"` //集群标签
	Kubeconfig string `gorm:"not null" json:"kubeconfig"`
}

type ClusterRespoense struct {
	Name   string            `json:"name"`
	Label  string            `json:"label"`
	status map[string]string `json:"status"`
}

func (c *ClusterInfo) TableName() string {
	return "cluster_info"
}

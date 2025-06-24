package service

import (
	"github.com/maolchen/krm-backend/models"
)

// 核心接口定义
type ResourceService interface {
	Create(basicInfo *models.BasicInfo) error
	Delete(basicInfo *models.BasicInfo) error
	Update(basicInfo *models.BasicInfo) error
	Get(basicInfo *models.BasicInfo) (interface{}, error)
	List(basicInfo *models.BasicInfo) (interface{}, error)
	Restart(basicInfo *models.BasicInfo) error // 添加重启接口
	Rollback(req *models.RollbackRequest) error
	ListRevisions(basicInfo *models.BasicInfo) (interface{}, error) // 添加查询历史版本接口
}

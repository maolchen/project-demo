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
}

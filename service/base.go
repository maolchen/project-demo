package service

import (
	"errors"
	"github.com/maolchen/krm-backend/models"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

// 定义一个BaseResource结构体，字段为一个Clientset
type BaseResource struct {
	Clientset kubernetes.Interface
}

// NewBaseClientset
func NewBaseResource(clientset kubernetes.Interface) *BaseResource {
	return &BaseResource{Clientset: clientset}
}

// InitObjectFromItem 将 接收的json中的 item 解析为 unstructured 对象并转换为目标结构体
func (b *BaseResource) InitObjectFromItem(basicInfo *models.BasicInfo, obj interface{}) error {
	unstructuredMap, ok := basicInfo.Item.(map[string]interface{})
	if !ok {
		return errors.New("item 不是 map[string]interface{}")
	}
	//runtime.DefaultUnstructuredConverter.FromUnstructured 官方推荐的转换方法
	return runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredMap, obj)
}

// GetName 获取资源名称
func (b *BaseResource) GetName(basicInfo *models.BasicInfo) string {
	return basicInfo.Name
}

// GetNamespace 获取命名空间
func (b *BaseResource) GetNamespace(basicInfo *models.BasicInfo) string {
	return basicInfo.Namespace
}

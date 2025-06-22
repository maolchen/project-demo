package resources

import (
	"context"
	"errors"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// NamespaceService 是 namespace 资源的服务实现
type NamespaceService struct {
	service.BaseResource // 基础服务封装了通用逻辑
}

// NewNamespaceService 创建一个 NamespaceService 实例
func NewNamespaceService(clientset kubernetes.Interface) service.ResourceService {
	return &NamespaceService{
		BaseResource: *service.NewBaseResource(clientset),
	}
}

// Create 创建一个 Kubernetes Namespace
func (n *NamespaceService) Create(basicInfo *models.BasicInfo) error {
	var ns corev1.Namespace
	if err := n.InitObjectFromItem(basicInfo, &ns); err != nil {
		return err
	}
	ns.Name = basicInfo.Name // 设置名称

	_, err := n.Clientset.CoreV1().Namespaces().Create(context.TODO(), &ns, metav1.CreateOptions{})
	if err != nil {
		zap.S().Errorf("创建命名空间失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("创建命名空间 %s 成功!!!", basicInfo.Name)
	return nil
}

// Delete 删除指定名称的 Namespace
func (n *NamespaceService) Delete(basicInfo *models.BasicInfo) error {
	name := basicInfo.Name
	if name == "kube-system" {
		zap.S().Errorf("删除命名空间失败：不允许删除 kube-system")
		return errors.New("不允许删除 kube-system")
	}
	err := n.Clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		zap.S().Errorf("删除命名空间失败：%s----->%s", name, err.Error())
		return err
	}
	zap.S().Infof("删除命名空间 %s 成功!!!", name)
	return nil
}

// Update 更新一个已有的 Namespace（注意：metadata.name 等字段不能修改）
func (n *NamespaceService) Update(basicInfo *models.BasicInfo) error {
	var ns corev1.Namespace
	if err := n.InitObjectFromItem(basicInfo, &ns); err != nil {
		return err
	}
	ns.Name = basicInfo.Name

	_, err := n.Clientset.CoreV1().Namespaces().Update(context.TODO(), &ns, metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("更新命名空间失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("更新命名空间 %s 成功!!!", basicInfo.Name)
	return nil
}

// Get 获取指定名称的 Namespace 的详细信息
func (n *NamespaceService) Get(basicInfo *models.BasicInfo) (interface{}, error) {
	name := basicInfo.Name
	ns, err := n.Clientset.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		zap.S().Errorf("获取命名空间详情失败：%s----->%s", name, err.Error())
		return nil, err
	}
	zap.S().Infof("获取命名空间 %s 成功!!!", name)
	return ns, nil
}

// List 列出所有 Namespace
func (n *NamespaceService) List(basicInfo *models.BasicInfo) (interface{}, error) {
	list, err := n.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.S().Errorf("列出命名空间失败：%s----->%s", basicInfo.Name, err.Error())
		return nil, err
	}
	zap.S().Infof("列出命名空间成功！共 %d 个", len(list.Items))
	return list.Items, nil
}

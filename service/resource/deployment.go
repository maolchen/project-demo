package resources

import (
	"context"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeploymentService 是 Deployment 资源的服务实现
type DeploymentService struct {
	service.BaseResource
}

// NewDeploymentService 创建一个 DeploymentService 实例
func NewDeploymentService(clientset kubernetes.Interface) service.ResourceService {
	return &DeploymentService{BaseResource: *service.NewBaseResource(clientset)}
}

// Create 创建 Deployment
func (d *DeploymentService) Create(basicInfo *models.BasicInfo) error {
	var deploy appsv1.Deployment
	if err := d.InitObjectFromItem(basicInfo, &deploy); err != nil {
		return err
	}
	deploy.Name = basicInfo.Name
	deploy.Namespace = basicInfo.Namespace

	_, err := d.Clientset.AppsV1().Deployments(deploy.Namespace).Create(context.TODO(), &deploy, metav1.CreateOptions{})
	if err != nil {
		zap.S().Errorf("创建 Deployment 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("创建 Deployment %s 成功!!!", basicInfo.Name)
	return nil
}

// Delete 删除 Deployment
func (d *DeploymentService) Delete(basicInfo *models.BasicInfo) error {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	err := d.Clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		zap.S().Errorf("删除 Deployment 失败：%s----->%s", name, err.Error())
		return err
	}
	zap.S().Infof("删除 Deployment %s 成功!!!", name)
	return nil
}

// Update 更新 Deployment
func (d *DeploymentService) Update(basicInfo *models.BasicInfo) error {
	var deploy appsv1.Deployment
	if err := d.InitObjectFromItem(basicInfo, &deploy); err != nil {
		return err
	}
	deploy.Name = basicInfo.Name
	deploy.Namespace = basicInfo.Namespace

	_, err := d.Clientset.AppsV1().Deployments(deploy.Namespace).Update(context.TODO(), &deploy, metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("更新 Deployment 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("更新 Deployment %s 成功!!!", basicInfo.Name)
	return nil
}

// Get 获取 Deployment 详情
func (d *DeploymentService) Get(basicInfo *models.BasicInfo) (interface{}, error) {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	deploy, err := d.Clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		zap.S().Errorf("获取 Deployment 详情失败：%s----->%s", name, err.Error())
		return nil, err
	}
	zap.S().Infof("获取 Deployment %s 成功!!!", name)
	return deploy, nil
}

// List 列出所有 Deployment
func (d *DeploymentService) List(basicInfo *models.BasicInfo) (interface{}, error) {
	namespace := basicInfo.Namespace
	list, err := d.Clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.S().Errorf("列出 Deployment 失败：%s", err.Error())
		return nil, err
	}
	zap.S().Infof("列出 Deployment 成功！共 %d 个", len(list.Items))
	return list.Items, nil
}

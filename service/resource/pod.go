package resources

import (
	"context"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// PodService 是 Pod 资源的服务实现
type PodService struct {
	service.BaseResource
}

// NewPodService 创建一个 PodService 实例
func NewPodService(clientset kubernetes.Interface) service.ResourceService {
	return &PodService{BaseResource: *service.NewBaseResource(clientset)}
}

// Create 在指定 Namespace 中创建 Pod
func (p *PodService) Create(basicInfo *models.BasicInfo) error {
	var pod corev1.Pod
	if err := p.InitObjectFromItem(basicInfo, &pod); err != nil {
		return err
	}
	pod.Name = basicInfo.Name
	pod.Namespace = basicInfo.Namespace

	_, err := p.Clientset.CoreV1().Pods(pod.Namespace).Create(context.TODO(), &pod, metav1.CreateOptions{})
	if err != nil {
		zap.S().Errorf("创建 Pod 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("创建 Pod %s 成功!!!", basicInfo.Name)
	return nil
}

// Delete 删除指定 Namespace 下的 Pod
func (p *PodService) Delete(basicInfo *models.BasicInfo) error {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	err := p.Clientset.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		zap.S().Errorf("删除 Pod 失败：%s----->%s", name, err.Error())
		return err
	}
	zap.S().Infof("删除 Pod %s 成功!!!", name)
	return nil
}

// Update 更新 Pod（注意：metadata.name 不能更改）
func (p *PodService) Update(basicInfo *models.BasicInfo) error {
	var pod corev1.Pod
	if err := p.InitObjectFromItem(basicInfo, &pod); err != nil {
		return err
	}
	pod.Name = basicInfo.Name
	pod.Namespace = basicInfo.Namespace

	_, err := p.Clientset.CoreV1().Pods(pod.Namespace).Update(context.TODO(), &pod, metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("更新 Pod 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("更新 Pod %s 成功!!!", basicInfo.Name)
	return nil
}

// Get 获取指定 Namespace 下的 Pod 详情
func (p *PodService) Get(basicInfo *models.BasicInfo) (interface{}, error) {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	pod, err := p.Clientset.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		zap.S().Errorf("获取 Pod 详情失败：%s----->%s", name, err.Error())
		return nil, err
	}
	zap.S().Infof("获取 Pod %s 成功!!!", name)
	return pod, nil
}

// List 列出指定 Namespace 下的所有 Pod
func (p *PodService) List(basicInfo *models.BasicInfo) (interface{}, error) {
	namespace := basicInfo.Namespace
	list, err := p.Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.S().Errorf("列出 Pod 失败：%s", err.Error())
		return nil, err
	}
	zap.S().Infof("列出 Pod 成功！共 %d 个", len(list.Items))
	return list.Items, nil
}

// Restart 重启指定 Namespace 下的 Pod
func (p *PodService) Restart(basicInfo *models.BasicInfo) error {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	err := p.Clientset.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		zap.S().Errorf("重启 Pod 失败：%s----->%s", name, err.Error())
		return err
	}
	zap.S().Infof("重启 Pod %s 成功!!!", name)
	return nil
}

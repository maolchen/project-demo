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

// StatefulSetService 是 StatefulSet 资源的服务实现
type StatefulSetService struct {
	service.BaseResource
}

// NewStatefulSetService 创建一个 StatefulSetService 实例
func NewStatefulSetService(clientset kubernetes.Interface) service.ResourceService {
	return &StatefulSetService{BaseResource: *service.NewBaseResource(clientset)}
}

// Create 创建 StatefulSet
func (s *StatefulSetService) Create(basicInfo *models.BasicInfo) error {
	var sts appsv1.StatefulSet
	if err := s.InitObjectFromItem(basicInfo, &sts); err != nil {
		return err
	}
	sts.Name = basicInfo.Name
	sts.Namespace = basicInfo.Namespace

	_, err := s.Clientset.AppsV1().StatefulSets(sts.Namespace).Create(context.TODO(), &sts, metav1.CreateOptions{})
	if err != nil {
		zap.S().Errorf("创建 StatefulSet 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("创建 StatefulSet %s 成功!!!", basicInfo.Name)
	return nil
}

// Delete 删除 StatefulSet
func (s *StatefulSetService) Delete(basicInfo *models.BasicInfo) error {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	err := s.Clientset.AppsV1().StatefulSets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		zap.S().Errorf("删除 StatefulSet 失败：%s----->%s", name, err.Error())
		return err
	}
	zap.S().Infof("删除 StatefulSet %s 成功!!!", name)
	return nil
}

// Update 更新 StatefulSet
func (s *StatefulSetService) Update(basicInfo *models.BasicInfo) error {
	var sts appsv1.StatefulSet
	if err := s.InitObjectFromItem(basicInfo, &sts); err != nil {
		return err
	}
	sts.Name = basicInfo.Name
	sts.Namespace = basicInfo.Namespace

	_, err := s.Clientset.AppsV1().StatefulSets(sts.Namespace).Update(context.TODO(), &sts, metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("更新 StatefulSet 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("更新 StatefulSet %s 成功!!!", basicInfo.Name)
	return nil
}

// Get 获取 StatefulSet 详情
func (s *StatefulSetService) Get(basicInfo *models.BasicInfo) (interface{}, error) {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	sts, err := s.Clientset.AppsV1().StatefulSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		zap.S().Errorf("获取 StatefulSet 详情失败：%s----->%s", name, err.Error())
		return nil, err
	}
	zap.S().Infof("获取 StatefulSet %s 成功!!!", name)
	return sts, nil
}

// List 列出所有 StatefulSet
func (s *StatefulSetService) List(basicInfo *models.BasicInfo) (interface{}, error) {
	namespace := basicInfo.Namespace
	list, err := s.Clientset.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.S().Errorf("列出 StatefulSet 失败：%s", err.Error())
		return nil, err
	}
	zap.S().Infof("列出 StatefulSet 成功！共 %d 个", len(list.Items))
	return list.Items, nil
}

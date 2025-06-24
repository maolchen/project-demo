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

// DaemonSetService 是 DaemonSet 资源的服务实现
type DaemonSetService struct {
	service.BaseResource
}

// NewDaemonSetService 创建一个 DaemonSetService 实例
func NewDaemonSetService(clientset kubernetes.Interface) service.ResourceService {
	return &DaemonSetService{BaseResource: *service.NewBaseResource(clientset)}
}

// Create 创建 DaemonSet
func (d *DaemonSetService) Create(basicInfo *models.BasicInfo) error {
	var ds appsv1.DaemonSet
	if err := d.InitObjectFromItem(basicInfo, &ds); err != nil {
		return err
	}
	ds.Name = basicInfo.Name
	ds.Namespace = basicInfo.Namespace

	_, err := d.Clientset.AppsV1().DaemonSets(ds.Namespace).Create(context.TODO(), &ds, metav1.CreateOptions{})
	if err != nil {
		zap.S().Errorf("创建 DaemonSet 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("创建 DaemonSet %s 成功!!!", basicInfo.Name)
	return nil
}

// Delete 删除 DaemonSet
func (d *DaemonSetService) Delete(basicInfo *models.BasicInfo) error {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	err := d.Clientset.AppsV1().DaemonSets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		zap.S().Errorf("删除 DaemonSet 失败：%s----->%s", name, err.Error())
		return err
	}
	zap.S().Infof("删除 DaemonSet %s 成功!!!", name)
	return nil
}

// Update 更新 DaemonSet
func (d *DaemonSetService) Update(basicInfo *models.BasicInfo) error {
	var ds appsv1.DaemonSet
	if err := d.InitObjectFromItem(basicInfo, &ds); err != nil {
		return err
	}
	ds.Name = basicInfo.Name
	ds.Namespace = basicInfo.Namespace

	_, err := d.Clientset.AppsV1().DaemonSets(ds.Namespace).Update(context.TODO(), &ds, metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("更新 DaemonSet 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("更新 DaemonSet %s 成功!!!", basicInfo.Name)
	return nil
}

// Get 获取 DaemonSet 详情
func (d *DaemonSetService) Get(basicInfo *models.BasicInfo) (interface{}, error) {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	ds, err := d.Clientset.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		zap.S().Errorf("获取 DaemonSet 详情失败：%s----->%s", name, err.Error())
		return nil, err
	}
	zap.S().Infof("获取 DaemonSet %s 成功!!!", name)
	return ds, nil
}

// List 列出所有 DaemonSet
func (d *DaemonSetService) List(basicInfo *models.BasicInfo) (interface{}, error) {
	namespace := basicInfo.Namespace
	list, err := d.Clientset.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.S().Errorf("列出 DaemonSet 失败：%s", err.Error())
		return nil, err
	}
	zap.S().Infof("列出 DaemonSet 成功！共 %d 个", len(list.Items))
	return list.Items, nil
}

// Restart 重启指定 Namespace 下的 DaemonSet
func (d *DaemonSetService) Restart(basicInfo *models.BasicInfo) error {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	daemonSet, err := d.Clientset.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		zap.S().Errorf("获取 DaemonSet 详情失败：%s----->%s", name, err.Error())
		return err
	}

	// 增加版本号以触发滚动更新
	if daemonSet.Spec.Template.Annotations == nil {
		daemonSet.Spec.Template.Annotations = make(map[string]string)
	}
	currentVersion := daemonSet.Spec.Template.Annotations["restart-version"]
	newVersion := currentVersion
	if newVersion == "" {
		newVersion = "1"
	} else {
		newVersion = string(currentVersion[0] + 1)
	}
	daemonSet.Spec.Template.Annotations["restart-version"] = newVersion

	_, err = d.Clientset.AppsV1().DaemonSets(namespace).Update(context.TODO(), daemonSet, metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("重启 DaemonSet 失败：%s----->%s", name, err.Error())
		return err
	}
	zap.S().Infof("重启 DaemonSet %s 成功!!!", name)
	return nil
}

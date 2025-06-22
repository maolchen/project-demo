package resources

import (
	"context"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service"
	"go.uber.org/zap"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CronJobService 是 CronJob 资源的服务实现
type CronJobService struct {
	service.BaseResource
}

// NewCronJobService 创建一个 CronJobService 实例
func NewCronJobService(clientset kubernetes.Interface) service.ResourceService {
	return &CronJobService{BaseResource: *service.NewBaseResource(clientset)}
}

// Create 创建 CronJob
func (c *CronJobService) Create(basicInfo *models.BasicInfo) error {
	var job batchv1beta1.CronJob
	if err := c.InitObjectFromItem(basicInfo, &job); err != nil {
		return err
	}
	job.Name = basicInfo.Name
	job.Namespace = basicInfo.Namespace

	_, err := c.Clientset.BatchV1beta1().CronJobs(job.Namespace).Create(context.TODO(), &job, metav1.CreateOptions{})
	if err != nil {
		zap.S().Errorf("创建 CronJob 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("创建 CronJob %s 成功!!!", basicInfo.Name)
	return nil
}

// Delete 删除 CronJob
func (c *CronJobService) Delete(basicInfo *models.BasicInfo) error {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	err := c.Clientset.BatchV1beta1().CronJobs(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		zap.S().Errorf("删除 CronJob 失败：%s----->%s", name, err.Error())
		return err
	}
	zap.S().Infof("删除 CronJob %s 成功!!!", name)
	return nil
}

// Update 更新 CronJob
func (c *CronJobService) Update(basicInfo *models.BasicInfo) error {
	var job batchv1beta1.CronJob
	if err := c.InitObjectFromItem(basicInfo, &job); err != nil {
		return err
	}
	job.Name = basicInfo.Name
	job.Namespace = basicInfo.Namespace

	_, err := c.Clientset.BatchV1beta1().CronJobs(job.Namespace).Update(context.TODO(), &job, metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("更新 CronJob 失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("更新 CronJob %s 成功!!!", basicInfo.Name)
	return nil
}

// Get 获取 CronJob 详情
func (c *CronJobService) Get(basicInfo *models.BasicInfo) (interface{}, error) {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	job, err := c.Clientset.BatchV1beta1().CronJobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		zap.S().Errorf("获取 CronJob 详情失败：%s----->%s", name, err.Error())
		return nil, err
	}
	zap.S().Infof("获取 CronJob %s 成功!!!", name)
	return job, nil
}

// List 列出所有 CronJob
func (c *CronJobService) List(basicInfo *models.BasicInfo) (interface{}, error) {
	namespace := basicInfo.Namespace
	list, err := c.Clientset.BatchV1beta1().CronJobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.S().Errorf("列出 CronJob 失败：%s", err.Error())
		return nil, err
	}
	zap.S().Infof("列出 CronJob 成功！共 %d 个", len(list.Items))
	return list.Items, nil
}

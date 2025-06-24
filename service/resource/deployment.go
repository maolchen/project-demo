package resources

import (
	"context"
	"fmt"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/service"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
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

// Restart 重启指定 Namespace 下的 Deployment
func (d *DeploymentService) Restart(basicInfo *models.BasicInfo) error {
	name := basicInfo.Name
	namespace := basicInfo.Namespace
	deployment, err := d.Clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		zap.S().Errorf("获取 Deployment 详情失败：%s----->%s", name, err.Error())
		return err
	}

	// 增加版本号以触发滚动更新
	if deployment.Spec.Template.Annotations == nil {
		deployment.Spec.Template.Annotations = make(map[string]string)
	}
	currentVersion := deployment.Spec.Template.Annotations["restart-version"]
	newVersion := currentVersion
	if newVersion == "" {
		newVersion = "1"
	} else {
		newVersion = string(currentVersion[0] + 1)
	}
	deployment.Spec.Template.Annotations["restart-version"] = newVersion

	_, err = d.Clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("重启 Deployment 失败：%s----->%s", name, err.Error())
		return err
	}
	zap.S().Infof("重启 Deployment %s 成功!!!", name)
	return nil
}

// Rollback 回滚指定 Deployment 到指定的修订版本
func (d *DeploymentService) Rollback(req *models.RollbackRequest) error {
	namespace := req.Namespace
	deploymentName := req.Name
	replicaSetName := req.Revision // 实际上是 ReplicaSet 的 Name

	fmt.Println("replicaSetName:", replicaSetName)
	// Step 1: 获取目标 ReplicaSet
	rs, err := d.Clientset.AppsV1().ReplicaSets(namespace).Get(context.TODO(), replicaSetName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("获取 ReplicaSet [%s] 失败: %w", replicaSetName, err)
	}

	// Step 2: 获取当前 Deployment
	deployment, err := d.Clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("获取 Deployment [%s] 失败: %w", deploymentName, err)
	}

	// Step 3: 替换 Deployment 的 PodTemplateSpec 为目标 ReplicaSet 的模板
	updatedDeployment := deployment.DeepCopy()
	updatedDeployment.Spec.Template = rs.Spec.Template

	// Step 4: 更新 Deployment（触发滚动更新）
	_, err = d.Clientset.AppsV1().Deployments(namespace).Update(context.TODO(), updatedDeployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("更新 Deployment [%s] 失败: %w", deploymentName, err)
	}

	return nil
}

// ListRevisions 查询 Deployment 的历史版本
func (d *DeploymentService) ListRevisions(basicInfo *models.BasicInfo) (interface{}, error) {
	namespace := basicInfo.Namespace
	name := basicInfo.Name
	//// 获取 Deployment 对象
	deployment, err := d.Clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取 Deployment 失败: %w", err)
	}

	// Step 2: 构建 LabelSelector 字符串
	labelSelector := labels.Set(deployment.Spec.Selector.MatchLabels).String()

	replicaSets, err := d.Clientset.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})

	if err != nil {
		return nil, fmt.Errorf("查询 ControllerRevisions 失败: %w", err)
	}

	// 筛选出与 Deployment 相关的 ControllerRevision
	//var revisions []*appsv1.ControllerRevision
	//revisions = controllerRevisions.Items
	//for _, cr := range controllerRevisions.Items {
	//	for _, ownerRef := range cr.OwnerReferences {
	//		if ownerRef.Kind == "Deployment" && ownerRef.Name == name {
	//			revisions = append(revisions, &cr)
	//			break
	//		}
	//	}
	//}

	return replicaSets.Items, nil
}

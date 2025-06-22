package namespace

import (
	"context"
	"errors"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func UpdateNamespace(basicInfo *models.BasicInfo) error {
	clientset, err := models.NewClientSetForBasicInfo(basicInfo)
	if err != nil {
		return err
	}
	// 将 basicInfo.Item 转换为 *corev1.Namespace
	ns, ok := basicInfo.Item.(*corev1.Namespace)
	if !ok {
		// 如果不是期望类型，则尝试从 map 转换
		unstructuredMap, ok := basicInfo.Item.(map[string]interface{})
		if !ok {
			return errors.New("item 不是 map 或 *corev1.Namespace")
		}

		// 使用官方 unstructured converter 转换

		ns = &corev1.Namespace{}
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredMap, ns)
		if err != nil {
			zap.S().Errorf("无法将 item 转换为 Namespace：%s", err.Error())
			return err
		}
	}

	// 设置名字确保更新的是正确的对象
	ns.Name = basicInfo.Name

	_, err = clientset.CoreV1().Namespaces().Update(context.TODO(), ns, metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("更新命名空间失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}

	zap.S().Infof("更新命名空间%s成功!!!", basicInfo.Name)
	return nil

}

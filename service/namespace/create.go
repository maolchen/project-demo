package namespace

import (
	"context"
	"errors"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	corev1 "k8s.io/api/core/v1"
)

func CreateNamespace(basicInfo *models.BasicInfo) error {
	var ns *corev1.Namespace
	clientset, err := models.NewClientSetForBasicInfo(basicInfo)
	if err != nil {
		return err
	}
	// Step 1: 解析 Item
	if basicInfo.Item != nil {
		if tmp, ok := basicInfo.Item.(*corev1.Namespace); ok {
			ns = tmp
		} else if m, ok := basicInfo.Item.(map[string]interface{}); ok {
			ns = &corev1.Namespace{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(m, ns); err != nil {
				zap.S().Errorf("无法将 item 转换为 Namespace：%s", err.Error())
				return err
			}
		} else {
			return errors.New("item 类型不支持")
		}
	} else {
		return errors.New("item 为空")
	}

	// Step 2: 强制设置 Name
	ns.Name = basicInfo.Name

	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		zap.S().Errorf("创建命名空间失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("创建命名空间%s成功!!!", basicInfo.Name)
	return nil
}

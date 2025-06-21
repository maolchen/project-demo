package namespace

import (
	"context"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"
)

func CreateNamespace(basicInfo *models.BasicInfo) error {
	var namespace corev1.Namespace
	clientset, err := models.NewClientSetForBasicInfo(basicInfo)
	if err != nil {
		return err
	}
	namespace.Name = basicInfo.Name
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
	if err != nil {
		zap.S().Errorf("创建命名空间失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("创建命名空间%s成功!!!", basicInfo.Name)
	return nil
}

package namespace

import (
	"context"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func UpdateNamespace(basicInfo *models.BasicInfo) error {
	clientset, err := models.NewClientSetForBasicInfo(basicInfo)
	if err != nil {
		return err
	}

	_, err = clientset.CoreV1().Namespaces().Update(context.TODO(), basicInfo.Item.(*corev1.Namespace), metav1.UpdateOptions{})
	if err != nil {
		zap.S().Errorf("更新命名空间失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}

	zap.S().Infof("更新命名空间%s成功!!!", basicInfo.Name)
	return nil

}

package namespace

import (
	"context"
	"errors"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeleteNamespace(basicInfo *models.BasicInfo) error {

	clientset, err := models.NewClientSetForBasicInfo(basicInfo)
	if err != nil {
		return err
	}

	if basicInfo.Namespace == "kube-system" {
		zap.S().Errorf("删除命名空间失败,不允许删除%s", basicInfo.Name)
		return errors.New("不允许删除kube-system")
	}

	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), basicInfo.Name, metav1.DeleteOptions{})
	if err != nil {
		zap.S().Errorf("删除命名空间失败：%s----->%s", basicInfo.Name, err.Error())
		return err
	}
	zap.S().Infof("删除命名空间%s成功!!!", basicInfo.Name)
	return nil
}

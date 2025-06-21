package namespace

import (
	"context"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListNamespace(basicInfo *models.BasicInfo) (interface{}, error) {

	clientset, err := models.NewClientSetForBasicInfo(basicInfo)
	if err != nil {
		return nil, err
	}

	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.S().Errorf("获取命名空间失败：%s----->%s", basicInfo.Name, err.Error())
		return nil, err
	}
	items := namespaceList.Items
	zap.S().Infof("获取命名空间%s成功!!!", basicInfo.Name)
	return items, nil
}

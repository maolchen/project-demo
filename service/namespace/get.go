package namespace

import (
	"context"
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespace(basicInfo *models.BasicInfo) (interface{}, error) {

	clientset, err := models.NewClientSetForBasicInfo(basicInfo)
	if err != nil {
		return nil, err
	}

	ns, err := clientset.CoreV1().Namespaces().Get(context.TODO(), basicInfo.Name, metav1.GetOptions{})
	if err != nil {
		zap.S().Errorf("获取命名空间详情失败：%s----->%s", basicInfo.Name, err.Error())
		return nil, err
	}

	zap.S().Infof("获取命名空间%s成功!!!", basicInfo.Name)
	return ns, nil
}

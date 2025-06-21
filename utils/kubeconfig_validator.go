package utils

import (
	"fmt"
	. "github.com/maolchen/krm-backend/service/common"
	"go.uber.org/zap"
)

func ValidateKubeconfig(kubeconfig []byte) error {

	// 解析 kubeconfig
	clientset, err := NewClientSet(kubeconfig)
	if err != nil {
		return err
	}

	// 尝试调用一个简单的 API 来验证连接
	serverVersion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return fmt.Errorf("无法访问集群，请检查权限或网络: %v", err)
	}

	// 可选：打印节点列表以确认成功连接
	zap.S().Infof("Successfully connected to clusters,当前集群版本%s", serverVersion.String())

	return nil
}

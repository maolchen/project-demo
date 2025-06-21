package clusters

import (
	"encoding/base64"
	"github.com/maolchen/krm-backend/utils"
	"go.uber.org/zap"
)

func DecodeAndValidateKubeconfig(base64Kubeconfig string) ([]byte, error) {
	kubeconfig, err := base64.StdEncoding.DecodeString(base64Kubeconfig)
	if err != nil {
		zap.L().Error("kubeconfig 解码失败", zap.Error(err))
		return nil, err
	}

	if err := utils.ValidateKubeconfig(kubeconfig); err != nil {
		zap.L().Error("kubeconfig 校验失败", zap.Error(err))
		return nil, err
	}

	return kubeconfig, nil
}

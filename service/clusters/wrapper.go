package clusters

import (
	"github.com/maolchen/krm-backend/models"
	"go.uber.org/zap"
)

type ClusterOp func(*models.ClusterInfo) error
type ClusterNameOp func(string) error

func WithLog(op string, fn ClusterOp) ClusterOp {
	return func(cluster *models.ClusterInfo) error {
		zap.S().Info("执行集群操作: " + op)
		err := fn(cluster)
		if err != nil {
			zap.S().Errorw("集群操作失败", "operation", op, "error", err)
		} else {
			zap.S().Infow("集群操作成功", "operation", op)
		}
		return err
	}
}

func WithNameLog(op string, fn ClusterNameOp) ClusterNameOp {
	return func(name string) error {
		zap.S().Info("执行集群操作: " + op)
		err := fn(name)
		if err != nil {
			zap.S().Errorw("集群操作失败", "operation", op, "error", err)
		} else {
			zap.S().Infow("集群操作成功", "operation", op)
		}
		return err
	}
}

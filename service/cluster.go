package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/maolchen/krm-backend/models"
	"github.com/maolchen/krm-backend/utils"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strconv"
)

func AddCluster(cluster *models.ClusterInfo) error {
	zap.S().Info("添加集群...")
	// 校验 kubeconfig
	if err := utils.ValidateKubeconfig(cluster.Kubeconfig); err != nil {
		zap.L().Error("kubeconfig 验证失败", zap.Error(err))
		return err
	}
	return cluster.Insert()
}

func ListCluster() ([]models.ClusterListResponse, error) {
	zap.S().Info("查看所有集群...")
	clusters, err := models.GetAllClusters()
	if err != nil {
		return nil, err
	}
	var res []models.ClusterListResponse
	for _, c := range clusters {
		status := map[string]string{
			"reachable":  "false",
			"readyNodes": "0",
			"notReady":   "0",
			"pods":       "0",
		}

		// 1. 构建 kubeconfig client
		fmt.Println(status["构建kubconfig"])
		kc, _ := base64.StdEncoding.DecodeString(c.Kubeconfig)
		config, _ := clientcmd.RESTConfigFromKubeConfig(kc)

		// 2. 创建 clientset
		fmt.Println(status["创建clientset"])
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			zap.L().Warn("创建 clientset 失败", zap.String("cluster", c.Name), zap.Error(err))
			res = append(res, models.ClusterListResponse{
				Name:   c.Name,
				Label:  c.Label,
				Status: status,
			})
			continue
		}

		// 3. 测试是否可访问集群（比如 list nodes）
		fmt.Println(status["获取集群节点"])
		nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			zap.L().Warn("无法访问集群节点", zap.String("cluster", c.Name), zap.Error(err))
			res = append(res, models.ClusterListResponse{
				Name:   c.Name,
				Label:  c.Label,
				Status: status,
			})
			continue
		}

		// 更新 reachable 状态
		status["reachable"] = "true"

		// 4. 统计 Ready 和 NotReady 节点
		fmt.Println("统计节点信息")
		readyCount := 0
		notReadyCount := 0
		for _, node := range nodes.Items {
			isReady := false
			for _, cond := range node.Status.Conditions {
				if cond.Type == corev1.NodeReady && cond.Status == corev1.ConditionTrue {
					isReady = true
					break
				}
			}
			if isReady {
				readyCount++
			} else {
				notReadyCount++
			}
		}

		status["readyNodes"] = strconv.Itoa(readyCount)
		status["notReady"] = strconv.Itoa(notReadyCount)

		// 5. 获取所有 Pod 数量
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			zap.L().Warn("无法列出 pods", zap.String("cluster", c.Name), zap.Error(err))
		} else {
			status["pods"] = strconv.Itoa(len(pods.Items))
		}

		res = append(res, models.ClusterListResponse{
			Name:   c.Name,
			Label:  c.Label,
			Status: status,
		})
	}
	return res, nil
}

func UpdateClusterByName(name string, updates map[string]interface{}) error {
	zap.S().Info("更新集群...")
	clusterInfo := models.ClusterInfo{}

	if err := utils.ValidateKubeconfig(updates["kubeconfig"].(string)); err != nil {
		zap.L().Error("kubeconfig 验证失败", zap.Error(err))
		return err
	}
	return clusterInfo.Update(name, updates)
}

func DeleteCluster(name string) error {
	zap.S().Info("删除集群...")
	clusterInfo := models.ClusterInfo{}
	return clusterInfo.Delete(name)
}

func GetClusterEditByName(name string) (models.ClusterEditResponse, error) {
	zap.S().Info("编辑集群...")
	return (&models.ClusterInfo{}).GetByName(name)
}

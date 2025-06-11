package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

// ValidateKubeconfig 验证 kubeconfig 是否合法并可访问集群，带有超时控制
func ValidateKubeconfig(encodedKubeconfig string) error {
	// 解码 Base64
	kubeconfig, err := base64.StdEncoding.DecodeString(encodedKubeconfig)
	if err != nil {
		return fmt.Errorf("无效的 Base64 编码: %v", err)
	}

	// 设置上下文超时时间为10秒
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 解析 kubeconfig
	config, err := clientcmd.RESTConfigFromKubeConfig(kubeconfig)
	if err != nil {
		return fmt.Errorf("无效的 kubeconfig 格式: %v", err)
	}

	// 创建 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("无法创建 Kubernetes 客户端: %v", err)
	}

	// 尝试调用一个简单的 API 来验证连接
	nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("无法访问集群，请检查权限或网络: %v", err)
	}

	// 可选：打印节点列表以确认成功连接
	fmt.Printf("Successfully connected to cluster, found %d nodes\n", len(nodes.Items))

	return nil
}

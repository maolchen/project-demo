package factory

import (
	"github.com/maolchen/krm-backend/service"
	resources "github.com/maolchen/krm-backend/service/resource"
	"k8s.io/client-go/kubernetes"
)

type ResourceFactory struct{}

func (f *ResourceFactory) GetService(resourceType string, clientset kubernetes.Interface) service.ResourceService {
	switch resourceType {
	case "namespaces":
		return resources.NewNamespaceService(clientset)
	case "pods":

		return resources.NewPodService(clientset)
	case "deployments":

		return resources.NewDeploymentService(clientset)
	case "statefulsets":

		return resources.NewStatefulSetService(clientset)
	case "daemonsets":

		return resources.NewDaemonSetService(clientset)
	case "cronjobs":

		return resources.NewCronJobService(clientset)
	default:

		panic("不支持的资源类型：" + resourceType)
	}
}

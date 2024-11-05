package global

import (
	"k8s.io/client-go/kubernetes"
	"kubeimook/config"
)

var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)

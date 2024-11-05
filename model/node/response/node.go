package response

import (
	corev1 "k8s.io/api/core/v1"
	"kubeimook/model/base"
)

type Node struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Age        int64  `json:"age"`
	InternalIp string `json:"internalIp"`
	ExternalIp string `json:"externalIp"`
	//kubelet 版本
	Version          string `json:"version"`
	OsImage          string `json:"osImage"`
	KernelVersion    string `json:"kernelVersion"`
	ContainerRuntime string `json:"containerRuntime"`
	// 标签和污点
	Labels []base.ListMapItem `json:"labels"`
	Taints []corev1.Taint     `json:"taints"`
}

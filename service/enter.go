package service

import (
	"kubeimook/service/configmap"
	"kubeimook/service/node"
	"kubeimook/service/pod"
)

type ServiceGroup struct {
	PodServiceGroup       pod.PodServiceGroup
	NodeServiceGroup      node.Group
	ConfigMapServiceGroup configmap.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

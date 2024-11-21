package service

import (
	"kubeimook/service/configmap"
	"kubeimook/service/node"
	"kubeimook/service/pod"
	"kubeimook/service/secret"
)

type ServiceGroup struct {
	PodServiceGroup       pod.PodServiceGroup
	NodeServiceGroup      node.Group
	ConfigMapServiceGroup configmap.ServiceGroup
	SecretServiceGroup    secret.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

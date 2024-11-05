package service

import (
	"kubeimook/service/node"
	"kubeimook/service/pod"
)

type ServiceGroup struct {
	PodServiceGroup  pod.PodServiceGroup
	NodeServiceGroup node.Group
}

var ServiceGroupApp = new(ServiceGroup)

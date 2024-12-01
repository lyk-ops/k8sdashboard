package service

import (
	"kubeimook/service/configmap"
	"kubeimook/service/node"
	"kubeimook/service/pod"
	"kubeimook/service/pv"
	"kubeimook/service/pvc"
	"kubeimook/service/secret"
	"kubeimook/service/storageclass"
)

type ServiceGroup struct {
	PodServiceGroup          pod.PodServiceGroup
	NodeServiceGroup         node.Group
	ConfigMapServiceGroup    configmap.ServiceGroup
	SecretServiceGroup       secret.ServiceGroup
	PvServiceGroup           pv.ServiceGroup
	PvcServiceGroup          pvc.ServiceGroup
	StorageClassServiceGroup storageclass.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

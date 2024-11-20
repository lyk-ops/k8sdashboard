package k8s

import (
	"kubeimook/service"
	"kubeimook/validate"
)

type ApiGroup struct {
	PodApi
	NamespaceApi
	NodeApi
	ConfigMapApi
}

var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.ServiceGroupApp.PodServiceGroup.PodService
var nodeService = service.ServiceGroupApp.NodeServiceGroup.NodeService
var configMapService = service.ServiceGroupApp.ConfigMapServiceGroup.ConfigMapService

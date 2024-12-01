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
	SecretApi
	PVApi
	PVCApi
	StorageClassApi
}

var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.ServiceGroupApp.PodServiceGroup.PodService
var nodeService = service.ServiceGroupApp.NodeServiceGroup.NodeService
var configMapService = service.ServiceGroupApp.ConfigMapServiceGroup.ConfigMapService
var secretService = service.ServiceGroupApp.SecretServiceGroup.SecretService
var pvService = service.ServiceGroupApp.PvServiceGroup.PVService
var pvcService = service.ServiceGroupApp.PvcServiceGroup.PVCService
var storageClassService = service.ServiceGroupApp.StorageClassServiceGroup.StorageClassService

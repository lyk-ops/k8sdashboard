package api

import (
	"kubeimook/api/example"
	"kubeimook/api/k8s"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8sApiGroup     k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

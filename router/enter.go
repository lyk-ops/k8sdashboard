package router

import (
	"kubeimook/router/example"
	"kubeimook/router/k8s"
)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
	K8sRouterGroup     k8s.K8sRouter
}

var RouterGroupApp = new(RouterGroup)

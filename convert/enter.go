package convert

import (
	"kubeimook/convert/configmap"
	"kubeimook/convert/node"
	"kubeimook/convert/pod"
)

type ConvertGroup struct {
	PodConvert       pod.PodConvertGroup
	NodeConvert      node.Group
	ConfigMapConvert configmap.Req2K8s
	ConfigConvert    configmap.K8s2Res
}

var ConvertGroupApp = new(ConvertGroup)

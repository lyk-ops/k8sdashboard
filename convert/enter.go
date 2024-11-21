package convert

import (
	"kubeimook/convert/configmap"
	"kubeimook/convert/node"
	"kubeimook/convert/pod"
	"kubeimook/convert/secret"
)

type ConvertGroup struct {
	PodConvert       pod.PodConvertGroup
	NodeConvert      node.Group
	ConfigMapConvert configmap.Req2K8s
	ConfigConvert    configmap.K8s2Res
	//configMapConvert configmap.ConvertGroup
	SecretConvert secret.ConvertGroup
}

var ConvertGroupApp = new(ConvertGroup)

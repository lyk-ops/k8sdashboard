package convert

import (
	"kubeimook/convert/node"
	"kubeimook/convert/pod"
)

type ConvertGroup struct {
	PodConvert  pod.PodConvertGroup
	NodeConvert node.Group
}

var ConvertGroupApp = new(ConvertGroup)

package node

import (
	"kubeimook/convert"
)

type Group struct {
	NodeService
}

var nodeConvert = convert.ConvertGroupApp.NodeConvert

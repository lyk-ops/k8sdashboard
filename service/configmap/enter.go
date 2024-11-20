package configmap

import "kubeimook/convert"

type ServiceGroup struct {
	ConfigMapService
}

var configMapConvert = convert.ConvertGroupApp.ConfigMapConvert
var ConfigConvert = convert.ConvertGroupApp.ConfigConvert

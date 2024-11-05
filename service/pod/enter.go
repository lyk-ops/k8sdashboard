package pod

import "kubeimook/convert"

var podConvert = convert.ConvertGroupApp.PodConvert

type PodServiceGroup struct {
	PodService
}

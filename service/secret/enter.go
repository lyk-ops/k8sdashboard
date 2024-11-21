package secret

import "kubeimook/convert"

type ServiceGroup struct {
	SecretService
}

var secretConvert = convert.ConvertGroupApp.SecretConvert

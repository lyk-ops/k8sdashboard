package configmap

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	configmapreq "kubeimook/model/configmap/request"
	"kubeimook/utils"
)

type Req2K8s struct {
}

func (c *Req2K8s) Cm2K8sReqConvert(comfigMap configmapreq.ConfigMap) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      comfigMap.Name,
			Namespace: comfigMap.Namespace,
			Labels:    utils.ToMap(comfigMap.Labels),
		},
		Data: utils.ToMap(comfigMap.Data),
	}
}

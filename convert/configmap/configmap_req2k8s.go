package configmap

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimook/model/base"
	configmapreq "kubeimook/model/configmap/request"
)

type Req2K8s struct {
}

func (c *Req2K8s) Cm2K8sReqConvert(comfigMap configmapreq.ConfigMap) *corev1.ConfigMap {
	var mapItem base.ListMapItem
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      comfigMap.Name,
			Namespace: comfigMap.Namespace,
			Labels:    mapItem.ToMap(comfigMap.Labels),
		},
		Data: mapItem.ToMap(comfigMap.Data),
	}
}

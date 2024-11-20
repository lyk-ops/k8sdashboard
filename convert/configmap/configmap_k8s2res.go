package configmap

import (
	corev1 "k8s.io/api/core/v1"
	"kubeimook/model/base"
	configmapres "kubeimook/model/configmap/response"
)

type K8s2Res struct{}

func (*K8s2Res) GetCmReqItem(configMap corev1.ConfigMap) configmapres.ConfigMap {
	return configmapres.ConfigMap{
		Name:      configMap.ObjectMeta.Name,
		Namespace: configMap.ObjectMeta.Namespace,
		DataNum:   len(configMap.Data),
		Age:       configMap.CreationTimestamp.Unix(),
	}

}
func (this *K8s2Res) GetCmReqDetail(configMap corev1.ConfigMap) configmapres.ConfigMap {
	detail := this.GetCmReqItem(configMap)
	var baseListMapItem base.ListMapItem
	detail.Labels = baseListMapItem.ToList(configMap.Labels)
	detail.Data = baseListMapItem.ToList(configMap.Data)
	return detail

}

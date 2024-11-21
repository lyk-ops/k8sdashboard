package configmap

import (
	corev1 "k8s.io/api/core/v1"
	configmapres "kubeimook/model/configmap/response"
	"kubeimook/utils"
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
	detail.Labels = utils.ToList(configMap.Labels)
	detail.Data = utils.ToList(configMap.Data)
	return detail

}

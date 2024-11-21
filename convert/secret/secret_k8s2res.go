package secret

import (
	corev1 "k8s.io/api/core/v1"
	secretRes "kubeimook/model/secret/response"
	"kubeimook/utils"
)

type K8s2Res struct {
}

func (K8s2Res) SecretK8sResItemConvert(secret corev1.Secret) secretRes.Secret {
	return secretRes.Secret{
		Name:      secret.Name,
		Namespace: secret.Namespace,
		DataNum:   len(secret.Data),
		Age:       secret.CreationTimestamp.Unix(),
		Type:      secret.Type,
	}
}
func (K8s2Res) SecretK8s2ResDetailConvert(secret corev1.Secret) secretRes.Secret {
	return secretRes.Secret{
		Name:      secret.Name,
		Namespace: secret.Namespace,
		DataNum:   len(secret.StringData),
		Age:       secret.CreationTimestamp.Unix(),
		Type:      secret.Type,
		Data:      utils.ToListWithMapByte(secret.Data),
		Labels:    utils.ToList(secret.Labels),
	}
}

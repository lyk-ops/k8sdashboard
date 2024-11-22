package secret

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimook/global"
	secretreq "kubeimook/model/secret/request"
	secretRes "kubeimook/model/secret/response"
	"strings"
)

type SecretService struct {
}

func (SecretService) GetSecretDetail(namespace, name string) (*secretRes.Secret, error) {
	secretK8s, err := global.KubeConfigSet.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	secretRes := secretConvert.SecretK8s2ResDetailConvert(*secretK8s)
	return &secretRes, nil

}
func (SecretService) GetSecretList(namespace string, keyword string) ([]secretRes.Secret, error) {
	list, err2 := global.KubeConfigSet.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err2 != nil {
		return nil, err2
	}
	secretResList := make([]secretRes.Secret, 0)
	//for _, item := range list.Items {
	//	if keyword == "" || (item.Name+item.Namespace) == keyword {
	//		secretRes := secretConvert.SecretK8sResItemConvert(item)
	//		secretResList = append(secretResList, secretRes)
	//	}
	//}
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		secretRes := secretConvert.SecretK8sResItemConvert(item)
		secretResList = append(secretResList, secretRes)
	}
	return secretResList, nil

}
func (SecretService) CreateOrUpdateSecret(secret secretreq.Secret) (err error) {

	secretsApi := global.KubeConfigSet.CoreV1().Secrets(secret.Namespace)
	secretK8s := secretConvert.SecretReq2K8sConvert(secret)
	//查询是否存在Secret
	_, err = secretsApi.Get(context.TODO(), secret.Name, metav1.GetOptions{})
	if err == nil {
		_, err = secretsApi.Update(context.TODO(), &secretK8s, metav1.UpdateOptions{})
	} else {
		_, err = secretsApi.Create(context.TODO(), &secretK8s, metav1.CreateOptions{})
	}

	return
}
func (SecretService) DeleteSecret(namespace, name string) (err error) {
	return global.KubeConfigSet.CoreV1().Secrets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

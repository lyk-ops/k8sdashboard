package configmap

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimook/global"
	configmapreq "kubeimook/model/configmap/request"
	configmapres "kubeimook/model/configmap/response"
	"strings"
)

type ConfigMapService struct {
}

func (*ConfigMapService) CreateOrUpdateConfigMap(configReq configmapreq.ConfigMap) error {
	// 将request数据转化为k8s数据
	convert := configMapConvert.Cm2K8sReqConvert(configReq)
	_, err2 := global.KubeConfigSet.CoreV1().ConfigMaps(convert.Namespace).Get(context.TODO(), convert.Name, metav1.GetOptions{})
	if err2 == nil {
		_, err := global.KubeConfigSet.CoreV1().ConfigMaps(convert.Namespace).Update(context.TODO(), convert, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	} else {
		_, err := global.KubeConfigSet.CoreV1().ConfigMaps(convert.Namespace).Create(context.TODO(), convert, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}
func (*ConfigMapService) GetConfigMapDetail(namespace string, name string) (cm configmapres.ConfigMap, err error) {
	configMapK8s, err := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return cm, err
	} else {
		// 将k8s数据转化为response数据
		cm := ConfigConvert.GetCmReqDetail(*configMapK8s)
		return cm, nil
	}
}

func (*ConfigMapService) GetConfigMapList(namespace string, keyword string) (cmList []configmapres.ConfigMap, err error) {
	// 从k8s获取数据
	list, err := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	//fmt.Printf("%v", list)

	if err != nil {
		return nil, err
	}
	// 转换为res(filter)数据
	configMapList := make([]configmapres.ConfigMap, 0)
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		configMapList = append(configMapList, ConfigConvert.GetCmReqItem(item))
	}
	fmt.Printf("查询到的 configmap为 %v", configMapList)

	return configMapList, nil
}

func (*ConfigMapService) DeleteConfigMap(ns string, name string) error {
	return global.KubeConfigSet.CoreV1().ConfigMaps(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

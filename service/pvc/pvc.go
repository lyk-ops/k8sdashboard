package pvc

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimook/global"
	"kubeimook/model/base"
	pvc_req "kubeimook/model/pvc/request"
	pvc_res "kubeimook/model/pvc/response"
	"kubeimook/utils"
	"strconv"
	"strings"
)

type PVCService struct {
}

func (*PVCService) CreatePVC(pvcReq pvc_req.PersistentVolumeClaim) error {
	pvc := corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pvcReq.Name,
			Namespace: pvcReq.Namespace,
			Labels:    utils.ToMap(pvcReq.Labels),
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: utils.ToMap(pvcReq.Selector),
			},
			AccessModes: pvcReq.AccessModes,
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(pvcReq.Capacity)) + "Gi"),
				},
			},
			StorageClassName: &pvcReq.StorageClassName,
		},
	}
	if pvc.Spec.StorageClassName != nil {
		pvc.Spec.Selector = nil
	}
	ctx := context.TODO()
	_, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(pvc.Namespace).Create(ctx, &pvc, metav1.CreateOptions{})
	return err

}
func (*PVCService) DeletePVC(namespace, name string) error {
	ctx := context.TODO()
	return global.KubeConfigSet.CoreV1().PersistentVolumeClaims(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}
func (*PVCService) GetPVCList(namespace string, keyword string) ([]pvc_res.PersistentVolumeClaim, error) {
	pvcResList := make([]pvc_res.PersistentVolumeClaim, 0)
	list, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		matchLabels := make([]base.ListMapItem, 0)
		if item.Spec.Selector != nil {
			matchLabels = utils.ToList(item.Spec.Selector.MatchLabels)
		}
		// k8s item-->response
		pvcResItem := pvc_res.PersistentVolumeClaim{
			Name:             item.Name,
			Namespace:        item.Namespace,
			Labels:           utils.ToList(item.Labels),
			AccessModes:      item.Spec.AccessModes,
			Capacity:         int32(item.Spec.Resources.Requests.Storage().Value() / 1024 / 1024),
			StorageClassName: *item.Spec.StorageClassName,
			Age:              item.CreationTimestamp.UnixMilli(),
			Volume:           item.Spec.VolumeName,
			Status:           item.Status.Phase,
			Selector:         matchLabels,
		}
		pvcResList = append(pvcResList, pvcResItem)
	}
	return pvcResList, nil
}

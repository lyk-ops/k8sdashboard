package pv

import (
	"context"
	"errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimook/global"
	pv_req "kubeimook/model/pv/request"
	pv_res "kubeimook/model/pv/response"
	"kubeimook/utils"
	"strconv"
)

type PVService struct {
}

func (PVService) CreatePV(pvReq pv_req.PersistentVolume) error {
	//参数转换
	var volumeSource corev1.PersistentVolumeSource
	switch pvReq.VolumeSource.Type {
	case "nfs":
		volumeSource.NFS = &corev1.NFSVolumeSource{
			Server:   pvReq.VolumeSource.NfsVolumeSource.NfsServer,
			Path:     pvReq.VolumeSource.NfsVolumeSource.NfsPath,
			ReadOnly: pvReq.VolumeSource.NfsVolumeSource.NfsReadOnly,
		}
	default:
		return errors.New("不支持的存储类型")
	}
	pv := corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pvReq.Name,
			Labels: utils.ToMap(pvReq.Labels),
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(pvReq.Capacity)) + "Mi"),
			},
			AccessModes:                   pvReq.AccessModes,
			PersistentVolumeSource:        volumeSource,
			PersistentVolumeReclaimPolicy: pvReq.ReclaimPolicy[0],
		},
	}
	ctx := context.TODO()
	_, err := global.KubeConfigSet.CoreV1().PersistentVolumes().Create(ctx, &pv, metav1.CreateOptions{})
	return err
}
func (PVService) DeletePV(_ string, name string) error {
	err := global.KubeConfigSet.CoreV1().PersistentVolumes().Delete(context.TODO(), name, metav1.DeleteOptions{})
	return err
}
func (PVService) GetPvList() ([]pv_res.PersistentVolume, error) {
	pvList, err := global.KubeConfigSet.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	pvResList := make([]pv_res.PersistentVolume, 0)
	for _, item := range pvList.Items {
		//k8s-->response
		claim := ""
		if item.Spec.ClaimRef != nil {
			claim = item.Spec.ClaimRef.Name
		}
		pvRes := pv_res.PersistentVolume{

			Name:             item.Name,
			Labels:           utils.ToList(item.Labels),
			Capacity:         int32(item.Spec.Capacity.Storage().Value() / 1024 / 1024),
			AccessModes:      item.Spec.AccessModes,
			ReclaimPolicy:    item.Spec.PersistentVolumeReclaimPolicy,
			Status:           item.Status.Phase,
			Claim:            claim,
			Age:              item.CreationTimestamp.Unix(),
			Reason:           item.Status.Reason,
			StorageClassName: "",
		}
		pvResList = append(pvResList, pvRes)
	}
	return pvResList, nil
}

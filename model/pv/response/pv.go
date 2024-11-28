package response

import (
	corev1 "k8s.io/api/core/v1"
	"kubeimook/model/base"
)

type PersistentVolume struct {
	Name          string                               `json:"name"`
	Capacity      int32                                `json:"capacity"`
	Labels        []base.ListMapItem                   `json:"labels"`
	AccessModes   []corev1.PersistentVolumeAccessMode  `json:"accessModes"`
	ReclaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"reclaimPolicy"`
	//VolumeSource  VolumeSource                           `json:"volumeSource"`
	Status           corev1.PersistentVolumePhase `json:"status"`
	Claim            string                       `json:"claim"`  // 绑定
	Age              int64                        `json:"age"`    // 创建时间
	Reason           string                       `json:"reason"` // 状态
	StorageClassName string                       `json:"storageClassName"`
}

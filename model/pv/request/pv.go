package request

import (
	corev1 "k8s.io/api/core/v1"
	"kubeimook/model/base"
)

type NfsVolumeSource struct {
	NfsServer   string `json:"server"`
	NfsPath     string `json:"path"`
	NfsReadOnly bool   `json:"nfsReadOnly"`
}
type VolumeSource struct {
	Type            string          `json:"type"`
	NfsVolumeSource NfsVolumeSource `json:"nfsVolumeSource"`
}
type PersistentVolume struct {
	Name string `json:"name"`
	//Namespace不需要
	//Namespace string             `json:"namespace"`
	Labels []base.ListMapItem `json:"labels"`
	// 容量
	Capacity int32 `json:"capacity"`
	// 访问模式
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	// pv回收策略
	ReclaimPolicy []corev1.PersistentVolumeReclaimPolicy `json:"reclaimPolicy"`
	// VolumeSource
	VolumeSource VolumeSource `json:"volumeSource"`
}

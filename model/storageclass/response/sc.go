package response

import (
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"kubeimook/model/base"
)

type StorageClass struct {
	Name string `json:"name"`
	//Namespace string             `json:"namespace"`
	Labels []base.ListMapItem `json:"labels"`
	//制备器
	Provisioner string             `json:"provisioner"`
	Parameters  []base.ListMapItem `json:"parameters"` //制备器入参
	//绑定选项
	MountOptions         []string                             `json:"mountOptions"`         //挂载选项
	ReclaimPolicy        corev1.PersistentVolumeReclaimPolicy `json:"reclaimPolicy"`        //回收策略
	AllowVolumeExpansion bool                                 `json:"allowVolumeExpansion"` //是否允许卷扩展
	VolumeBindingMode    storagev1.VolumeBindingMode          `json:"volumeBindingMode"`
	Age                  int64                                `json:"age"` //创建时间
}

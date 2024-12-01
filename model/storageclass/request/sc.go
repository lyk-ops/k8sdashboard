package request

import (
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"kubeimook/model/base"
)

/*
	type StorageClass struct {
	    v1.TypeMeta      `json:",inline"`
	    v1.ObjectMeta    `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	    Provisioner          string                            `json:"provisioner" protobuf:"bytes,2,opt,name=provisioner"`
	    Parameters           map[string]string                 `json:"parameters,omitempty" protobuf:"bytes,3,rep,name=parameters"`
	    ReclaimPolicy        *v1.PersistentVolumeReclaimPolicy `json:"reclaimPolicy,omitempty" protobuf:"bytes,4,opt,name=reclaimPolicy,casttype=k8s.io/api/core/v1.PersistentVolumeReclaimPolicy"`
	    MountOptions         []string                          `json:"mountOptions,omitempty" protobuf:"bytes,5,opt,name=mountOptions"`
	    AllowVolumeExpansion *bool                             `json:"allowVolumeExpansion,omitempty" protobuf:"varint,6,opt,name=allowVolumeExpansion"`
	    VolumeBindingMode    *VolumeBindingMode                `json:"volumeBindingMode,omitempty" protobuf:"bytes,7,opt,name=volumeBindingMode"`
	    AllowedTopologies    []v1.TopologySelectorTerm         `json:"allowedTopologies,omitempty" protobuf:"bytes,8,rep,name=allowedTopologies"`
	}
*/
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
}

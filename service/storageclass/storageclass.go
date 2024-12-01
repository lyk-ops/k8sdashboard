package storageclass

import (
	"context"
	"fmt"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimook/global"
	sc_req "kubeimook/model/storageclass/request"
	sc_res "kubeimook/model/storageclass/response"
	"kubeimook/utils"
	"strings"
)

type StorageClassService struct {
}

func (*StorageClassService) GetStorageClasses(keyword string) ([]sc_res.StorageClass, error) {

	list, err := global.KubeConfigSet.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var scList []sc_res.StorageClass
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		//item --> response
		if item.AllowVolumeExpansion == nil {
			item.AllowVolumeExpansion = new(bool)
		}
		mountOptions := make([]string, 0)
		if item.MountOptions != nil {
			mountOptions = item.MountOptions
		}

		scList = append(scList, sc_res.StorageClass{
			Name:                 item.Name,
			Labels:               utils.ToList(item.Labels),
			Provisioner:          item.Provisioner,
			Parameters:           utils.ToList(item.Parameters),
			MountOptions:         mountOptions,
			ReclaimPolicy:        *item.ReclaimPolicy,
			AllowVolumeExpansion: *item.AllowVolumeExpansion,
			VolumeBindingMode:    *item.VolumeBindingMode,
			Age:                  item.CreationTimestamp.Unix(),
		})
	}
	return scList, nil
}
func (*StorageClassService) CreateStorageClass(scReq sc_req.StorageClass) error {

	//判断 provisioner 参数是否合法
	provisioner := strings.Split(global.CONF.System.Provisioner, ",")
	for _, item := range provisioner {
		if scReq.Provisioner == item {
			break
		} else if item == provisioner[len(provisioner)-1] {
			err := fmt.Errorf("provisioner 参数错误,期望值: " + strings.Join(provisioner, ","))
			return err

		}
	}
	sc := storagev1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:   scReq.Name,
			Labels: utils.ToMap(scReq.Labels),
		},
		Provisioner:          scReq.Provisioner,
		MountOptions:         scReq.MountOptions,
		ReclaimPolicy:        &scReq.ReclaimPolicy,
		AllowVolumeExpansion: &scReq.AllowVolumeExpansion,
		VolumeBindingMode:    &scReq.VolumeBindingMode,
		Parameters:           utils.ToMap(scReq.Parameters),
	}
	ctx := context.TODO()
	_, err := global.KubeConfigSet.StorageV1().StorageClasses().Create(ctx, &sc, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
func (*StorageClassService) DeleteStorageClass(name string) error {
	ctx := context.TODO()
	err := global.KubeConfigSet.StorageV1().StorageClasses().Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

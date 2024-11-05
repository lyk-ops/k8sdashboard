package validate

import (
	"errors"
	pod_req "kubeimook/model/pod/request"
)

const (
	IMAGE_PULL_PILICY_INNOTPRESENT = "IfNotPresent"
	RESTART_POLICY_ALWAYS          = "Always"
)

type PodValidate struct{}

func (*PodValidate) Validate(podReq *pod_req.Pod) error {
	// 校验必填项
	if podReq.Base.Name == "" {
		return errors.New("请定义pod名字")
	}
	if len(podReq.Containers) == 0 {
		return errors.New("请定义pod容器")
	}
	// 非必填项赋值默认值
	if len(podReq.Containers) > 0 {
		for index, container := range podReq.InitContainers {
			if container.Name == "" {
				return errors.New("请定义初始化容器名字")
			}
			if container.Image == "" {
				return errors.New("请定义初始化容器镜像")
			}
			if container.ImagePullPolicy == "" {
				podReq.InitContainers[index].ImagePullPolicy = IMAGE_PULL_PILICY_INNOTPRESENT
			}
		}
	}
	if len(podReq.Containers) > 0 {
		for index, container := range podReq.Containers {
			if container.Name == "" {
				return errors.New("请定义容器名字")
			}
			if container.Image == "" {
				return errors.New("请定义容器镜像")
			}
			if container.ImagePullPolicy == "" {
				podReq.InitContainers[index].ImagePullPolicy = IMAGE_PULL_PILICY_INNOTPRESENT
			}
		}
	}
	if podReq.Base.RestartPolicy == "" {
		podReq.Base.RestartPolicy = RESTART_POLICY_ALWAYS
	}

	return nil
}

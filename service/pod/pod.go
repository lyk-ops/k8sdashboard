package pod

import (
	"context"
	"errors"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	k8serror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"kubeimook/global"
	pod_req "kubeimook/model/pod/request"
	pod_res "kubeimook/model/pod/response"
	"strings"
)

type PodService struct {
}

func (*PodService) GetPodList(namespace string, keyword string, nodename string) ([]pod_res.PodListItem, error) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	podList := make([]pod_res.PodListItem, len(list.Items))
	for _, item := range list.Items {
		if nodename != "" && item.Spec.NodeName != nodename {
			continue
		}
		if strings.Contains(item.Name, keyword) {
			podRItem := podConvert.PodK8s2ItemRes(item)
			podList = append(podList, podRItem)
		}

	}
	return podList, nil
}
func (*PodService) GetPodDetail(namespace, name string) (podReq pod_req.Pod, err error) {
	ctx := context.TODO()
	podApi := global.KubeConfigSet.CoreV1().Pods(namespace)
	k8sGetpod, err := podApi.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]查询失败，detail: %s", namespace, name, err)
		err = errors.New(errMsg)
		return
	}
	// 将 k8s 对象转换为自定义对象
	podReq = podConvert.K8s2ReqConvert.PodK8s2Req(*k8sGetpod)
	return
}
func (*PodService) CreateOrUpdatePod(podReq pod_req.Pod) (msg string, err error) {
	k8sPod := podConvert.Req2K8sConvert.PodReq2K8s(podReq)
	ctx := context.TODO()
	podApi := global.KubeConfigSet.CoreV1().Pods(k8sPod.Namespace)
	// [no]update [no] patch [yes] delete+create
	if k8sGetPod, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); err == nil {
		//参数是否合理
		k8sPodCopy := *k8sPod
		k8sPodCopy.Name = k8sPod.Name + "-validate"
		_, err := podApi.Create(ctx, &k8sPodCopy, metav1.CreateOptions{
			DryRun: []string{metav1.DryRunAll},
		})
		if err != nil {
			errMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]创建失败，detail: %s", k8sPod.Namespace, k8sPod.Name, err)
			return errMsg, nil
		}
		//删除操作 -- 强制删除
		background := metav1.DeletePropagationBackground
		var gracePeriodSeconds int64 = 0
		err = podApi.Delete(ctx, k8sPod.Name, metav1.DeleteOptions{
			GracePeriodSeconds: &gracePeriodSeconds, //可以配置缩短时间强制删除
			PropagationPolicy:  &background,
		})
		if err != nil {
			errMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]删除失败，detail: %s", k8sPod.Namespace, k8sPod.Name, err)
			return errMsg, nil
		}
		//创建pod
		//pod可能处于terminating状态，监听pod删除完成后才能执行create
		var labelSelector []string
		for k, v := range k8sGetPod.Labels {
			labelSelector = append(labelSelector, fmt.Sprintf("%s=%s", k, v))
		}
		//label格式 k1=v1,k2=v2
		watcher, err := podApi.Watch(ctx, metav1.ListOptions{
			LabelSelector: strings.Join(labelSelector, ","),
		})
		if err != nil {
			errMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]监听失败，detail: %s", k8sPod.Namespace, k8sPod.Name, err)
			return errMsg, nil
		}
		for event := range watcher.ResultChan() {
			switch event.Type {
			case watch.Deleted:
				k8sPodChan := event.Object.(*corev1.Pod)
				//有时terminating状态时间会很短，不需要等待查询很长时间，直接创建pod
				if _, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{}); k8serror.IsNotFound(err) {
					createPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
					if err != nil {
						errMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]创建失败，detail: %s", k8sPod.Namespace, k8sPod.Name, err)
						return errMsg, nil
					} else {
						successMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]创建成功", createPod.Namespace, createPod.Name)
						return successMsg, nil
					}
				}
				//重新创建pod
				if k8sPodChan.Name != k8sPod.Name {
					continue
				}
				createPod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
				if err != nil {
					errMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]创建失败，detail: %s", k8sPod.Namespace, k8sPod.Name, err)
					return errMsg, nil
				} else {
					successMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]创建成功", createPod.Namespace, createPod.Name)
					return successMsg, nil
				}
			}
		}
		return "", nil
	} else {
		pod, err := podApi.Create(ctx, k8sPod, metav1.CreateOptions{})
		if err != nil {
			errMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]创建失败，detail: %s", k8sPod.Namespace, k8sPod.Name, err)
			return errMsg, nil
		} else {
			successMsg := fmt.Sprintf("pod[namespace:%s-pod:%s]创建成功", pod.Namespace, pod.Name)
			return successMsg, nil
		}
	}
	return "", nil
}
func (*PodService) DeletePod(namespace, name string) error {
	ctx := context.TODO()
	background := metav1.DeletePropagationBackground
	var gracePeriodSeconds int64 = 0
	return global.KubeConfigSet.CoreV1().Pods(namespace).Delete(ctx, name, metav1.DeleteOptions{
		GracePeriodSeconds: &gracePeriodSeconds,
		PropagationPolicy:  &background,
	})
}

package k8s

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"kubeimook/global"
	pod_req "kubeimook/model/pod/request"
	"kubeimook/response"
)

type PodApi struct {
}

// update pod  update 注意：update操作受限，只能更新部分字段
func UpdatePod(ctx context.Context, pod *corev1.Pod) error {
	// update 注意：update操作受限，只能更新部分字段
	_, err := global.KubeConfigSet.CoreV1().Pods(pod.Namespace).Update(ctx, pod, metav1.UpdateOptions{})
	if err != nil {
		return err
	} else {
		return nil
	}
	return nil
}
func PatchPod(patchData map[string]interface{}, k8sPod *corev1.Pod, ctx context.Context) error {

	patchDataBytes, err := json.Marshal(patchData)
	if err != nil {
		panic(err)
	}
	_, err = global.KubeConfigSet.CoreV1().Pods(k8sPod.Namespace).Patch(
		ctx,
		k8sPod.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	return err
}
func (*PodApi) CreateOrUpdatePod(c *gin.Context) {
	var podReq pod_req.Pod

	if err := c.ShouldBind(&podReq); err != nil {
		response.FailWithMessage(c, "参数解析失败"+err.Error())
		return
	}
	err := podValidate.Validate(&podReq)
	if err != nil {
		response.FailWithMessage(c, "参数验证失败，detail: "+err.Error())
		return
	}
	msg, err := podService.CreateOrUpdatePod(podReq)
	if err != nil {
		response.FailWithMessage(c, msg+err.Error())
	} else {
		response.SuccessWithMessage(c, msg)
	}
}

func (*PodApi) GetPodListOrDetail(c *gin.Context) {
	namespace := c.Param("namespace") // c.Param用于从URL的路径参数中提取值
	name := c.Query("name")           //用于从URL的查询字符串中提取值
	keyword := c.Query("keyword")
	if name != "" {
		req, err := podService.GetPodDetail(namespace, name)
		if err != nil {
			response.FailWithMessage(c, "查询失败"+err.Error())
		} else {
			response.SuccessWithDetailed(c, "获取成功", req)
		}
	} else {
		req, err := podService.GetPodList(namespace, keyword, c.Query("nodename"))
		if err != nil {
			response.FailWithMessage(c, "查询失败"+err.Error())
		} else {
			response.SuccessWithDetailed(c, "获取成功", req)
		}

	}
}

func (*PodApi) DeletePod(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	err := podService.DeletePod(namespace, name)
	if err != nil {
		response.FailWithMessage(c, "删除失败"+err.Error())
	} else {
		response.SuccessWithMessage(c, "删除成功")
	}
}

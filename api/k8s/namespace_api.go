package k8s

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimook/global"
	namespace_res "kubeimook/model/namespace/response"
	"kubeimook/response"
)

type NamespaceApi struct {
}

func (*NamespaceApi) GetNamespaceList(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	namespaceList := make([]namespace_res.Namespace, 0)
	for _, v := range list.Items {
		namespaceList = append(namespaceList, namespace_res.Namespace{
			Name:              v.GetName(),
			CreationTimestamp: v.GetCreationTimestamp().String(),
			Status:            string(v.Status.Phase),
		})
	}
	response.SuccessWithDetailed(c, "获取成功", namespaceList)

}

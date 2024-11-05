package k8s

import (
	"github.com/gin-gonic/gin"
	"kubeimook/api"
)

type K8sRouter struct{}

func (*K8sRouter) InitK8sRouter(r *gin.Engine) {
	group := r.Group("k8s")
	apiGroup := api.ApiGroupApp.K8sApiGroup
	group.POST("/pod", apiGroup.CreateOrUpdatePod)
	group.GET("/pod/:namespace", apiGroup.GetPodListOrDetail)
	group.DELETE("/pod/:namespace/:name", apiGroup.DeletePod)
	group.GET("/namespace", apiGroup.GetNamespaceList)

	// node schedule相关
	group.GET("/node", apiGroup.GetNodeDetailOrList)

}

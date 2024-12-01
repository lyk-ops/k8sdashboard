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
	group.PUT("/node/label", apiGroup.UpdatedNodeLabel)
	group.PUT("/node/taint", apiGroup.UpdaNodeTaint)

	// ConfigMap
	group.GET("/configmap/:namespace", apiGroup.GetConfigMapDetailOrList)
	group.POST("/configmap", apiGroup.CreateOrUpdateConfigMap)
	group.DELETE("/configmap/:namespace/:name", apiGroup.DeleteConfigMap)

	//Secret
	group.POST("/secret", apiGroup.CreateOrUpdateSecret)
	group.GET("/secret/:namespace", apiGroup.GetSecretDetailOrList)
	group.DELETE("/secret/:namespace/:name", apiGroup.DeleteSecret)

	//PV
	group.POST("/pv", apiGroup.CreatePV)
	group.DELETE("/pv/:namespace/:name", apiGroup.DeletePV)
	group.GET("/pv/:name", apiGroup.GetPVList)

	//pvc
	group.POST("/pvc", apiGroup.CreatePVC)
	group.DELETE("/pvc/:namespace/:name", apiGroup.DeletePVC)
	group.GET("/pvc/:namespace", apiGroup.GetPVCList)

	//storageClass
	group.POST("/sc", apiGroup.CreateStroageClass)
	group.DELETE("/sc/:namespace/:name", apiGroup.DeleteStroageClass)
	group.GET("/sc/:name", apiGroup.ListStroageClass)
}

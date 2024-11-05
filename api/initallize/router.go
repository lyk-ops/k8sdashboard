package initallize

import (
	"github.com/gin-gonic/gin"
	"kubeimook/middleware"
	"kubeimook/router"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors)
	exampleGroup := router.RouterGroupApp.ExampleRouterGroup
	exampleGroup.InitExample(r)
	k8sGroup := router.RouterGroupApp.K8sRouterGroup // 初始化k8s路由
	k8sGroup.InitK8sRouter(r)
	return r
}

package k8s

import (
	"fmt"
	"github.com/gin-gonic/gin"
	configmapreq "kubeimook/model/configmap/request"
	"kubeimook/response"
)

type ConfigMapApi struct {
}

// GetConfigMapDetailOrList 获取configmap详情或列表
func (*ConfigMapApi) GetConfigMapDetailOrList(c *gin.Context) {
	name := c.Query("name")
	namespace := c.Param("namespace")
	keyword := c.Query("keyword")
	fmt.Printf("namespace: %s, name: %s\n", namespace, name)

	if name == "" {
		list, err := configMapService.GetConfigMapList(namespace, keyword)
		fmt.Printf("查询到的 Configmap列表: %v\n", list)
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		} else {
			response.SuccessWithDetailed(c, "查询Configmap列表成功", list)
		}
	} else {
		cm, err := configMapService.GetConfigMapDetail(namespace, name)
		if err != nil {
			response.FailWithMessage(c, "查询Configmap详情失败")
		} else {
			response.SuccessWithDetailed(c, "查询Configmap详情成功", cm)
		}
	}
}

func (*ConfigMapApi) CreateOrUpdateConfigMap(c *gin.Context) {
	var configMapReq configmapreq.ConfigMap
	err := c.ShouldBind(&configMapReq)
	if err != nil {
		response.FailWithMessage(c, "Configmap参数解析失败")
		return
	}
	err = configMapService.CreateOrUpdateConfigMap(configMapReq)
	if err != nil {
		response.FailWithMessage(c, "创建或更新Configmap失败")
	} else {
		response.Success(c)
	}
}

func (*ConfigMapApi) DeleteConfigMap(c *gin.Context) {
	err := configMapService.DeleteConfigMap(c.Param("namespace"), c.Param("name"))
	if err != nil {
		response.FailWithMessage(c, err.Error())
	} else {
		response.Success(c)
	}
}

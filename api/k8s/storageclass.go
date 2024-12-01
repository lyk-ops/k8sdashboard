package k8s

import (
	"github.com/gin-gonic/gin"
	sc_req "kubeimook/model/storageclass/request"
	"kubeimook/response"
)

type StorageClassApi struct {
}

func (*StorageClassApi) CreateStroageClass(c *gin.Context) {
	var scReq sc_req.StorageClass
	if err := c.ShouldBindJSON(&scReq); err != nil {
		response.FailWithMessage(c, "参数错误")
		return
	}
	err := storageClassService.CreateStorageClass(scReq)
	if err != nil {
		response.FailWithMessage(c, err.Error())
	} else {
		response.Success(c)
	}
}
func (*StorageClassApi) DeleteStroageClass(c *gin.Context) {
	err := storageClassService.DeleteStorageClass(c.Query("name"))
	if err != nil {
		response.FailWithMessage(c, err.Error())
	} else {
		response.Success(c)
	}
}
func (*StorageClassApi) ListStroageClass(c *gin.Context) {
	classes, err := storageClassService.GetStorageClasses(c.Query("keyword"))
	if err != nil {
		response.FailWithMessage(c, err.Error())
	} else {
		response.SuccessWithDetailed(c, "获取成功", classes)
	}

}

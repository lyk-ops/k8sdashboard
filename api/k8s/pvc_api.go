package k8s

import (
	"github.com/gin-gonic/gin"
	pvc_req "kubeimook/model/pvc/request"
	"kubeimook/response"
)

type PVCApi struct {
}

func (*PVCApi) CreatePVC(c *gin.Context) {

	var pvcReq pvc_req.PersistentVolumeClaim
	if err := c.ShouldBindJSON(&pvcReq); err != nil {
		response.FailWithMessage(c, "参数错误")
		return
	}
	err2 := pvcService.CreatePVC(pvcReq)
	if err2 != nil {
		response.FailWithMessage(c, "创建失败")
	} else {
		response.Success(c)
	}

}
func (*PVCApi) DeletePVC(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	err := pvcService.DeletePVC(namespace, name)
	if err != nil {
		response.FailWithMessage(c, "参数错误")
	} else {
		response.Success(c)
	}
}
func (*PVCApi) GetPVCList(c *gin.Context) {
	namespace := c.Param("namespace")
	keywords := c.Query("keywords")
	pvcList, err2 := pvcService.GetPVCList(namespace, keywords)
	if err2 != nil {
		response.FailWithMessage(c, "参数错误")
	} else {
		response.SuccessWithDetailed(c, "获取成功", pvcList)
	}
}

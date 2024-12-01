package k8s

import (
	"github.com/gin-gonic/gin"
	pv_req "kubeimook/model/pv/request"
	"kubeimook/response"
)

type PVApi struct {
}

func (PVApi) CreatePV(c *gin.Context) {

	var pvReq pv_req.PersistentVolume
	if err := c.ShouldBind(&pvReq); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err := pvService.CreatePV(pvReq)
	if err != nil {
		response.FailWithMessage(c, err.Error())
	} else {
		response.Success(c)
	}

}
func (PVApi) DeletePV(c *gin.Context) {
	err := pvService.DeletePV(c.Param("namespace"), c.Param("name"))
	if err != nil {
		response.FailWithMessage(c, err.Error())
	} else {
		response.Success(c)
	}
}
func (PVApi) GetPVList(c *gin.Context) {
	list, err := pvService.GetPvList(c.Query("keyword"))
	if err != nil {
		response.FailWithMessage(c, err.Error())
	} else {
		response.SuccessWithDetailed(c, "获取数据成功", list)
	}
}

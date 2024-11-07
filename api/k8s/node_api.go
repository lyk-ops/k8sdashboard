package k8s

import (
	"github.com/gin-gonic/gin"
	node_req "kubeimook/model/node/request"
	"kubeimook/response"
)

type NodeApi struct {
}

func (*NodeApi) GetNodeDetailOrList(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	nodeName := ctx.Query("nodeName")
	if nodeName != "" {
		detail, err := nodeService.GetNodeDetail(nodeName)
		if err != nil {
			response.FailWithMessage(ctx, "获取节点详情失败")
		} else {
			response.SuccessWithDetailed(ctx, "获取节点详情成功", detail)
		}
	} else {
		list, err := nodeService.GetNodeList(keyword)
		//fmt.Printf("list 列表%v", list)
		if err != nil {
			response.FailWithMessage(ctx, "获取节点列表失败")
		} else {
			response.SuccessWithDetailed(ctx, "获取节点列表成功", list)
		}
	}

	return
}
func (*NodeApi) UpdatedNodeLabel(ctx *gin.Context) {
	var updateLabels node_req.UpdatedLabel
	if err := ctx.ShouldBindJSON(&updateLabels); err != nil {
		response.FailWithMessage(ctx, "参数解析失败")
		return
	} else {
		err := nodeService.UpdateNodeLabel(updateLabels)
		if err != nil {
			response.FailWithMessage(ctx, "更新节点标签失败")
		} else {
			response.SuccessWithMessage(ctx, "更新节点标签成功")
		}
	}
}

func (*NodeApi) UpdaNodeTaint(ctx *gin.Context) {
	var updateTaints node_req.UpdatedTaint
	if err := ctx.ShouldBindJSON(&updateTaints); err != nil {
		response.FailWithMessage(ctx, "参数解析失败")
	} else {
		err := nodeService.UpdateNodeTaint(updateTaints)
		if err != nil {
			response.FailWithMessage(ctx, "更新节点污点失败")
		} else {
			response.SuccessWithMessage(ctx, "更新节点污点成功")
		}
	}
}

package k8s

import (
	"github.com/gin-gonic/gin"
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

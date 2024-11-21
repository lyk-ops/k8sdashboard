package k8s

import (
	"github.com/gin-gonic/gin"
	secretreq "kubeimook/model/secret/request"
	"kubeimook/response"
)

type SecretApi struct {
}

func (*SecretApi) CreateOrUpdateSecret(ctx *gin.Context) {
	var secretReq secretreq.Secret
	if err := ctx.ShouldBindJSON(&secretReq); err != nil {
		response.FailWithMessage(ctx, "参数绑定失败")
		return
	}
	err := secretService.CreateOrUpdateSecret(secretReq)
	if err != nil {
		response.FailWithMessage(ctx, "创建/更新失败")
		return
	}
	response.Success(ctx)
}
func (*SecretApi) GetSecretDetailOrList(ctx *gin.Context) {
	name := ctx.Query("name")
	namespace := ctx.Param("namespace")
	keyword := ctx.Query("keyword")
	var data interface{}
	var err error
	if name != "" {
		data, err = secretService.GetSecretDetail(namespace, name)
	} else {
		data, err = secretService.GetSecretList(namespace, keyword)

	}
	if err != nil {
		response.SuccessWithMessage(ctx, "获取Secret失败")
	} else {
		response.SuccessWithDetailed(ctx, "获取secret成功", data)
	}
}
func (*SecretApi) DeleteSecret(ctx *gin.Context) {

	err := secretService.DeleteSecret(ctx.Param("namespace"), ctx.Param("name"))
	if err != nil {
		response.FailWithMessage(ctx, "删除Secret失败"+err.Error())
	} else {
		response.Success(ctx)
	}

}

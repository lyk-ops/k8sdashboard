package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	success = iota
	fail
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  "success",
	})
}
func SuccessWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  message,
	})
}
func SuccessWithDetailed(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  message,
		"data": data,
	})
}
func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  "fail",
	})
}
func FailWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  message,
	})
}
func FailWithDetailed(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  message,
		"data": data,
	})
}

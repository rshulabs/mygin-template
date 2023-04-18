package utils

import (
	"backup/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应结构体
type Response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

// Success 响应成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0, data, "success",
	})
}

// Fail 响应失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode, nil, msg,
	})
}

// FailByError 返回自定义错误
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidatorFail 参数验证失败
func ValidatorFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidatorError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}

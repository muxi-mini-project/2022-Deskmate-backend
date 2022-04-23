package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Data interface{} `json:"data"`
}

// Response 请求响应
type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"msg"`
	Data    interface{} `json:"data"`
} //@name response

// 200
func SendResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

// 401
func SendResponse401(c *gin.Context, message string, data error) {
	c.JSON(http.StatusOK, Response{
		Code:    401,
		Message: message,
		Data:    data,
	})
}

// 404
func SendResponse404(c *gin.Context, message string, data error) {
	c.JSON(http.StatusOK, Response{
		Code:    404,
		Message: message,
		Data:    data,
	})
}

// 400
func SendBadRequest(c *gin.Context, message interface{}, data error) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    data,
	})
}

// 500
func SendResponse500(c *gin.Context, message string, data error) {
	c.JSON(http.StatusOK, Response{
		Code:    500,
		Message: message,
		Data:    data,
	})
}
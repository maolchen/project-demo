package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 定义了API响应的通用格式
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // 使用omitempty避免返回空对象
}

// SuccessWithData 返回带有数据的成功响应
func SuccessWithData(message string, data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
}

// SuccessNoData 返回没有数据的成功响应
func SuccessNoData(message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
	}
}

// Error 返回错误响应
func Error(code int, message string) *Response {
	return &Response{
		Status:  code,
		Message: message,
	}
}

// Send 将响应发送到Gin上下文
func (r *Response) Send(ctx *gin.Context) {
	if r.Data != nil {
		ctx.JSON(r.Status, gin.H{
			"status":  r.Status,
			"message": r.Message,
			"data":    r.Data,
		})
	} else {
		ctx.JSON(r.Status, gin.H{
			"status":  r.Status,
			"message": r.Message,
		})
	}
}

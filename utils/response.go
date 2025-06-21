package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 定义了API响应的通用格式
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty 避免空值
}

// DataItem 单个对象结构
type DataItem struct {
	Item interface{} `json:"item"`
}

// DataItems 列表结构
type DataItems struct {
	Items interface{} `json:"items"`
}

// SuccessWithItem 返回单个对象
func SuccessWithItem(message string, item interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Data: DataItem{
			Item: item,
		},
	}
}

// SuccessWithItems 返回多个对象
func SuccessWithItems(message string, items interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Data: DataItems{
			Items: items,
		},
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

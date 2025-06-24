package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/constants"
	"go.uber.org/zap"
	"net/http"
)

func BindJSON(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		Error(http.StatusBadRequest, constants.RequestParmsError).Send(ctx)
		zap.S().Errorf("解析参数失败---->%s", err.Error())
		return false
	}
	//fmt.Println(obj)
	return true
}

func BindQuery(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindQuery(obj); err != nil {
		Error(http.StatusBadRequest, constants.RequestParmsError).Send(ctx)
		zap.S().Errorf("解析参数失败---->%s", err.Error())
		return false
	}
	//fmt.Print("BindQuery", obj)
	return true
}

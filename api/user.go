package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/constants"
	"github.com/maolchen/krm-backend/service"
	"github.com/maolchen/krm-backend/utils"
	"net/http"
	"strconv"
)

// 用户认证登录
func Login(ctx *gin.Context) {
	user, err := service.UserLogin(ctx)
	if err != nil || user.Username == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  constants.CodeAuthFail,
			"message": err.Error(),
		})
		return
	}

	// 生成token
	newJwt := utils.NewJWT()

	token, err := newJwt.AccessToken(user.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  constants.CodeAuthFail,
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  constants.CodeSuccess,
		"message": constants.LoginSuccess,
		"data": gin.H{
			"token": token,
		},
	})
}

func UserAuthenticate(ctx *gin.Context) {
	var data map[string]string = map[string]string{
		"username": "chenml",
		"password": "111",
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  constants.CodeSuccess,
		"message": constants.AuthSuccess,
		"data":    data,
	})
}

func UserCreate(ctx *gin.Context) {
	var data map[string]string = map[string]string{
		"id": "",
	}
	id, err := service.CreateUser(ctx)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  constants.CodeCreateUserfail,
			"message": constants.CreateUserFail + ":" + err.Error(),
		})
		return
	}
	data["id"] = strconv.Itoa(int(id))
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  constants.CodeSuccess,
		"message": constants.CreateUserSuccess,
		"data":    data,
	})
}

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/constants"
	"github.com/maolchen/krm-backend/utils"
	"go.uber.org/zap"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUrl := c.FullPath()
		if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
			zap.S().Debugf("当前请求路径是%s", requestUrl)
			c.Next()
			return
		}
		// 获取请求头中的 token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"status":  constants.CodeAuthFail,
				"message": constants.TokenInvalid,
			})
			return
		}

		zap.S().Debugf("get token: %#v", token)

		j := utils.NewJWT() // 使用全局 SecretKey

		// 解析 token
		claims, err := j.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"status":  constants.CodeAuthFail,
				"message": err.Error(),
			})
			return
		}

		// 设置用户信息到上下文
		c.Set("claims", claims)
		c.Set("access_token", token)

		// 检查是否需要刷新 token
		newToken, err := j.RefreshToken(token)
		if err == nil && newToken != token {
			// 返回新的 token 到 header
			c.Header("X-Token", newToken)
		}

		c.Next()
	}
}

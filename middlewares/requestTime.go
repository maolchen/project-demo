package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// 记录请求时间到日志
func RequestTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		latency := time.Since(startTime).Milliseconds()
		path := c.Request.URL.Path
		method := c.Request.Method
		statusCode := c.Writer.Status()

		zap.L().Debug("Request Success.......",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status_code", statusCode),
			zap.String("client_ip", c.ClientIP()),
			zap.String("cost_time", fmt.Sprintf("%d ms", latency)),
		)
	}
}

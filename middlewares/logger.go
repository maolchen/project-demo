package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

// 记录请求时间到日志
// LoggerMiddleware 统一日志中间件，包含请求追踪ID和请求耗时记录
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成并设置 trace_id
		traceID := uuid.New().String()
		c.Set("trace_id", traceID)

		// 记录请求开始
		startTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		zap.L().Info("request started",
			zap.String("trace_id", traceID),
			zap.String("method", method),
			zap.String("path", path),
		)

		// 执行下一个中间件或路由处理器
		c.Next()

		// 请求结束后记录相关信息
		latency := time.Since(startTime).Milliseconds()
		statusCode := c.Writer.Status()

		zap.L().Info("request completed",
			zap.String("trace_id", traceID),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status_code", statusCode),
			zap.String("client_ip", c.ClientIP()),
			zap.String("cost_time", fmt.Sprintf("%d ms", latency)),
		)
	}
}

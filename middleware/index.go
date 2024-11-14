package middleware

import (
	"fmt"
	"lyp-go/logger"
	"lyp-go/model"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoadMidde(c *gin.Engine) {
	// 手动添加Logger中间件
	c.Use(zapLog())
	c.Use(respHeader())
	c.Use(customRecovery())
}

func safeHeader(ctx *gin.Context) {
	ctx.Header("X-Frame-Options", "DENY")
	ctx.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
	ctx.Header("X-XSS-Protection", "1; mode=block")
	ctx.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	ctx.Header("Referrer-Policy", "strict-origin")
	ctx.Header("X-Content-Type-Options", "nosniff")
	ctx.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
}

func respHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		safeHeader(ctx)
		ctx.Next()
	}
}

// 自定义的Recovery中间件
func customRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 获取堆栈信息
				stack := make([]byte, 1024*8)
				length := runtime.Stack(stack, false)

				// 打印堆栈信息（可选）
				logger.GetLogger().Error(fmt.Sprintf("Panic info is: [%v], stack is: [%s]\n", r, stack[:length]))

				// 如果是自定义异常，特殊处理
				if lErr, ok := r.(*model.LError); ok {
					// 返回JSON格式的错误响应
					ctx.JSON(http.StatusInternalServerError, model.LErr2Json(lErr))
				} else {
					// 返回JSON格式的错误响应
					ctx.JSON(http.StatusInternalServerError, model.Json(
						-1,
						fmt.Sprintf("服务器内部错误: %v", r),
						nil,
					))
				}

				ctx.Abort() // 中止请求处理
			}
		}()
		ctx.Next() // 执行下一个中间件或处理函数
	}
}

func zapLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if logger.GetLogger() != nil {
			start := time.Now()
			ctx.Next()
			log := logger.GetLogger()
			log.Info("handled request",
				zap.String("method", ctx.Request.Method),
				zap.String("path", ctx.Request.URL.Path),
				zap.Int("status", ctx.Writer.Status()),
				zap.Duration("duration", time.Since(start)),
			)
		}
	}
}

package middleware

import (
	"fmt"
	"lyp-go/logger"
	"lyp-go/resp"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoadMidde(c *gin.Engine) {
	c.Use(reqt2resp())
	c.Use(customRecovery())
}

func reqt2resp() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Header("Access-Control-Allow-Origin", "https://lyp.ink")
		c.Header("Vary", "Origin")

		if logger.GetLogger() != nil {
			c.Next()
			logger.GetLogger().Info("handled request",
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.Int("status", c.Writer.Status()),
				zap.Duration("duration", time.Since(start)),
			)
		} else {
			c.Next()
		}
	}
}

// 自定义的Recovery中间件
func customRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 获取堆栈信息
				stack := make([]byte, 1024*8)
				length := runtime.Stack(stack, false)

				// 如果是自定义异常，特殊处理
				if lErr, ok := r.(*resp.LError); ok {
					logger.Errorf("Panic info is: %s, stack is: \n%s", resp.Err2Str(lErr), stack[:length])
					// 返回JSON格式的错误响应
					c.JSON(http.StatusInternalServerError, lErr)
				} else {
					logger.Errorf("Panic info is: %s, stack is: \n%s", r, stack[:length])
					// 返回JSON格式的错误响应
					c.JSON(http.StatusInternalServerError, resp.Err(300, fmt.Sprintf("服务器内部错误: %+v", r), nil))
				}

				c.Abort() // 中止请求处理
			}
		}()
		c.Next() // 执行下一个中间件或处理函数
	}
}

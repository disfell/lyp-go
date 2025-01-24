package middleware

import (
	"fmt"
	"lyp-go/logger"
	"lyp-go/model"
	"lyp-go/output"
	"net/http"
	"runtime/debug"
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
		if gin.Mode() == gin.ReleaseMode {
			c.Header("X-Frame-Options", "DENY")
			c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
			c.Header("X-XSS-Protection", "1; mode=block")
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
			c.Header("X-Content-Type-Options", "nosniff")
			c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
			c.Header("Access-Control-Allow-Origin", "https://lyp.ink")
			c.Header("Vary", "Origin")
			c.Header("Referrer-Policy", "strict-origin")
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}

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
				// 如果是自定义异常，特殊处理
				if lErr, ok := r.(*output.LError); ok {
					logger.Errorf("Panic info is: %s, stack is: \n%s", output.Err2Str(lErr), string(debug.Stack()))
					// 返回JSON格式的错误响应
					c.JSON(http.StatusOK, lErr)
				} else {
					logger.Errorf("Panic info is: %s, stack is: \n%s", r, string(debug.Stack()))
					// 返回JSON格式的错误响应
					c.JSON(http.StatusOK, output.Err(model.ErrorCode, fmt.Sprintf("服务器内部错误: %+v", r), nil))
				}

				c.Abort() // 中止请求处理
			}
		}()
		c.Next() // 执行下一个中间件或处理函数

		// 处理非 panic 的常规错误
		if len(c.Errors) > 0 {
			lastError := c.Errors.Last().Err
			logger.Errorf("Error info is: %s, stack is: \n%s", lastError, string(debug.Stack()))
			c.JSON(http.StatusInternalServerError, output.Err(model.ErrorCode, fmt.Sprintf("服务器内部错误: %+v", lastError.Error()), nil))
			c.Abort()
		}
	}
}

func Cache1day() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 Cache-Control 头，使浏览器缓存该响应 1 天
		c.Writer.Header().Set("Cache-Control", "public, max-age=86400")
		c.Next()
	}
}

func Cache1min() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 Cache-Control 头，使浏览器缓存该响应 1 天
		c.Writer.Header().Set("Cache-Control", "public, max-age=60")
		c.Next()
	}
}

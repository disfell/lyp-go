package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/logger"
	"lyp-go/model"
	"lyp-go/output"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func UrlProxyHandler(c *gin.Context) {
	// 构造目标 URL
	targetURL, err := url.Parse(c.Param("target")[1:])
	logger.Debugf("proxy target path: %s", targetURL)

	if err != nil {
		c.JSON(http.StatusBadRequest, output.Err(model.ErrorCode, "Invalid target URL", nil))
		return
	}

	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// 修改请求的 Host 头为目标 URL 的 Host
	c.Request.URL.Host = targetURL.Host
	c.Request.URL.Scheme = targetURL.Scheme
	c.Request.Header.Set("X-Forwarded-Host", c.Request.Header.Get("Host"))
	c.Request.Host = targetURL.Host

	// 转发请求
	proxy.ServeHTTP(c.Writer, c.Request)
}

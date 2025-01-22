package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func UrlProxyHandler(c *gin.Context) {
	// 获取目标路径（如 /proxy/github.com -> github.com）
	targetPath := strings.TrimPrefix(c.Param("target"), "/")
	if targetPath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Target URL is required"})
		return
	}

	// 构造目标 URL
	targetURL, err := url.Parse("https://" + targetPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target URL"})
		return
	}

	// 发送 GET 请求到目标服务器
	resp, err := http.Get(targetURL.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reach target server"})
		return
	}
	defer resp.Body.Close()

	// 将目标服务器的响应体返回给客户端
	c.Status(resp.StatusCode)
	_, _ = io.Copy(c.Writer, resp.Body)
}

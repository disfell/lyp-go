package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"lyp-go/model"
	"lyp-go/output"
	"net/http"
	"net/url"
	"strings"
)

type ProxyController struct{}

func (pc *ProxyController) UrlProxyHandler(c *gin.Context) {
	// 获取目标路径（如 /proxy/github.com -> github.com）
	targetPath := strings.TrimPrefix(c.Param("target"), "/")
	if targetPath == "" {
		c.JSON(http.StatusOK, output.Err(model.ErrorCode, "Target URL is required", nil))
		return
	}

	// 构造目标 URL
	targetURL, err := url.Parse("https://" + targetPath)
	if err != nil {
		c.JSON(http.StatusOK, output.Err(model.ErrorCode, "Invalid target URL", nil))
		return
	}

	// 发送 GET 请求到目标服务器
	resp, err := http.Get(targetURL.String())
	if err != nil {
		c.JSON(http.StatusOK, output.Err(model.ErrorCode, "Failed to reach target server", nil))
		return
	}
	defer resp.Body.Close()

	// 将目标服务器的响应体返回给客户端
	c.Status(resp.StatusCode)
	_, _ = io.Copy(c.Writer, resp.Body)
}

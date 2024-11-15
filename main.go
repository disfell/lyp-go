package main

import (
	"lyp-go/logger"
	"lyp-go/middleware"
	"lyp-go/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 确保所有日志都写入
	defer logger.Sync()
	c := gin.New()
	// 注册中间件
	middleware.LoadMidde(c)
	// 注册路由
	router.InitRouter(c)
	// 启动服务
	http.ListenAndServe(":8180", c)
}

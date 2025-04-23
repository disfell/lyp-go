package main

import (
	"github.com/gin-gonic/gin"
	"lyp-go/config"
	"lyp-go/logger"
	"lyp-go/middleware"
	"lyp-go/router"
	"net/http"
)

var (
	app *gin.Engine
)

// @title           Gin Swagger 示例 API
// @version         1.0
// @description     基于 Gin 的 Swagger 接口文档示例
// @host            localhost:8180
// @BasePath        /
func main() {
	config.Take()
	// 确保所有日志都写入
	defer logger.Sync()
	app = gin.New()
	// 注册中间件
	middleware.LoadMidde(app)
	// 注册路由
	router.InitRouter(app)
	// 启动服务
	_ = http.ListenAndServe(":8180", app)
}

package main

import (
	"lyp-go/logger"
	"lyp-go/middleware"
	"lyp-go/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func main() {
	// 确保所有日志都写入
	defer logger.Sync()
	app = gin.New()
	// 注册中间件
	middleware.LoadMidde(app)
	// 注册路由
	router.InitRouter(app)
	// 启动服务
	err := http.ListenAndServe(":8180", app)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return
	}
}

package api

import (
	"github.com/gin-gonic/gin"
	"lyp-go/config"
	"lyp-go/cron"
	"lyp-go/logger"
	"lyp-go/middleware"
	"lyp-go/router"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

var (
	app *gin.Engine
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	config.Take()
	// 确保所有日志都写入
	defer logger.Sync()
	app = gin.New()
	// 注册中间件
	middleware.LoadMidde(app)
	// 注册路由
	router.InitRouter(app)
	// 初始化定时任务
	c := cron.Init()
	defer c.Stop()
}

package api

import (
	"github.com/gin-gonic/gin"
	"lyp-go/config"
	"lyp-go/logger"
	"lyp-go/middleware"
	"lyp-go/router"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

var (
	app *gin.Engine
)

func init() {
	// 设置中国时区
	_, _ = time.LoadLocation("Asia/Shanghai")
	// 生产模式
	gin.SetMode(gin.ReleaseMode)
	config.Take()
	// 确保所有日志都写入
	defer logger.Sync()
	app = gin.New()
	// 注册中间件
	middleware.LoadMidde(app)
	// 注册路由
	router.InitRouter(app)
}

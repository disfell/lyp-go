package api

import (
	"github.com/gin-gonic/gin"
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

func Start(c *gin.Engine) {
	// 确保所有日志都写入
	defer logger.Sync()
	c = gin.New()
	// 注册中间件
	middleware.LoadMidde(c)
	// 注册路由
	router.InitRouter(c)
	// 设置静态文件目录
	c.Static("/static", "./static")
	gin.SetMode(gin.ReleaseMode)
	logger.Infof("发布模式=%s", gin.Mode())
}

func init() {
	Start(app)
}

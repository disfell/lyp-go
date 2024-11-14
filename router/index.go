package router

import (
	"lyp-go/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(ctx *gin.Engine) {
	// 注册路径
	ctx.GET("/", handler.HelloHandler)
	ctx.GET("/testError", handler.TestError)
	ctx.POST("/testSqlite", handler.TestSqlite)
}

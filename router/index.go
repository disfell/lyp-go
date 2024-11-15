package router

import (
	"lyp-go/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Engine) {
	// 注册路径
	c.GET("/", handler.HelloHandler)
	c.POST("/testSqlite", handler.TestSqlite)
}

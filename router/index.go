package router

import (
	"lyp-go/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Engine) {
	// 注册路径
	c.GET("/api", handler.HelloHandler)
	c.POST("/api/testSqlite", handler.TestSqlite)
}

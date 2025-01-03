package router

import (
	"github.com/gin-gonic/gin"
	"lyp-go/handler"
)

func InitRouter(c *gin.Engine) {
	c.GET("/", handler.HelloHandler)
	c.POST("/", handler.HelloHandler)
	c.PUT("/", handler.HelloHandler)
	c.DELETE("/", handler.HelloHandler)

	// 可公开的路由
	v1 := c.Group("/api")
	{
		v1.GET("/public", handler.ApiHandler)
	}
}

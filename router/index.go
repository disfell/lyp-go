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
	api := c.Group("/api")
	{
		public := api.Group("/public")
		{
			public.GET("/steam", handler.SteamGamesHandler)
		}
	}

	// 自定义404页面
	c.NoRoute(handler.NotFound)
}

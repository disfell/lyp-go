package router

import (
	"github.com/gin-gonic/gin"
	"lyp-go/handler"
)

func InitRouter(c *gin.Engine) {
	c.GET("/", handler.HelloHandler)

	// 可公开的路由
	api := c.Group("/api")
	{
		public := api.Group("/public")
		{
			public.GET("/steam", handler.SteamHandler)
		}
	}

	c.NoRoute(handler.NotFoundHandler)
}

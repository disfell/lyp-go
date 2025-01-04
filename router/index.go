package router

import (
	"github.com/gin-gonic/gin"
	"lyp-go/handler"
	"lyp-go/middleware"
)

func InitRouter(c *gin.Engine) {
	// 可公开的路由
	api := c.Group("/api").Group("/public")
	{
		api.POST("/steam/games", middleware.Cache1day(), handler.SteamHandler)
		api.GET("/steam/status", middleware.Cache1min(), handler.SteamStatus)
	}

	c.NoRoute(handler.NotFoundHandler)
}

package router

import (
	"github.com/gin-gonic/gin"
	"lyp-go/handler"
)

func InitRouter(c *gin.Engine) {
	// 可公开的路由
	api := c.Group("/api").Group("/public")
	{
		api.POST("/steam/games", handler.SteamHandler)
		api.GET("/steam/status", handler.SteamStatus)

		api.GET("/github/trending", handler.GitHubTrendingHandler)

		api.GET("/notion/qrytable", handler.NotionDatabaseQryHandler)
	}

	proxy := c.Group("/proxy")
	{
		proxy.Any("/*target", handler.UrlProxyHandler)
	}

	c.NoRoute(handler.NotFoundHandler)
}

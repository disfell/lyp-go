package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "lyp-go/docs" // 导入生成的 docs 包
	"lyp-go/handler"
)

func InitRouter(c *gin.Engine) {

	steamCont := handler.SteamController{}
	githubCont := handler.GitHUbController{}
	notionCont := handler.NotionController{}
	proxyCont := handler.ProxyController{}
	errCont := handler.ErrController{}
	dtCont := handler.DateTimeController{}

	// 注册 Swagger 路由
	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 可公开的路由
	api := c.Group("/api").Group("/public")
	{
		api.POST("/steam/games", steamCont.SteamHandler)
		api.GET("/steam/status", steamCont.SteamStatus)
		api.GET("/github/trending", githubCont.GitHubTrendingHandler)
		api.POST("/notion/qryDatabase", notionCont.NotionDatabaseQryHandler)
		api.GET("/datetime/calculateDateDiff", dtCont.CalculateDateDiff)
	}

	proxy := c.Group("/proxy")
	{
		proxy.Any("/*target", proxyCont.UrlProxyHandler)
	}

	c.NoRoute(errCont.NotFoundHandler)
}

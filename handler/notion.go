package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/client"
	"lyp-go/output"
	"net/http"
)

func NotionDatabaseQryHandler(c *gin.Context) {
	c.JSON(http.StatusOK, output.Suc("", client.NotionTable()))
}

package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/client"
	"lyp-go/model"
	"lyp-go/output"
	"net/http"
)

func NotionDatabaseQryHandler(c *gin.Context) {
	databaseId := c.Query("databaseId")

	var reqBody = map[string]interface{}{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.JSON(http.StatusOK, output.Err(model.ErrorCode, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, output.Suc("", client.NotionDatabaseQry(databaseId, reqBody)))
}

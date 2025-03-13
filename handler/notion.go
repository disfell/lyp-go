package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/client"
	"lyp-go/model"
	"lyp-go/output"
	"net/http"
)

// NotionDatabaseQryHandler notion api docs: https://developers.notion.com/reference/intro
func NotionDatabaseQryHandler(c *gin.Context) {
	databaseId := c.Query("databaseId")
	filterProperties := c.Query("filter_properties")

	var reqBody = map[string]interface{}{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.JSON(http.StatusOK, output.Err(model.ErrorCode, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, output.Suc("", client.NotionDatabaseQry(databaseId, filterProperties, reqBody)))
}

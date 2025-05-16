package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/model"
	"lyp-go/output"
	"lyp-go/service"
	"net/http"
)

type NotionController struct{}

// NotionDatabaseQryHandler notion api docs: https://developers.notion.com/reference/intro
func (nc *NotionController) NotionDatabaseQryHandler(c *gin.Context) {
	databaseId := c.Query("databaseId")
	filterProperties := c.Query("filter_properties")

	var reqBody = map[string]interface{}{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		c.JSON(http.StatusOK, output.Err(model.ErrorCode, err.Error(), err))
		return
	}
	serv := service.NotionServ{}
	c.JSON(http.StatusOK, output.Suc("", serv.NotionDatabaseQry(databaseId, filterProperties, reqBody)))
}

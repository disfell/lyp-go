package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/output"
	"lyp-go/service"
	"net/http"
)

type DateTimeController struct {
}

var dts = service.DateTimeServ{}

func (dc *DateTimeController) CalculateDateDiff(c *gin.Context) {
	c.JSON(http.StatusOK, output.Suc("", dts.CalculateDateDiff(
		dts.ParseDate(c.Query("beginDate")),
		dts.ParseDate(c.Query("endDate")))))
}

package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/output"
	"net/http"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusOK, output.Suc("404", nil))
}

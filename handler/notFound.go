package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/resp"
	"net/http"
)

func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusOK, resp.Suc("404", nil))
}

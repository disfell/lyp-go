package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/output"
	"net/http"
)

type ErrController struct{}

func (ec *ErrController) NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusOK, output.Suc("404", nil))
}

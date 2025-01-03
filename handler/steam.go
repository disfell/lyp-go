package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/resp"
	"net/http"
)

func SteamHandler(c *gin.Context) {
	c.JSON(http.StatusOK, resp.Suc("steam index", nil))
}

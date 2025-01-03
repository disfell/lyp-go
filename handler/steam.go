package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/model"
	"net/http"
)

func SteamGamesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Suc("steam index"))
}

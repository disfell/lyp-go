package handler

import (
	"github.com/gin-gonic/gin"
	"lyp-go/model"
	"net/http"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusOK, model.Suc("无效路径"))
}

package handler

import (
	"errors"
	"lyp-go/db"
	"lyp-go/model"
	"lyp-go/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, resp.Suc("Hello Wrold", nil))
}

func TestSqlite(c *gin.Context) {
	action := c.Query("action")
	if action == "0" {
		var article model.Article
		err := db.DB.First(&article, c.Query("id")).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic("数据不存在，id=" + c.Query("id"))
		}
		c.JSON(http.StatusOK, resp.Suc("Hello Wrold", article))
	} else if action == "1" {
		var article model.Article
		if err := c.ShouldBindJSON(&article); err != nil {
			panic(err.Error())
		}
		db.DB.Create(&article)
		c.JSON(http.StatusOK, resp.Suc("Hello Wrold", article))
	} else if action == "2" {
		var article model.Article
		db.DB.Delete(&article, "id = ?", c.Query("id"))
		c.JSON(http.StatusOK, resp.Suc("删除成功", nil))
	} else {
		panic(resp.Err(300, "不支持的action类型: "+action, action))
	}
}

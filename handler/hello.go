package handler

import (
	"fmt"
	"lyp-go/db"
	"lyp-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.JsonSuc("你好!", nil))
}

func TestError(c *gin.Context) {
	action := c.Query("action")
	if action == "1" {
		panic(model.NewErr(300, "测试错误", map[string]string{"field1": "val1", "field2": "val2"}))
	} else {
		arr := []int{1, 2, 3}
		fmt.Println(arr[5])
	}
}

func TestSqlite(c *gin.Context) {
	action := c.Query("action")
	if action == "0" {
		var article model.Article
		db.DB.First(&article, c.Query("id"))
		c.JSON(http.StatusOK, model.JsonSuc("suc!", article))
	}
	if action == "1" {
		var article model.Article
		if err := c.ShouldBindJSON(&article); err != nil {
			panic(err.Error())
		}
		db.DB.Create(&article)
		fmt.Println(article)
		c.JSON(http.StatusOK, model.JsonSuc("suc!", article))
	}
}

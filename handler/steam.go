package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lyp-go/client"
	"lyp-go/lhttp"
	"lyp-go/logger"
	"lyp-go/model"
	"lyp-go/output"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type SteamController struct{}

func (sc *SteamController) SteamHandler(c *gin.Context) {
	// 先查原有的数据
	data := client.SupaGet("games", &map[string]string{"select": "*"})
	location, _ := time.LoadLocation("Asia/Shanghai")
	chinaTime := time.Now().In(location)

	if data != nil && len(data) > 0 {
		createdAtStr := data[0]["created_at"].(string)
		parsedTime, _ := time.Parse(time.RFC3339, createdAtStr)

		diff := chinaTime.Sub(parsedTime)
		if diff.Hours() < 1 {
			c.JSON(http.StatusOK, output.Suc("", data))
			return
		}
	}

	// 创建查询参数
	params := &url.Values{}
	params.Add("key", viper.GetString("steam.token"))
	params.Add("steamid", viper.GetString("steam.id"))
	recentUrl := viper.GetString("steam.recentUrl")
	steamRet := lhttp.Get[map[string]interface{}](recentUrl, params, nil)

	games := steamRet["response"].(map[string]interface{})["games"].([]interface{})

	logger.Debugf("steam resp: %v", games)

	var collection []map[string]interface{}

	if games != nil {
		for _, game := range games {
			v := game.(map[string]interface{})
			collection = append(collection, map[string]interface{}{
				"name":       v["name"].(string),
				"name_cn":    model.SteamDict[strconv.Itoa(int(v["appid"].(float64)))],
				"game_id":    v["appid"].(float64),
				"play_time":  v["playtime_forever"].(float64),
				"original":   v,
				"updated_at": time.Now(),
				"created_at": time.Now(),
			})
		}
		client.SupaDelete("games", &map[string]string{"id": "gt.0"})
		client.SupaInsert("games", &collection)
	}

	logger.Debugf("steam build collection: %v", collection)

	c.JSON(http.StatusOK, output.Suc("", collection))
}

func (sc *SteamController) SteamStatus(c *gin.Context) {
	// 创建查询参数
	params := &url.Values{}
	params.Add("key", viper.GetString("steam.token"))
	params.Add("steamids", viper.GetString("steam.id"))
	recentUrl := viper.GetString("steam.userStatus")
	steamRet := lhttp.Get[map[string]interface{}](recentUrl, params, nil)
	c.JSON(http.StatusOK, output.Suc("", steamRet))
}

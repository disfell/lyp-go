package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lyp-go/client"
	http2 "lyp-go/http"
	"lyp-go/logger"
	"lyp-go/model"
	"lyp-go/output"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func SteamHandler(c *gin.Context) {
	// 创建查询参数
	params := &url.Values{}
	params.Add("key", viper.GetString("steam.token"))
	params.Add("steamid", viper.GetString("steam.id"))
	recentUrl := viper.GetString("steam.recentUrl")
	steamRet := http2.GetMap(recentUrl, params, nil)

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

	c.JSON(http.StatusOK, output.Suc("", collection))
}

func SteamStatus(c *gin.Context) {
	// 创建查询参数
	params := &url.Values{}
	params.Add("key", viper.GetString("steam.token"))
	params.Add("steamids", viper.GetString("steam.id"))
	recentUrl := viper.GetString("steam.userStatus")
	steamRet := http2.GetMap(recentUrl, params, nil)
	c.JSON(http.StatusOK, output.Suc("", steamRet))
}

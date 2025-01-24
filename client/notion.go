package client

import (
	"github.com/spf13/viper"
	"lyp-go/lhttp"
	"lyp-go/logger"
	"strings"
)

func NotionHeader() map[string]string {
	token := viper.GetString("notion.token")
	return map[string]string{"Authorization": "Bearer " + token, "Content-Type": "application/json", "Notion-Version": "2022-06-28"}
}

func NotionTable() []map[string]interface{} {
	api := viper.GetString("notion.api")
	database := viper.GetString("notion.database.qry")
	databaseId := viper.GetString("notion.database.blogCollection")
	database = strings.Replace(database, "${database_id}", databaseId, -1)
	url := api + database
	resp := lhttp.Post[map[string]interface{}](url, nil, nil, NotionHeader())
	logger.Debugf("Notion response: %v", resp)

	NotionRespCheck(resp)

	var list []map[string]interface{}
	results := resp["results"].([]interface{})

	for _, item := range results {
		properties := item.(map[string]interface{})["properties"].(map[string]interface{})
		url := properties["url"].(map[string]interface{})["url"].(string)
		name := properties["名称"].(map[string]interface{})["title"].([]interface{})[0].(map[string]interface{})["plain_text"].(string)

		list = append(list, map[string]interface{}{
			"url":  url,
			"name": name,
		})
	}
	return list
}

func NotionRespCheck(respBody map[string]interface{}) {
	//if "0" == respBody["code"] {
	//	return
	//}
	//
	//panic(output.Err(model.ErrorCode, respBody["msg"].(string), respBody))
}

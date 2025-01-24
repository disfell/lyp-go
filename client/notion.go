package client

import (
	"github.com/spf13/viper"
	"lyp-go/lhttp"
	"lyp-go/model"
	"lyp-go/output"
	"strings"
)

func NotionHeader() map[string]string {
	token := viper.GetString("notion.token")
	return map[string]string{"Authorization": "Bearer " + token, "Content-Type": "application/json", "Notion-Version": "2022-06-28"}
}

func NotionDatabaseQry(databaseId string, reqBody map[string]interface{}) []map[string]interface{} {
	api := viper.GetString("notion.api")
	qryUrl := viper.GetString("notion.database.qry")
	database := strings.Replace(qryUrl, "${database_id}", databaseId, -1)
	url := api + database

	resp := lhttp.Post[map[string]interface{}](url, nil, reqBody, NotionHeader())

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
	if "error" != respBody["object"] {
		return
	}

	panic(output.Err(model.ErrorCode, respBody["message"].(string), respBody))
}

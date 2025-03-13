package client

import (
	"github.com/spf13/viper"
	"lyp-go/lhttp"
	"lyp-go/model"
	"lyp-go/output"
	"net/url"
	"strings"
	"time"
)

func NotionHeader() map[string]string {
	token := viper.GetString("notion.token")
	return map[string]string{"Authorization": "Bearer " + token, "Content-Type": "application/json", "Notion-Version": "2022-06-28"}
}

func NotionDatabaseQry(databaseId string, filterProperties string, reqBody map[string]interface{}) any {
	api := viper.GetString("notion.api")
	qryUrl := viper.GetString("notion.database.qry")
	database := strings.Replace(qryUrl, "${database_id}", databaseId, -1)
	databaseUrl := api + database

	params := &url.Values{}
	if filterProperties != "" {
		params.Add("filter_properties", filterProperties)
	}
	resp := lhttp.Post[map[string]interface{}](databaseUrl, params, reqBody, NotionHeader())

	notionRespCheck(resp)

	results := resp["results"].([]interface{})

	if len(results) <= 0 {
		return make([]map[string]interface{}, 0)
	}

	if "list" == resp["object"] {
		return seekList(resp["results"].([]interface{}))
	} else {
		panic(output.Err(model.ErrorCode, "未适配的数据格式", resp["object"]))
	}
}

func seekList(objs []interface{}) interface{} {

	items := make([]interface{}, len(objs))

	for idx, item := range objs {
		properties := item.(map[string]interface{})["properties"].(map[string]interface{})

		oneObj := make(map[string]interface{})
		for key, value := range properties {
			type_ := value.(map[string]interface{})["type"].(string)
			oneObj[key] = matchData(type_, value)
		}

		items[idx] = oneObj
	}
	return items
}

func matchData(type_ string, value interface{}) interface{} {
	if "created_time" == type_ {
		return getDateTime("created_time", value.(map[string]interface{}))
	}
	if "last_edited_time" == type_ {
		return getDateTime("last_edited_time", value.(map[string]interface{}))
	}
	if "url" == type_ {
		return getString("url", value.(map[string]interface{}))
	}
	if "multi_select" == type_ {
		return getMultiSelect(value.(map[string]interface{}))
	}
	if "title" == type_ {
		return getTitle(value.(map[string]interface{}))
	}
	panic(output.Err(model.ErrorCode, "未适配数据处理格式: "+type_, type_))

}

func getTitle(item map[string]interface{}) interface{} {
	return getString("plain_text", item["title"].([]interface{})[0].(map[string]interface{}))
}

func notionRespCheck(respBody map[string]interface{}) {
	if "error" != respBody["object"] {
		return
	}

	panic(output.Err(model.ErrorCode, respBody["message"].(string), respBody))
}

func getDateTime(key string, item map[string]interface{}) time.Time {
	t, err := time.Parse(model.TimeFormatISO860_UTC_WITH_MS, item[key].(string))
	if err != nil {
		panic(err)
	}
	return t
}

func getString(key string, item map[string]interface{}) string {
	return item[key].(string)
}

func getMultiSelect(item map[string]interface{}) []string {
	multiSelect := item["multi_select"].([]interface{})
	selectItems := make([]string, len(multiSelect))

	for idx, selectItem := range multiSelect {
		selectItems[idx] = selectItem.(map[string]interface{})["name"].(string)
	}

	return selectItems
}

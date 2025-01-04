package client

import (
	"github.com/spf13/viper"
	http2 "lyp-go/http"
	"lyp-go/logger"
	"lyp-go/model"
	"lyp-go/output"
	"net/url"
)

func SupaDelete(table string, cond *map[string]string) map[string]interface{} {
	supabaseUrl := viper.GetString("supabase.url") + "/rest/v1/" + table
	headers := map[string]string{
		"apikey":        viper.GetString("supabase.key"),
		"Authorization": "Bearer " + viper.GetString("supabase.key"),
	}
	params := &url.Values{}
	if cond != nil {
		for key, value := range *cond {
			params.Add(key, value)
		}
	}

	ret := http2.DeleteMap(supabaseUrl, params, nil, headers)
	logger.Debugf("del data resp: %v", ret)

	if nil != ret["code"] {
		panic(output.Err(model.ErrorCode, ret["code"].(string)+": "+ret["message"].(string), ret))
	}
	return ret
}

func SupaInsert(table string, data *[]map[string]interface{}) map[string]interface{} {
	supabaseUrl := viper.GetString("supabase.url") + "/rest/v1/" + table
	headers := map[string]string{
		"Prefer":        "return=minimal",
		"apikey":        viper.GetString("supabase.key"),
		"Authorization": "Bearer " + viper.GetString("supabase.key"),
	}
	ret := http2.PostMap(supabaseUrl, nil, data, headers)

	logger.Debugf("insert data resp: %v", ret)

	if nil != ret["code"] {
		panic(output.Err(model.ErrorCode, ret["code"].(string)+": "+ret["message"].(string), ret))
	}
	return ret
}

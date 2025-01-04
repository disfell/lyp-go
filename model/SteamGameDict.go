package model

import (
	"encoding/json"
	"io"
	"os"
)

// SteamDict 初始化缓存
var SteamDict = make(map[string]interface{})

func init() {
	fileContent, err := os.Open("config/steam.dict.json")
	if err != nil {
		return
	}
	defer func(fileContent *os.File) {
		err := fileContent.Close()
		if err != nil {
			return
		}
	}(fileContent)

	byteResult, _ := io.ReadAll(fileContent)

	err = json.Unmarshal(byteResult, &SteamDict)
	if err != nil {
		return
	}
}

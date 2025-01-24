package lhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"lyp-go/logger"
	"net/http"
	"net/url"
)

func Get[T any](url string, params *url.Values, headers map[string]string) T {
	recentUrl := url
	if params != nil && len(params.Encode()) > 0 {
		recentUrl = url + "?" + params.Encode()
	}
	req, err := http.NewRequest("GET", recentUrl, nil)
	if err != nil {
		panic(err.Error())
	}
	// 设置请求头
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	// 发送请求
	client := &http.Client{}
	logger.Debugf("get url: %s, header: %s", recentUrl, headers)
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)

	return retT[T](body)
}

func Post[T any](url string, params *url.Values, requestBody interface{}, headers map[string]string) T {
	recentUrl := url
	// 将请求参数附加到URL
	if params != nil {
		recentUrl = recentUrl + "?" + params.Encode()
	}

	logger.Debugf("req url: %v", recentUrl)
	logger.Debugf("req body: %v", requestBody)
	logger.Debugf("req headers: %v", headers)
	// 将请求体编码为JSON
	var body io.Reader
	if requestBody != nil {
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			panic(err.Error())
		}
		body = bytes.NewBuffer(jsonBody)
	}

	// 创建POST请求
	req, err := http.NewRequest("POST", recentUrl, body)
	if err != nil {
		panic(err.Error())
	}

	// 设置请求头
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	if requestBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(resp.Body)

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	return retT[T](respBody)
}

func Delete[T any](url string, params *url.Values, requestBody interface{}, headers map[string]string) T {
	recentUrl := url
	// 将请求参数附加到URL
	if params != nil {
		recentUrl = recentUrl + "?" + params.Encode()
	}

	// 将请求体编码为JSON
	var body io.Reader
	if requestBody != nil {
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			panic(err.Error())
		}
		body = bytes.NewBuffer(jsonBody)
	}

	// 创建POST请求
	req, err := http.NewRequest("DELETE", recentUrl, body)
	if err != nil {
		panic(err.Error())
	}

	// 设置请求头
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	if requestBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 发送请求
	client := &http.Client{}
	logger.Debugf("delete url: %s, header: %s, body: %s", recentUrl, headers, body)
	resp, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err.Error())
		}
	}(resp.Body)

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	return retT[T](respBody)
}

func retT[T any](data []byte) T {
	var empty T
	if "" == string(data) {
		return empty
	}
	logger.Debugf("get data resp: %v", string(data))

	var ret T
	err := json.Unmarshal(data, &ret)
	if err != nil {
		panic(err.Error())
	}
	return ret
}

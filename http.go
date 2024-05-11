package game_center_client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func postRequest(url string, postData map[string]interface{}) (res CommonResp, err error) {
	// 将 POST 数据转换为 JSON 格式
	postBody, err := json.Marshal(postData)
	if err != nil {
		return res, err
	}

	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 创建一个 POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		return res, err
	}

	// 设置默认的请求头
	req.Header.Set("Content-Type", "application/json")

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	// 打印响应体
	//fmt.Println(string(body))

	// 将响应体转换为 CommonResp 结构体
	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

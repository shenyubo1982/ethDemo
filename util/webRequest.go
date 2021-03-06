package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func WebGetRequest(requestUrl string) (body []byte, statusCode int, err error) {
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	//if resp.Body == nil {
	//	return nil, 0, err
	//}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, resp.StatusCode, err2
	}
	if resp.StatusCode == 200 {
		return body, resp.StatusCode, err
	} else {
		return nil, resp.StatusCode, err
	}
}

type GetInfoJson struct {
	Data Data `json:"data"`
}

type Data struct {
	AddressCount string `json:"addressCount"`
	TransCount   string `json:"transCount"`
	BlockNumber  string `json:"blockNumber"`
	NodeNumber   int    `json:"nodeNumber"`
	Ts           int64  `json:"ts"`
}

// SearchBody2Json TODO 如何动态转换成json返回
func SearchBody2Json(body []byte, getInfoJson GetInfoJson, searchKey string) string {
	tokenGet := string(body)                                   // 请求结果string格式
	infoJson := json.Unmarshal([]byte(tokenGet), &getInfoJson) // 将string 格式转成json格式
	if infoJson == nil {
		log.Printf("Error Json : %v\n", infoJson) // 错误写进日志文件
	}
	t := reflect.TypeOf(getInfoJson.Data)
	v := reflect.ValueOf(getInfoJson.Data)
	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		value := v.Field(k)
		if searchKey == name {
			//匹配成功
			return value.String()
		}
	}
	//没有匹配
	return ""
}

package util

import (
	"testing"
)

//
//  TestWebGetRequest
//  @Description:
//  @param t
//
func TestWebGetRequest(t *testing.T) {
	t.Run("GetWebRequest|Correct", func(t *testing.T) {
		t.Helper()
		requestUrl := "http://testing-metain-js-official-node-service.jiansutech.com/api/mt/getInfo"
		body, statusCode, err := WebGetRequest(requestUrl)
		if err != nil {
			t.Errorf("Erro is Exception : %v\n expeted is nil.\n", err)
		}
		if statusCode != 200 {
			t.Errorf("statudcode is : %v\n expeted 200\n", statusCode)
		}
		if body == nil {
			t.Errorf("request body is : %v\n expeted not nil.\n", body)
		}

		getInfoJson := GetInfoJson{}
		webBN := SearchBody2Json(body, getInfoJson, "BlockNumber")
		//转换body to json
		if webBN == "" {
			t.Errorf("没有找到数据。")
		}
	})

	t.Run("GetWebRequest|NotFound", func(t *testing.T) {
		t.Helper()
		requestUrl := "http://testing-metain-js-official-node-service.jiansutech.com/api/mt/getInfo"
		body, statusCode, err := WebGetRequest(requestUrl)
		if err != nil {
			t.Errorf("Erro is Exception : %v\n expeted is nil.\n", err)
		}
		if statusCode != 200 {
			t.Errorf("statudcode is : %v\n expeted 200\n", statusCode)
		}
		if body == nil {
			t.Errorf("request body is : %v\n expeted not nil.\n", body)
		}

		getInfoJson := GetInfoJson{}
		webBN := SearchBody2Json(body, getInfoJson, "ErrorBlockNumber")
		//转换body to json
		if webBN != "" {
			t.Errorf("应该没有找到数据")
		}
	})

	t.Run("GetWebRequest|httpsNotCorrect", func(t *testing.T) {
		t.Helper()
		requestUrl := "https://testing-metain-js-official-node-service.jiansutech.com/api/mt/getInfo"
		_, _, err := WebGetRequest(requestUrl)
		if err == nil {
			t.Errorf("Erro is Exception : %v\n expeted is nil.\n", err)
		}
		//if statusCode != 200 {
		//	t.Errorf("statudcode is : %v\n expeted 200\n", statusCode)
		//}
		//if body == nil {
		//	t.Errorf("request body is : %v\n expeted not nil.\n", body)
		//}
	})
}

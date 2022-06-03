package util

import (
	"log"
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
		} else {
			log.Println(string(body))
		}

		//转换body to json
		//ConvertBody2Json(body)
	})
}

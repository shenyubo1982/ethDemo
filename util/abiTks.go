package util

import (
	"fmt"
	"io/ioutil"
)

// ReadAbi
//  @Description: abi文件读取
//  @param filePath abi File path & File name, ex. folder/filename.abi
//  @return string abi file body.
//
func ReadAbi(filePath string) string {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}

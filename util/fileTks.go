package util

import (
	"os"
	"path/filepath"
	"regexp"
)

func GetRunPath(dirName string, fileName string) string {
	pwd, _ := os.Getwd() // 获取到当前目录，相当于python里的os.getcwd()
	//fmt.Println("当前pwd路径为:", pwd)
	fileFullPath := filepath.Join(pwd, dirName, fileName)
	return fileFullPath
}

func IsDigit(checkStr string) bool {
	pattern := "\\d+" //反斜杠要转义
	//str := "124534"
	result, _ := regexp.MatchString(pattern, checkStr)
	return result
}

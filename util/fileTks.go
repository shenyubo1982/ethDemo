package util

import (
	"bufio"
	"fmt"
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

// ReadFileAllContent ReadFileContent 读取文件方法
func ReadFileAllContent(dirName string, fileName string) []string {
	var fileContent []string

	fileFullPath := dirName + fileName
	fmt.Println(fileFullPath)

	readFile, err := os.Open(fileFullPath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileContent = append(fileContent, fileScanner.Text())
		//fmt.Println(fileScanner.Text())
	}

	_ = readFile.Close()

	return fileContent

	//f, err := ioutil.ReadFile(fileFullPath)
	//if err != nil {
	//	fmt.Println("read fail", err)
	//}
	//return string(f)
}

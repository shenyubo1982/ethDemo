package util

import (
	"fmt"
	"io/ioutil"
	"os"
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

// Exists
//  Exists  file or dir is existed
//  @Description:
//  @param path
//  @return bool
//
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// GetFilesInDir
//  @Description: 读取XX的目录，将所有文件的名称返回
//  @param keyFileDir 相对路径 ../mt
//  @return []string  目录下所有的文件名称
//
func GetFilesInDir(dirName string) []string {
	if dirName == "" {
		panic("Dir is empty.")
	}

	if !IsDir(dirName) {
		panic("Dir is not correct.")
	}

	filesName := make([]string, 0, 10)
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		fileFullName := dirName + "/" + f.Name()
		if IsFile(fileFullName) {
			filesName = append(filesName, f.Name())
		}
	}
	return filesName
}
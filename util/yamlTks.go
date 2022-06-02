package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

const YamlPath = "../config/"

////yaml文件内容影射的结构体，注意结构体成员要大写开头
//type Stu struct {
//	Name  string `yaml:"Name"`
//	Age   string `yaml:"Age"`
//	Sex   string `yaml:"Sex"`
//	Class string `yaml:"class"`
//}

// YamlContent
//  @Description: 与yaml配置文件中的key匹配
//
type YamlContent struct {
	ChainUrl           string `yaml:"chainUrl"`
	ContractAddressHex string `yaml:"contractAddressHex"`
	AdminAddressHex    string `yaml:"adminAddressHex"`
	AdminPrivateKeyHex string `yaml:"adminPrivateKeyHex"`
	ChainID            string `yaml:"chainID"`
	Abi                string `yaml:"abi"`
	GasLimit           string `yaml:"gasLimit"`
	GasPrice           string `yaml:"gasPrice"`
}

// NewYamlContent
//  @Description: 构造器
//  @param filePath
//  @return *YamlContent
//
func NewYamlContent(filePath string) *YamlContent {
	yc := YamlContent{}
	//读取yaml文件到缓存中
	config, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	//yaml文件内容影射到结构体中
	err2 := yaml.Unmarshal([]byte(config), &yc)
	if err2 != nil {
		log.Fatalf("cannot unmarshal data: %v", err2)
	}
	return &yc
}

func (YamlContent *YamlContent) getChainUrl() string {
	return YamlContent.ChainUrl
}

type ChainTestYaml struct {
	filePath   string
	yamContent YamlContent
}

// Load
//  @Description: 装载不同测试环境的链相关配置信息(统一模板）
//  @param yamlFile 配置文件不需要路径
//  @return *YamlContent
//
func Load(yamlFileName string) *YamlContent {
	yamlFile := YamlPath + yamlFileName
	return NewYamlContent(yamlFile)
}

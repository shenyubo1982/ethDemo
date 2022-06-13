package util

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const YamlPath = "../config/"

// YamlContent
//  @Description: 与yaml配置文件中的key匹配
//  yaml文件内容影射的结构体，注意结构体成员要大写开头
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

type ChainTestYaml struct {
	filePath    string
	YamlContent YamlContent
}

// NewChainTestYaml Load NewChainConfig
//  @Description: 构造器
//  @param filePath
//  @return *YamlContent
//
//func (self *ChainConfig)NewChainConfig(filePath string) *YamlContent {
func NewChainTestYaml(fileName string) *ChainTestYaml {
	instance := new(ChainTestYaml)
	instance.filePath = YamlPath + fileName
	yc := YamlContent{}
	//读取yaml文件到缓存中
	config, err := ioutil.ReadFile(instance.filePath)
	if err != nil {
		return nil
	}
	//yaml文件内容影射到结构体中
	err2 := yaml.Unmarshal([]byte(config), &yc)
	if err2 != nil {
		return nil
	}
	instance.YamlContent = yc
	return instance
}

func (cc *ChainTestYaml) GetYamlContent() *YamlContent {
	return &cc.YamlContent
}

package util

import (
	"testing"
)

const ConfigFile = "metaTest.yaml"

//
//  TestLaunch
//  @Description:
//  @param t
//
func TestLoad(t *testing.T) {
	t.Run("LoadConfigFile-Success", func(t *testing.T) {
		t.Helper()
		myChainConfig := NewChainTestYaml(ConfigFile)
		if myChainConfig.YamlContent.ChainUrl != "http://172.17.4.13:7755" {
			t.Errorf("Value is not correct! ")
		}
		myChainConfig2 := myChainConfig.GetYamlContent()
		if myChainConfig2.ChainUrl != "http://172.17.4.13:7755" {
			t.Errorf("Value is not correct! ")
		}
	})

	t.Run("LoadConfigFile-Failed", func(t *testing.T) {
		t.Helper()
		myChainConfig := NewChainTestYaml("ErrorFileName")
		if myChainConfig != nil {
			t.Errorf("应该加载失败，返回nil")
		}

	})

}

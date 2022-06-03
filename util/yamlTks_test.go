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
	t.Run("LoadConfigFile", func(t *testing.T) {
		t.Helper()
		myChainConfig := NewChainTestYaml(ConfigFile)
		if myChainConfig.YamlContent.ChainUrl != "http://172.17.4.13:7755" {
			t.Fatal("Value is not correct! ")
		}
	})
}

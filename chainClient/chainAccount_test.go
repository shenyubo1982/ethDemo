package chainClient

import (
	"log"
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
		keyDir := "../keys/mt"
		iWantCnt := 15
		cas := LoadChainAccount(iWantCnt, keyDir)
		if cas.cnt != iWantCnt {
			t.Errorf("生成Account 错误！应该是%v ，实际是 %v", iWantCnt, cas.cnt)
		}

		for _, ca := range cas.accounts {
			address := ca.address
			priKey := ca.priKey
			keyFile := ca.keyFile
			log.Printf("File : %v \nAddress is %v \n privateKey is %v\n", keyFile, address, priKey)
			log.Printf("--------------------------------")
		}
	})
}

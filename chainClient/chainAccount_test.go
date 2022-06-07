package chainClient

import (
	"fmt"
	"log"
	"testing"
)

const ConfigFile = "metaTest.yaml"

//
//  TestLaunch
//  @Description:
//  @param t
//
func TestLoadChainAccount(t *testing.T) {
	t.Run("LoadChainAccount", func(t *testing.T) {
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
			//keyFile := ca.keyFile
			log.Printf("Address is %v \n privateKey is %v\n", address, priKey)
			log.Printf("--------------------------------")
		}
	})
}

func TestAddAccount(t *testing.T) {
	keyDir := "../keys/mt"

	t.Run("Add 1 Account", func(t *testing.T) {
		t.Helper()
		keyfileNames := []string{"UTC--2022-06-05T09-29-29.614223000Z--b319c54942ab601d89cab2654afaf9496ddaaa7d"}
		cas := new(chainAccounts)
		cas.AddAccount(keyDir, keyfileNames[0])
		if cas.Cnt() != 1 {
			t.Errorf("应该加载%v个Account,实际是%v", len(keyfileNames), cas.Cnt())
		}
		//显示加载账号的信息
		for _, account := range cas.accounts {
			log.Printf("priKeyHex %v\n", account.PriKeyHex())
			log.Printf("PubkeyHex %v\n", account.PubKeyHex())
			log.Printf("AddressHex %v\n", account.AddressHex())
		}
	})

	t.Run("Add 2 Account", func(t *testing.T) {
		t.Helper()
		keyfileNames := []string{"UTC--2022-06-05T09-29-29.614223000Z--b319c54942ab601d89cab2654afaf9496ddaaa7d",
			"UTC--2022-06-05T09-29-30.994525000Z--bbf1e2529432fa7735a573ffbbe0f698dd527062"}
		cas := new(chainAccounts)
		for i := 0; i < len(keyfileNames); i++ {
			cas.AddAccount(keyDir, keyfileNames[i])
		}
		if cas.Cnt() != len(keyfileNames) {
			t.Errorf("应该加载%v个Account,实际是%v", len(keyfileNames), cas.Cnt())
		}
	})
}

func TestGetPrivateKey(t *testing.T) {
	keyDir := "../keys/mt"
	keyfileName := "UTC--2022-06-05T09-29-29.614223000Z--b319c54942ab601d89cab2654afaf9496ddaaa7d"
	storeFile := keyDir + PathSymbol + keyfileName
	pwd := accountPwd

	t.Run("GetPrivateKey", func(t *testing.T) {
		t.Helper()
		privateKey := GetPrivateKey(storeFile, pwd)
		fmt.Println(privateKey)
	})

	t.Run("GetPrivateKeyHex", func(t *testing.T) {
		t.Helper()
		privateKey := GetPrivateKeyHex(storeFile, pwd)
		fmt.Println(privateKey)
	})
}

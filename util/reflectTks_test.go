package util

import (
	"context"
	"ethDemo/chainClient"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"log"
	"testing"
)

func Test_reflectDemo(t *testing.T) {

	u := User{Id: 1001, Name: "aaa", Address: Address{Add: "ccccccccc", Res: 12}}

	t.Run("Test_reflectDemo", func(t *testing.T) {
		reflectDemo(u)
	})
}

func Test_reflectRecipe(t *testing.T) {
	//连接区块链
	myChainConfig := NewChainTestYaml(ConfigFile)
	myChainClient := chainClient.Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		log.Printf("Can't get Client")
	}

	txHex := "0x006a8884e8a2314a5bb2f39962a0235f63a304346b4a5067221c5d4fe9253ba1"
	txHash, err2 := chainhash.NewHashFromStr(txHex)

	if err2 != nil {
		return
	}
	receipt, err := myChainClient.Client().TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	t.Run("reflectDemo", func(t *testing.T) {
		reflectReceipt(u)
	})
}

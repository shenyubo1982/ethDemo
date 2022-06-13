package util

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	client, err := ethclient.Dial("http://172.17.4.13:7755")
	if err != nil {
		log.Fatal(err)
	}

	t.Run("reflectRecipe-don't found", func(t *testing.T) {
		txHex := "0x006a8884e8a2314a5bb2f39962a0235f63a304346b4a5067221c5d4fe9253ba1"
		txHash := common.HexToHash(txHex)

		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			log.Fatal(err)
		}
		ReflectReceipt(*receipt, "", false)
	})

	t.Run("reflectRecipe-1lv-found status.", func(t *testing.T) {
		txHex := "0x006a8884e8a2314a5bb2f39962a0235f63a304346b4a5067221c5d4fe9253ba1"
		txHash := common.HexToHash(txHex)

		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			log.Fatal(err)
		}
		status := ReflectReceipt(*receipt, "Status", true)
		if status == nil {
			t.Errorf("没有找到")
		}
	})

	t.Run("reflectRecipe-1lv-notFound status.", func(t *testing.T) {
		txHex := "0x006a8884e8a2314a5bb2f39962a0235f63a304346b4a5067221c5d4fe9253ba1"
		txHash := common.HexToHash(txHex)

		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			log.Fatal(err)
		}
		status := ReflectReceipt(*receipt, "StatusAAAA", false)
		if status != nil {
			t.Errorf("没有找到，应该返回nil")
		}
	})

	t.Run("reflectRecipe-txHex is not correct.", func(t *testing.T) {
		txHex := "0xaaaaaaaaaa"
		txHash := common.HexToHash(txHex)

		_, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			t.Errorf("应该报错，txHex是错误的。")
		}
	})
}

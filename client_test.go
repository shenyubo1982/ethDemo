package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"reflect"
	"testing"
)

//
//  TestLaunch
//  @Description: 启动函数测试
//  @param t
//
func TestLaunch(t *testing.T) {
	t.Run("Launch", func(t *testing.T) {
		t.Helper()
		chainUrl := "http://172.17.4.13:7755"
		ansClient, err := launch(chainUrl)
		if ansClient == nil || err != nil {
			t.Fatal("Can't get Client")
		}
	})
}

func TestDo(t *testing.T) {
	t.Run("Do", func(t *testing.T) {
		t.Helper()
		//client, err := ethclient.Dial("ADD_YOUR_ETHEREUM_NODE_URL")
		//client, err := ethclient.Dial("https://cloudflare-eth.com")
		// 教程中使用了cloudflare-ech.com
		// 也可以使用自己的eth节点，这里我们尝试使用infura的免费服务。infura 需要新建项目，并且配置相应的测试网络和网络节点url。
		//https://infura.io/dashboard/ethereum/3f97ae7214cc4e2794bee5bdc3bd6b95/settings
		// Details > KEYS > ENDPOINTS (KOVAN) > https://kovan.infura.io/v3/<projectkey>
		// kovan 测试网络中的测试币可以通过水龙头领取
		infuraKovanUrl := "https://kovan.infura.io/v3/3f97ae7214cc4e2794bee5bdc3bd6b95"
		client, err := ethclient.Dial(infuraKovanUrl)
		//client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/3f97ae7214cc4e2794bee5bdc3bd6b95")
		fmt.Printf("Type %v is %v \n", reflect.TypeOf(client), client)
		if err != nil {
			log.Fatalf("Oops! There was a problem %v", err)
			os.Exit(100)
		} else {
			fmt.Println("Success! you are connected to the Ethereum Network")
		}
		myAddress := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
		kovanTestEthAddress := "0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c"
		myAddressEthValue := getBalanceFromAddress(*client, myAddress)
		KovanTestEthValue := getBalanceFromAddress(*client, kovanTestEthAddress)
		fmt.Printf("myAddress is %v,it's Eth Value in Kovan Testnet is %v\n", myAddress, myAddressEthValue)
		fmt.Printf("myAddress is %v,it's Eth Value in Kovan Testnet is %v\n", kovanTestEthAddress, KovanTestEthValue)

		// 根据账户地址和区块高度查询，区块交易的金额。 这个功能需要https://infura.io/dashboard 中购买 archive Data 功能才能调用api
		//KovanTestEthBlockNums := []int64{31866217, 31866202}
		//KovanTestEthValueAtBlock := getBalanceFromBlockNum(*client, kovanTestEthAddress, KovanTestEthBlockNums[0])
		//fmt.Printf("myAddress is %v,it's in %v BlockNumber ,and Kovan Testnet is %v\n", kovanTestEthAddress, KovanTestEthBlockNums[0], KovanTestEthValueAtBlock)
	})
}

package chainClient

import (
	"ethDemo/util"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"reflect"
	"testing"
)

// test to meta in chain with laboratory
const ConfigFile = "metaTest.yaml"
const blockInfoRequestUrl = "http://testing-metain-js-official-node-service.jiansutech.com/api/mt/getInfo"

// 通用的业务测试功能
type ChainOptionTest interface {
	CheckBlockNum(t *testing.T)
	IsConnected(t *testing.T)
	CallContract(t *testing.T)
	CheckChainNum(t *testing.T)
}

// 元生链 实验室 测试
type MetaChainLabTest struct {
}

func (metaChainLab MetaChainLabTest) CheckChainNum(t *testing.T) {
	myChainConfig := util.NewChainTestYaml(ConfigFile)
	myChainClient := Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
	}

	nowBlockNum := myChainClient.getBlockNumber()
	responseBody, _, _ := util.WebGetRequest(blockInfoRequestUrl)
	if responseBody == nil {
		t.Errorf("blockNum in chain is %v\nblockNum in web si %v ", nowBlockNum, nil)
	}

	var getInfoJson = util.GetInfoJson{}
	webBlockNum := util.ConvertBody2Json(responseBody, getInfoJson, "BlockNumber")
	if nowBlockNum != webBlockNum {
		t.Errorf("blockNum in chain is %v\nblockNum in web si %v ", nowBlockNum, webBlockNum)
	}

}

func (metaChainLab MetaChainLabTest) CheckBlockNum(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func (metaChainLab MetaChainLabTest) CallContract(t *testing.T) {
	myChainConfig := util.NewChainTestYaml(ConfigFile)
	myChainClient := Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
	}

	title := "Title-goLang"
	name := "bobo-go"
	content := "content-by-golang"
	//myChainClient.CallContract(*myChainClient, myChainConfig.ContractAddressHex, myChainConfig.AdminPrivateKeyHex, title, name, content)
	myChainClient.CallContract(title, name, content)

}

func (metaChainLab MetaChainLabTest) IsConnected(t *testing.T) {
	myChainConfig := util.NewChainTestYaml(ConfigFile)
	myChainClient := Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
	}

	//myChainConfig := util.LoadConfig(ConfigFile)
	////println(myChainConfig.ChainUrl)
	//ansClient, err := Launch(myChainConfig.ChainUrl)
	//if ansClient == nil || err != nil {
	//	t.Errorf("Can't get Client")
	//}
}

//
//  TestLaunch
//  @Description: 区块链网络连接测试
//  @param t
//
func TestLaunch(t *testing.T) {
	t.Run("Launch", func(t *testing.T) {
		t.Helper()
		var metaChainOptionTest ChainOptionTest
		metaChainOptionTest = new(MetaChainLabTest)
		metaChainOptionTest.IsConnected(t)
	})
}

//
//  TestGetBlockNumber
//  @Description:
//  @param t
//
func TestGetBlockNumber(t *testing.T) {
	t.Run("GetBlockNumber", func(t *testing.T) {
		t.Helper()
		var metaChainOptionTest ChainOptionTest
		metaChainOptionTest = new(MetaChainLabTest)
		metaChainOptionTest.CheckChainNum(t)
	})
}

func TestCallContract(t *testing.T) {
	t.Run("CAllContract", func(t *testing.T) {
		t.Helper()

		var metaChainOptionTest ChainOptionTest
		metaChainOptionTest = new(MetaChainLabTest)
		metaChainOptionTest.CallContract(t)
		//chainUrl := "http://172.17.4.13:7755"
		//ansClient, err := Launch(chainUrl)
		//if ansClient == nil || err != nil {
		//	t.Fatal("Can't get Client")
		//}
		//privateKeyHex := "794c479028076af7673a6941185af09a51c86a44082b438dbdfca70b6c6829ed"
		//contractAddressHex := "0x03Bc2D794B2FcDA47a9dBb1d43B1fA7B05260282"
		//title := "Title-goLang"
		//name := "bobo-go"
		//content := "content-by-golang"
		//CallContract(*ansClient, contractAddressHex, privateKeyHex, title, name, content)

	})
}

func TestDo(t *testing.T) {
	t.Run("Do", func(t *testing.T) {
		t.Helper()
		//chainClient, err := ethclient.Dial("ADD_YOUR_ETHEREUM_NODE_URL")
		//chainClient, err := ethclient.Dial("https://cloudflare-eth.com")
		// 教程中使用了cloudflare-ech.com
		// 也可以使用自己的eth节点，这里我们尝试使用infura的免费服务。infura 需要新建项目，并且配置相应的测试网络和网络节点url。
		//https://infura.io/dashboard/ethereum/3f97ae7214cc4e2794bee5bdc3bd6b95/settings
		// Details > KEYS > ENDPOINTS (KOVAN) > https://kovan.infura.io/v3/<projectkey>
		// kovan 测试网络中的测试币可以通过水龙头领取
		infuraKovanUrl := "https://kovan.infura.io/v3/3f97ae7214cc4e2794bee5bdc3bd6b95"
		client, err := ethclient.Dial(infuraKovanUrl)
		//chainClient, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/3f97ae7214cc4e2794bee5bdc3bd6b95")
		fmt.Printf("Type %v is %v \n", reflect.TypeOf(client), client)
		if err != nil {
			log.Fatalf("Oops! There was a problem %v", err)
			os.Exit(100)
		} else {
			fmt.Println("Success! you are connected to the Ethereum Network")
		}
		myAddress := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
		kovanTestEthAddress := "0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c"
		myAddressEthValue := GetBalanceFromAddress(*client, myAddress)
		KovanTestEthValue := GetBalanceFromAddress(*client, kovanTestEthAddress)
		fmt.Printf("myAddress is %v,it's Eth Value in Kovan Testnet is %v\n", myAddress, myAddressEthValue)
		fmt.Printf("myAddress is %v,it's Eth Value in Kovan Testnet is %v\n", kovanTestEthAddress, KovanTestEthValue)

		// 根据账户地址和区块高度查询，区块交易的金额。 这个功能需要https://infura.io/dashboard 中购买 archive Data 功能才能调用api
		//KovanTestEthBlockNums := []int64{31866217, 31866202}
		//KovanTestEthValueAtBlock := getBalanceFromBlockNum(*chainClient, kovanTestEthAddress, KovanTestEthBlockNums[0])
		//fmt.Printf("myAddress is %v,it's in %v BlockNumber ,and Kovan Testnet is %v\n", kovanTestEthAddress, KovanTestEthBlockNums[0], KovanTestEthValueAtBlock)
	})
}

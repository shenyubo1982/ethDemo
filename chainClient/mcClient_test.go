package chainClient

// chain 链测试 package
import (
	"ethDemo/util"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"testing"
)

// t.Error t.Errorf :错误不停止继续进行test
// t.Fatal t.Fatalf :遇错即停

// meta in chain 实验室用例类
type MetaChainLabTest struct {
	// 测试用例中需要使用的常量，可以在这边先定义变量名称。在NewXXXX中，用的常量内容。
	ConfigFile          string // "metaTest.yaml"
	blockInfoRequestUrl string //"http://testing-metain-js-official-node-service.jiansutech.com/api/mt/getInfo"
}

// 测试用例的构造器(在运行单独测试用例时，必须先构造生成测试实例 , 测试用例需要用到的个性配置可以在此处通过变量的方式预存)
func NewMetaChainLabTest() *MetaChainLabTest {
	instance := new(MetaChainLabTest)
	//预存测试过程中需要使用的配置
	instance.ConfigFile = "metaTest.yaml"
	instance.blockInfoRequestUrl = "http://testing-metain-js-official-node-service.jiansutech.com/api/mt/getInfo"
	return instance
}

func (ct MetaChainLabTest) Transact(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func (ct MetaChainLabTest) CreateAccount(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func (ct MetaChainLabTest) PressureAttack(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

//
//  CheckChainNum
//  @Description: 测试区块链高度是否显示正常。
// 	测试用例逻辑：获取区块链最新区块num，与区块链浏览器中显示的区块链最新区块链是否一致。
//  @ct MetaChainLabTest
//
func (ct MetaChainLabTest) CheckChainNum(t *testing.T) {
	//连接区块链
	myChainConfig := util.NewChainTestYaml(ct.ConfigFile)
	myChainClient := Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
	}
	// 获取区块链上的最新区块高度
	nowBlockNum := myChainClient.getBlockNumber()
	responseBody, _, _ := util.WebGetRequest(ct.blockInfoRequestUrl)
	if responseBody == nil {
		t.Errorf("blockNum in chain is %v\nblockNum in web si %v ", nowBlockNum, nil)
	}
	//调用接口浏览器使用的接口，获取最新区块高度信息
	var getInfoJson = util.GetInfoJson{}
	webBlockNum := util.ConvertBody2Json(responseBody, getInfoJson, "BlockNumber")
	//比较区块高度时是否一致
	if nowBlockNum != webBlockNum {
		t.Errorf("blockNum in chain is %v\nblockNum in web si %v ", nowBlockNum, webBlockNum)
	}
}

//
//  CallContract
//  @Description:
//  @receiver ct
//  @param t
//
func (ct MetaChainLabTest) CallContract(t *testing.T) {
	//连接区块链
	myChainConfig := util.NewChainTestYaml(ct.ConfigFile)
	myChainClient := Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
	}

	//根据业务内容，设定此次调用合约需要的参数
	title := "Title-goLang"
	name := "bobo-go"
	content := "content-by-golang"

	//myChainClient.CallContract(*myChainClient, myChainConfig.ContractAddressHex, myChainConfig.AdminPrivateKeyHex, title, name, content)
	transactHash := myChainClient.CallContract(title, name, content)
	if transactHash == "" {
		t.Errorf("交易失败，没有获取交易hash")
	} else {
		fmt.Printf("call contract succeed. hash is %v\n", transactHash)
	}
	//TODO 应该再根据交易has去链上验证交易内容。（交易查询功能还未实现，可以放在chainClient类中,作为基础功能来实现后，此处作为业务测试用例进行调用）
}

//
//  IsConnected
//  @Description: 区块链网络连接测试
//  @receiver ct
//  @param t
//
func (ct MetaChainLabTest) IsConnected(t *testing.T) bool {
	//连接区块链
	myChainConfig := util.NewChainTestYaml(ct.ConfigFile)
	myChainClient := Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
		return false
	}
	return true
}

func (ct MetaChainLabTest) mySelfTestFunc(t *testing.T) {
	infuraKovanUrl := "https://kovan.infura.io/v3/3f97ae7214cc4e2794bee5bdc3bd6b95"
	client, err := ethclient.Dial(infuraKovanUrl)
	if err != nil {
		log.Fatalf("Oops! There was a problem %v", err)
		os.Exit(100)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}
	kovanTestEthAddress := "0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c"
	KovanTestEthValue := GetBalanceFromAddress(*client, kovanTestEthAddress)
	fmt.Printf("myAddress is %v,it's Eth Value in Kovan Testnet is %v\n", kovanTestEthAddress, KovanTestEthValue)

	// 根据账户地址和区块高度查询，区块交易的金额。 这个功能需要https://infura.io/dashboard 中购买 archive Data 功能才能调用api
	//KovanTestEthBlockNums := []int64{31866217, 31866202}
	//KovanTestEthValueAtBlock := getBalanceFromBlockNum(*chainClient, kovanTestEthAddress, KovanTestEthBlockNums[0])
	//fmt.Printf("myAddress is %v,it's in %v BlockNumber ,and Kovan Testnet is %v\n", kovanTestEthAddress, KovanTestEthBlockNums[0], KovanTestEthValueAtBlock)

}

//
//  TestLaunch
//  @Description: 区块链网络连接测试用例逻辑
//  @param t
//
func TestLaunchMetaLab(t *testing.T) {
	t.Run("TestLaunchMetaLab", func(t *testing.T) {
		t.Helper()
		var metaChainOptionTest ChainTestingCase
		metaChainOptionTest = NewMetaChainLabTest()
		astRet := metaChainOptionTest.IsConnected(t)
		if astRet != true {
			t.Errorf("应该 %v , 结果为 %v\n ", true, astRet)
		}
	})
}

//Test开始的是测试用例，用go test 工具会执行的测试用例。
//  TestChainNumMetaLab 链上最新区块高度,
//  @Description:
//  @param t
//
func TestChainNumMetaLab(t *testing.T) {
	t.Run("TestChainNumMetaLab", func(t *testing.T) {
		t.Helper()
		var metaChainOptionTest ChainTestingCase
		metaChainOptionTest = NewMetaChainLabTest()
		metaChainOptionTest.CheckChainNum(t)
	})
}

//Test开始的是测试用例，用go test 工具会执行的测试用例。
//  TestCallContractMetaLab 调用智能合约方法
//  @Description:
//  @param t
//
func TestCallContractMetaLab(t *testing.T) {
	t.Run("TestCallContractMetaLab", func(t *testing.T) {
		t.Helper()
		var metaChainOptionTest ChainTestingCase
		metaChainOptionTest = NewMetaChainLabTest()
		metaChainOptionTest.CallContract(t)
	})
}

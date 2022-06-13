package business

// 业务测试 mt（链名称） Lab（业务名称）_test(测试用例)
import (
	"ethDemo/chainClient"
	"ethDemo/util"
	"fmt"
	"math/big"
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

//
//  Transact
//  @Description: 单个账户转账测试 (case 不正确，需要调整，from地址的val变化没有区别，实际上是变化了)
//  @receiver ct
//  @param t
//
func (ct MetaChainLabTest) Transact(t *testing.T) {
	//Todo Admin Address to Test Account transact price.
	//连接区块链
	myChainConfig := util.NewChainTestYaml(ct.ConfigFile)
	myChainClient := chainClient.Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
	}

	// Admin account's priKeyHex
	priKeyHex := myChainClient.ChainConfig().AdminPrivateKeyHex

	//get a test Account
	keyDir := "../keys/mt"
	iWantCnt := 1
	cas := chainClient.LoadChainAccount(iWantCnt, keyDir)
	// to Account's Address
	toAddressHex := cas.Account(0).Address().Hex()
	//transBeforeToAddress := myChainClient.GetBalanceByAddress(toAddressHex)

	// transact From Account Balance
	//acc1Key, _ := crypto.HexToECDSA(priKeyHex)
	//fromAddressHex := crypto.PubkeyToAddress(acc1Key.PublicKey).Hex()
	//transBeforeFromAddress := myChainClient.GetBalanceByAddress(fromAddressHex)

	//转账金额
	price := int64(1000000000000000000) // in wei (1 eth)
	txHex, err := myChainClient.TransferExchange(priKeyHex, toAddressHex, price)
	if err != nil {
		t.Errorf("交易异常:%v", err)
	}
	//todo 校验exHex 中的交易的内容
	//log.Printf("交易完成了，交易hex:%v\n", txHex)
	//// 获取交易receipt
	//receipt, err := myChainClient.Client().TransactionReceipt(context.Background(), txHex)
	//if err != nil {
	//	t.Errorf("Receipt 异常:%v", err)
	//}

	succeed, err := myChainClient.ExchangeIsSucceed(txHex)
	if err != nil {
		t.Errorf("无法查询交易记录:%v", err)
	}
	if !succeed {
		t.Errorf("交易失败")
	}
	//print(receipt.Status)
	////todo check is rigth?
	//log.Println("===================================")
	//status := util.ReflectReceipt(*receipt, "Status")
	//log.Printf("tx_Status is %v\n", status)
	//log.Println("===================================")

	////transact From Account Balance
	//transAfterFromAddress := myChainClient.GetBalanceByAddress(fromAddressHex)
	////transact To Account Balance
	//transAfterToAddress := myChainClient.GetBalanceByAddress(toAddressHex)
	//
	//// Todo:Bug 验证的方法需要修改。获取From和To转账地址在转账前的金额。发起转账交易后，根据交易id，确认区块链已成功后，再确认From和To转账地址最新的金额。进行比较。
	//
	////检查Transact From 账号的val
	////1：前面的big.Int 实例大于cmp方法big.Int 参数
	//if -1 != transAfterFromAddress.Cmp(transBeforeFromAddress) {
	//	t.Errorf("转账%v没有成功.\nbefore Val:%v \nafter val:%v", fromAddressHex, transBeforeFromAddress, transAfterFromAddress)
	//} else {
	//	log.Printf("转账%v成功\n转账前:%v \n转账之后:%v", fromAddressHex, transBeforeFromAddress, transAfterFromAddress)
	//}
	////检查Transact To 账号的val
	//if 1 != transAfterToAddress.Cmp(transBeforeToAddress) {
	//	t.Errorf("转账%v没有成功.\nbefore Val:%v \nafter val:%v", toAddressHex, transBeforeToAddress, transAfterToAddress)
	//} else {
	//	log.Printf("转账%v成功\n转账前:%v\n转账之后:%v\n", toAddressHex, transBeforeToAddress, transAfterToAddress)
	//}

}

//
//  CreateAccount 接口定义的必须实现的测试case
//  @Description: 创建n个账号，并保存在store中
//  @receiver ct
//  @param newCount
//  @param t
//
func (ct MetaChainLabTest) CreateAccount(newCount int, t *testing.T) {
	keyDir := "../keys/mt"
	iWantCnt := newCount
	cas := chainClient.LoadChainAccount(iWantCnt, keyDir)
	if cas.Cnt() != iWantCnt {
		t.Errorf("生成Account 错误！应该是%v ，实际是 %v", iWantCnt, cas.Cnt())
	}
}

func (ct MetaChainLabTest) PressureAttack(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

//
//  CheckChainNum 自定义测试case
//  @Description: 测试区块链高度是否显示正常。
// 	测试用例逻辑：获取区块链最新区块num，与区块链浏览器中显示的区块链最新区块链是否一致。
//  @ct MetaChainLabTest
//
func (ct MetaChainLabTest) CheckChainNum(t *testing.T) {
	//连接区块链
	myChainConfig := util.NewChainTestYaml(ct.ConfigFile)
	myChainClient := chainClient.Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
	}

	//从业务浏览器中获取最新的区块高度
	var getInfoJson = util.GetInfoJson{}
	webBlockNum := new(big.Int)

	responseBody, _, _ := util.WebGetRequest(ct.blockInfoRequestUrl)
	if responseBody == nil {
		t.Errorf("blockNum in web is %v ", nil)
	}
	//调用接口浏览器使用的接口，获取最新区块高度信息
	webBN := util.SearchBody2Json(responseBody, getInfoJson, "BlockNumber")

	webBlockNum, ok := webBlockNum.SetString(webBN, 10)
	if !ok {
		panic("big int setString is error.")
	}
	//blocknum := util.SearchBody2Json(responseBody, getInfoJson, "BlockNumber")
	//webBlockNum := big.NewInt()

	//// 获取区块链上的最新区块高度
	//nowBlockNum := myChainClient.GetBlockNumber()

	chainBlock := chainClient.GetBlockInfo(*myChainClient, webBlockNum)
	nowBlockNum := chainBlock.BlockNum()
	//比较区块高度时是否一致
	if 0 != nowBlockNum.Cmp(webBlockNum) {
		t.Errorf("blockNum in chain is %v\nblockNum in web is %v ", nowBlockNum, webBlockNum)
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
	myChainClient := chainClient.Launch(myChainConfig.YamlContent)
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
	myChainClient := chainClient.Launch(myChainConfig.YamlContent)
	if myChainClient == nil {
		t.Errorf("Can't get Client")
		return false
	}
	return true
}

// business test from interface
//  TestIsConnected
//  @Description: 区块链网络连接测试用例逻辑
//  @param t
//
func TestIsConnected(t *testing.T) {
	t.Run("TestIsConnected", func(t *testing.T) {
		t.Helper()
		var metaChainOptionTest ChainTestingCase
		metaChainOptionTest = NewMetaChainLabTest()
		astRet := metaChainOptionTest.IsConnected(t)
		if astRet != true {
			t.Errorf("应该 %v , 结果为 %v\n ", true, astRet)
		}
	})
}

// business test from interface
func TestCreateAccount(t *testing.T) {
	t.Run("TestCreateAccount", func(t *testing.T) {
		t.Helper()
		var metaChainOptionTest ChainTestingCase
		metaChainOptionTest = NewMetaChainLabTest()
		metaChainOptionTest.CreateAccount(2, t)
	})
}

// business test
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

// business test
//Test开始的是测试用例，用go test 工具会执行的测试用例。
//  TestTransactExchange 转账
//  @Description:
//  @param t
//
func TestTransactExchange(t *testing.T) {
	t.Run("TestTransactExchange", func(t *testing.T) {
		t.Helper()
		var metaChainOptionTest ChainTestingCase
		metaChainOptionTest = NewMetaChainLabTest()
		metaChainOptionTest.Transact(t)
	})
}

/**/
func TestGetBlockNum(t *testing.T) {
	t.Run("TestGetBlockNum", func(t *testing.T) {
		t.Helper()
		metaChainOptionTest := NewMetaChainLabTest()
		metaChainOptionTest.CheckChainNum(t)
	})

}

////业务测试初始化函数 go test 会被先执行此函数
//func TestMain(m *testing.M) {
//	log.Println("---------测试开始---------")
//
//}

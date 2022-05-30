package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"reflect"
)

func convertWeiToBi(balance *big.Int) (ethValue *big.Float) {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}

func getBalanceFromAddress(client ethclient.Client, address string) *big.Float {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}
	return convertWeiToBi(balance)
}

func getBalanceFromBlockNum(client ethclient.Client, address string, blockNum int64) *big.Float {
	print(blockNum)
	account := common.HexToAddress(address)
	blockNumber := big.NewInt(blockNum)
	print(blockNumber)
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		panic(err)
	}
	return convertWeiToBi(balance)
}

//
//  launch
//  @Description: 启动区块链连接，返回网络客户端对象
//  @param chainUrl
//  @return *ethclient.Client
//  @return error
//
func launch(chainUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(chainUrl)
	fmt.Printf("Type %v is %v \n", reflect.TypeOf(client), client)
	if err != nil {
		log.Fatalf("Oops! There was a problem %v", err)
		return nil, err
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
		return client, nil
	}
	//myAddress := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	//myAddressEthValue := getBalanceFromAddress(*client, myAddress)
	//fmt.Printf("myAddress is %v,it's Eth Value in Kovan Testnet is %v\n", myAddress, myAddressEthValue)

}

func main() {
	////client, err := ethclient.Dial("ADD_YOUR_ETHEREUM_NODE_URL")
	////client, err := ethclient.Dial("https://cloudflare-eth.com")
	//// 教程中使用了cloudflare-ech.com
	//// 也可以使用自己的eth节点，这里我们尝试使用infura的免费服务。infura 需要新建项目，并且配置相应的测试网络和网络节点url。
	////https://infura.io/dashboard/ethereum/3f97ae7214cc4e2794bee5bdc3bd6b95/settings
	//// Details > KEYS > ENDPOINTS (KOVAN) > https://kovan.infura.io/v3/<projectkey>
	//// kovan 测试网络中的测试币可以通过水龙头领取
	//infuraKovanUrl := "https://kovan.infura.io/v3/3f97ae7214cc4e2794bee5bdc3bd6b95"
	//client, err := ethclient.Dial(infuraKovanUrl)
	////client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/3f97ae7214cc4e2794bee5bdc3bd6b95")
	//fmt.Printf("Type %v is %v \n", reflect.TypeOf(client), client)
	//if err != nil {
	//	log.Fatalf("Oops! There was a problem", err)
	//	os.Exit(100)
	//} else {
	//	fmt.Println("Success! you are connected to the Ethereum Network")
	//}
	//myAddress := "0x71c7656ec7ab88b098defb751b7401b5f6d8976f"
	//kovanTestEthAddress := "0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c"
	//myAddressEthValue := getBalanceFromAddress(*client, myAddress)
	//KovanTestEthValue := getBalanceFromAddress(*client, kovanTestEthAddress)
	//fmt.Printf("myAddress is %v,it's Eth Value in Kovan Testnet is %v\n", myAddress, myAddressEthValue)
	//fmt.Printf("myAddress is %v,it's Eth Value in Kovan Testnet is %v\n", kovanTestEthAddress, KovanTestEthValue)
	//
	//// 根据账户地址和区块高度查询，区块交易的金额。 这个功能需要https://infura.io/dashboard 中购买 archive Data 功能才能调用api
	////KovanTestEthBlockNums := []int64{31866217, 31866202}
	////KovanTestEthValueAtBlock := getBalanceFromBlockNum(*client, kovanTestEthAddress, KovanTestEthBlockNums[0])
	////fmt.Printf("myAddress is %v,it's in %v BlockNumber ,and Kovan Testnet is %v\n", kovanTestEthAddress, KovanTestEthBlockNums[0], KovanTestEthValueAtBlock)
}

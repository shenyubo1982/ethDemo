package etherscanClient

import (
	"github.com/nanmu42/etherscan-api"
	"math"
	"math/big"
	"time"
)

type ScanManage struct {
	client    *etherscan.Client
	apiKey    string
	scNetwork etherscan.Network
	//钱包保存所有地址和余额
	wallet map[string]string
	//wallet map[string]big.Float
}

//func (sm *ScanManage) Wallet() map[string]string {
//	return sm.wallet
//}

func (sm *ScanManage) Wallet() map[string]string {
	return sm.wallet
}

// NewClient 创建scanClient
func (sm *ScanManage) NewClient(network etherscan.Network, APIKey string) {
	if APIKey == "" {
		panic("APIKey is empty")
	}
	if network == "" {
		panic("network is empty")
	}
	sm.scNetwork = network
	sm.apiKey = APIKey

	sm.client = etherscan.New(sm.scNetwork, sm.apiKey)
	if sm.client == nil {
		panic("client is nil")
	}
}

//获取单个地址的余额
func (sm *ScanManage) getAccountBalance(address string) *big.Float {
	// 查询账户以太坊余额
	//0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c
	var balance, err = sm.client.AccountBalance(address)
	if err != nil {
		panic(err)
	}
	fBalance := new(big.Float)
	fBalance.SetString(balance.Int().String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}

//更新钱包中所有地址的余额
func (sm *ScanManage) updateAccountsBalance() {
	if sm.wallet == nil || len(sm.wallet) == 0 {
		//钱包为空，不需要更新
		return
	}
	for k, _ := range sm.wallet {
		//todo 将wallet中的地址作为参数传入，获取余额，并存入wallet中
		balance := sm.getAccountBalance(k)
		sm.wallet[k] = balance.String()
		//sm.wallet[k] = balance
	}
}

//将地址列表导入钱包,并导入地址余额
//参数说明：addresses:地址列表
func (sm *ScanManage) NewWallet(addresses []string) {
	walletSize := len(addresses)

	sm.wallet = make(map[string]string, walletSize)
	var apiCallCount = 0
	for _, address := range addresses {
		balance := sm.getAccountBalance(address)
		sm.wallet[address] = balance.String()
		//fmt.Println("Debug Info | address:", address, "balance:", balance.String())
		// Free API calls per second: 5 calls
		// 调用5次，停1秒
		apiCallCount++
		if apiCallCount >= 5 {
			time.Sleep(1 * time.Second)
			apiCallCount = 0
		}
	}
}

//func (sm *ScanManage) NewWallet(addresses []string) {
//	walletSize := len(addresses)
//
//	sm.wallet = make(map[string]*big.Float, walletSize)
//	for _, address := range addresses {
//		balance := sm.getAccountBalance(address)
//		sm.wallet[address] = balance
//	}
//}

// 返回钱包地址和余额
func (sm *ScanManage) getWallet() map[string]string {
	return sm.wallet
}

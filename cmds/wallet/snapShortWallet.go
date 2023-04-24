package wallet

import (
	"errors"
	"ethDemo/etherscanClient"
	"ethDemo/util"
	"fmt"
	"github.com/gookit/color"
	"github.com/nanmu42/etherscan-api"
	"math/big"
	"os"
	"strings"
	"time"
)

// walletAccount 钱包账户: addressNo, address, balance
type walletAccount struct {
	addressNO int       //account index
	address   string    //account address
	balance   big.Float //account balance
}

func (receiver *walletAccount) SetAddressNO(addressNO int) {
	receiver.addressNO = addressNO
}

func (receiver *walletAccount) SetAddress(address string) {
	receiver.address = address
}

func (receiver *walletAccount) SetBalance(balance big.Float) {
	receiver.balance = balance
}

func (receiver walletAccount) AddressNO() int {
	return receiver.addressNO
}

func (receiver walletAccount) Address() string {
	return receiver.address
}

func (receiver walletAccount) Balance() big.Float {
	return receiver.balance
}

// walletManager 构造器
func newWalletAccount(addressNo int, address string, balance big.Float) *walletAccount {
	return &walletAccount{
		addressNO: addressNo,
		address:   address,
		balance:   balance,
	}
}

func (receiver walletAccount) getBalanceByNo(addressNo int) (error, big.Float) {
	if receiver.addressNO != addressNo {
		return errors.New("\"addressNo is not match"), receiver.balance
	}
	return nil, receiver.balance
}

func (receiver walletAccount) getBalanceByAddress(address string) (error, big.Float) {
	if receiver.address != address {
		return errors.New("\"addressNo is not match"), receiver.balance
	}
	return nil, receiver.balance
}

// myWallet 钱包 struct
// accounts: 钱包中所有账户
// keystorePath: 钱包中所有账户的keystore文件夹
type myWallet struct {
	accounts     []walletAccount
	keystorePath string
}

// newWallet 构造器
func newWallet(keystorePath string) *myWallet {
	//todo read keyStorePath and initialize accounts

	//keyMenu := km.InitKeyManagerMenu(keystorePath)
	//if keyMenu == nil {
	//	color.Warn.Println("keyStorePath is not exist.")
	//	os.Exit(1)
	//}

	for i := 0; i < 10; i++ {
		//todo new walletAccount
		//address := "0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c"
		//balance := big.NewFloat(0)
		//wAccount := newWalletAccount(i, "0x1D9b2905b2EC7d9F64022c6e698c0d622A35225c", *big.NewFloat(0))

	}
	//save keystorePath and accounts to myWallet
	return &myWallet{
		keystorePath: keystorePath,
		accounts:     nil,
	}
}

// Run walletManager 入口
func Run() {
	// read Config file
	myWalletConfig := util.NewWalletYaml("walletConfig.yaml")

	// TODO 从终端输入钱包的keystore文件夹导入钱包地址
	//color.Warn.Println("Please input your Wallet keyStore folder.")
	//reader := bufio.NewReader(os.Stdin)
	//dir, _ := reader.ReadString('\n')
	//dir = strings.TrimSpace(dir)

	//read files from dir and get wallet address
	// todo addressPath and fileName from config file
	//addressPath := "/Users/bobo/Dropbox/web3/Account/"
	//fileName := "CoinAccount"
	addresses := initAddress(myWalletConfig.YamlContent.WalletAddressDir, myWalletConfig.YamlContent.WalletAddressFile)

	// import address into wallet
	var client = etherscanClient.ScanManage{}
	// todo newClient from config file
	client.NewClient(etherscan.Goerli, myWalletConfig.YamlContent.EthScanAPI)
	client.NewWallet(addresses)

	//验证：钱包地址余额显示正确
	//showBalances(client)
	//getBalancesSortDesc(client)
	successRows, fileName := snapshotAddress(client, myWalletConfig.YamlContent.WalletSnapShotDir)
	if successRows > 0 {
		// 写入成功
		color.Success.Printf("\nsnapShort success Records:【%d】\n", successRows)
		color.Success.Printf("snapshot file: %s\n", fileName)
		util.Confirm()
	}
}

func initAddress(addressPath string, fileName string) []string {
	// todo read files from dir and get wallet address
	var addresses = util.ReadFileAllContent(addressPath, fileName)
	return addresses
}

func showBalances(scanManager etherscanClient.ScanManage) {
	wallet := scanManager.Wallet()
	for k, v := range wallet {
		fmt.Println("address is ", k, "balance is ", v)
	}
}

//func getBalancesSortDesc(scanManager etherscanClient.ScanManage) {
//	wallet := scanManager.Wallet()
//	sortWallet := sortAddressByValue(wallet)
//	for k, v := range sortWallet {
//		fmt.Println("address is ", k, "balance is ", v)
//	}
//}

func snapshotAddress(scanManager etherscanClient.ScanManage, snapShotDir string) (int, string) {
	wallet := scanManager.Wallet()
	// save snapshot to file
	currentTime := time.Now().Format("20060102150405")

	//数据写入到csv文件
	//首行
	var titles string
	titles = "No,Address,Balance\n"

	var stringBuilder strings.Builder
	stringBuilder.WriteString(titles)

	var index int
	for k, v := range wallet {
		index++
		dataString := fmt.Sprintf("%d,%s,%s\n", index, k, v)
		stringBuilder.WriteString(dataString)
	}
	filename := snapShotDir + "snapShortBalance-" + currentTime + ".csv"
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
	dataString := stringBuilder.String()
	_, err := file.WriteString(dataString)
	if err != nil {
		fmt.Println("writeToCsv error:", err)
	}
	err = file.Close()
	if err != nil {
		fmt.Println("writeToCsv error:", err)
	}
	//写入的数据行数
	return index, filename
}

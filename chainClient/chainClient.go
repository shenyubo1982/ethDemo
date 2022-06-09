package chainClient

import (
	"context"
	"crypto/ecdsa"
	"errors"
	evidencecontract "ethDemo/abi"
	"ethDemo/util"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"reflect"
)

const PathSymbol = "/"

type chainClient struct {
	chainConfig util.YamlContent
	client      *ethclient.Client
	//keyStoreDir string //创建账户保存账号信息的目录
}

func (cc *chainClient) Client() *ethclient.Client {
	return cc.client
}

func (cc *chainClient) ChainConfig() util.YamlContent {
	return cc.chainConfig
}

//func (cc *chainClient) SetKeyStoreDir(keyStoreDir string) {
//	cc.keyStoreDir = keyStoreDir
//}
//
//func (cc *chainClient) KeyStoreDir() string {
//	return cc.keyStoreDir
//}

//
//  convertWeiToValue
//  @Description:  change xxWei to XX(eth)
//  @param balance :
//  @return ethValue
//
func convertWeiToValue(balance *big.Int) (ethValue *big.Float) {
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	ethValue = new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	//log.Printf("Value is : %v", ethValue)
	return ethValue
}

// GetBalanceByAddress
//  @Description: 获取地址上目前的最新balance
//  @receiver cc
//  @param addressHex
//  @return *big.Float
//
func (cc *chainClient) GetBalanceByAddress(addressHex string) *big.Float {
	account := common.HexToAddress(addressHex)
	balance, err := cc.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}
	//log.Printf("balance is : %v", balance)
	return convertWeiToValue(balance)
}

// GetBalanceFromBlockNum
//  @Description:
//  @param client
//  @param address
//  @param blockNum
//  @return *big.Float
//
func GetBalanceFromBlockNum(client ethclient.Client, address string, blockNum int64) *big.Float {
	account := common.HexToAddress(address)
	blockNumber := big.NewInt(blockNum)
	log.Printf("blockNum is : %v", blockNumber)
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		panic(err)
	}
	return convertWeiToValue(balance)
}

// CallContract
//  @Description: 调用区块链网络上的已部署成功的合约和方法， 返回调用合约的交易hax（hex)
//  @param chainClient
//  @param addressHex
//  @param privateKeyHex
//  @param title
//  @param name
//  @param content
//
func (cc *chainClient) CallContract(title string, name string, content string) string {
	// 2. put in your testing private key, make sure it has bsc testnet BNB
	privateKey, err := crypto.HexToECDSA(cc.chainConfig.AdminPrivateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 3. contract address
	address := common.HexToAddress(cc.chainConfig.ContractAddressHex)
	//abi := ReadAbi("./abi/EvidenceContract.abi")
	//fmt.Println(abi)
	//fmt.Println(address)

	instance, err := evidencecontract.NewEvidencecontract(address, cc.client)
	if err != nil {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(77))
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = 9999999
	auth.GasPrice = big.NewInt(1000000000)
	//auth.GasLimit = 9999999
	//auth.GasPrice = big.NewInt(1000000000)

	transactOpts := &bind.TransactOpts{
		From:   fromAddress,
		Nonce:  auth.Nonce,
		Signer: auth.Signer,
		//Value:    big.NewInt(0),
		Value:    big.NewInt(100),
		GasPrice: auth.GasPrice,
		GasLimit: auth.GasLimit,
		Context:  auth.Context,
		NoSend:   false,
	}

	// todo reflect methods
	typeOfContract := reflect.TypeOf(instance)
	for i := 0; i < typeOfContract.NumMethod(); i++ {
		fmt.Printf(
			"method is %s, type is %s, kind is %s.\n",
			typeOfContract.Method(i).Name,
			typeOfContract.Method(i).Type,
			typeOfContract.Method(i).Type.Kind(),
		)
	}
	method, _ := typeOfContract.MethodByName("AddInfo")
	fmt.Printf("method is %s, type is %s, kind is %s.\n", method.Name, method.Type, method.Type.Kind())

	// todo reflect methods

	tx, err := instance.AddInfo(transactOpts, title, name, content)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("tx sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex()
}

// GetBlockNumber
//  @Description: 获取区块头
//  @receiver cc: 区块链客户端
//  @return string 区块头编号
//
func (cc *chainClient) GetBlockNumber() *big.Int {
	header, err := cc.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return header.Number
}

func (cc chainClient) GetBlockInfo(blockNumber *big.Int) {
	block, err := cc.client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144
}

// TransferExchange
//  @Description: 发起交易(转账)
//  @receiver cc 链客户端
//  @param fromAccount 发起交易方
//  @param toAddress 交易目的
//  @param price
//
func (cc *chainClient) TransferExchange(priKeyHex string, toAddressHex string, price int64) (txHex common.Hash, err error) {
	txHex = common.Hash{}
	privateKey, err := crypto.HexToECDSA(priKeyHex)
	if err != nil {
		return txHex, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return txHex, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := cc.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return txHex, err
		//log.Fatal(err)
	}

	//value := big.NewInt(1000000000000000000) // in wei (1 eth)
	value := big.NewInt(price) // in wei (1 eth)
	gasLimit := uint64(21000)  // in units
	gasPrice, err := cc.client.SuggestGasPrice(context.Background())
	if err != nil {
		//log.Fatal(err)
		return txHex, err
	}

	toAddress := common.HexToAddress(toAddressHex)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := cc.client.NetworkID(context.Background())
	if err != nil {
		return txHex, err
		//log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return txHex, err
		//log.Fatal(err)
	}

	err = cc.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return txHex, err
		//log.Fatal(err)
	}

	//返回交易的hex
	txHex = signedTx.Hash()
	return txHex, err
}

//TODO
//查询交易
func (cc chainClient) GetTx(tx_hash string) {
	// 交易查询
	print("Todo")

}

// Launch
//  @Description: client 构造器：启动区块链连接，返回网络客户端对象
//  @param chainUrl
//  @return *ethclient.Client
//  @return error
//
func Launch(myChainConfig util.YamlContent) *chainClient {
	instance := new(chainClient)
	instance.chainConfig = myChainConfig
	client, err := ethclient.Dial(instance.chainConfig.ChainUrl)
	if err != nil {
		log.Fatalf("Oops! There was a problem %v", err)
		return nil
	} else {
		instance.client = client
		return instance
	}
}

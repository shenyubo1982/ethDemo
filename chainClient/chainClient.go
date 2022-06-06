package chainClient

import (
	"context"
	"crypto/ecdsa"
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
)

const PathSymbol = "/"

type chainClient struct {
	chainConfig util.YamlContent
	client      *ethclient.Client
	keyStoreDir string //创建账户保存账号信息的目录
}

func (cc *chainClient) SetKeyStoreDir(keyStoreDir string) {
	cc.keyStoreDir = keyStoreDir
}

func (cc *chainClient) KeyStoreDir() string {
	return cc.keyStoreDir
}

//
//  convertWeiToValue
//  @Description:  change xxWei to XX(eth)
//  @param balance :
//  @return ethValue
//
func convertWeiToValue(balance *big.Int) (ethValue *big.Float) {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	//log.Printf("Value is : %v", ethValue)
	return ethValue
}

func (cc *chainClient) getBalanceByAddress(addressHex string) *big.Float {
	account := common.HexToAddress(addressHex)
	balance, err := cc.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}
	//log.Printf("balance is : %v", balance)
	return convertWeiToValue(balance)
}

//func GetBalanceFromAddress(client ethclient.Client, address string) *big.Float {
//	account := common.HexToAddress(address)
//	balance, err := client.BalanceAt(context.Background(), account, nil)
//	if err != nil {
//		panic(err)
//	}
//	log.Printf("balance is : %v", balance)
//	return convertWeiToValue(balance)
//}

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
	tx, err := instance.AddInfo(transactOpts, title, name, content)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("tx sent: %s", tx.Hash().Hex())
	return tx.Hash().Hex()
}

func (cc *chainClient) getBlockNumber() string {
	header, err := cc.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String()) // 5671744
	return header.Number.String()
}

//
//  transferExchange
//  @Description: 发起交易(转账)
//  @receiver cc 链客户端
//  @param fromAccount 发起交易方
//  @param toAddress 交易目的
//  @param price
//
func (cc *chainClient) transferExchange(priKeyHex string, toAddressHex string, price int64) {

	privateKey, err := crypto.HexToECDSA(priKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := cc.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//value := big.NewInt(1000000000000000000) // in wei (1 eth)
	value := big.NewInt(price) // in wei (1 eth)
	gasLimit := uint64(21000)  // in units
	gasPrice, err := cc.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(toAddressHex)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := cc.client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = cc.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
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

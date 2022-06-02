package chainClient

import (
	"context"
	"crypto/ecdsa"
	evidencecontract "ethDemo/abi"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"reflect"
)

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
	log.Printf("Value is : %v", ethValue)
	return ethValue
}

func GetBalanceFromAddress(client ethclient.Client, address string) *big.Float {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}
	log.Printf("balance is : %v", balance)
	return convertWeiToValue(balance)
}

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
//  @Description: 调用区块链网络上的已部署成功的合约和方法
//  @param chainClient
//  @param addressHex
//  @param privateKeyHex
//  @param title
//  @param name
//  @param content
//
func CallContract(client ethclient.Client, addressHex string, privateKeyHex string, title string, name string, content string) {
	// 2. put in your testing private key, make sure it has bsc testnet BNB
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
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
	address := common.HexToAddress(addressHex)
	//abi := ReadAbi("./abi/EvidenceContract.abi")
	//fmt.Println(abi)
	//fmt.Println(address)

	instance, err := evidencecontract.NewEvidencecontract(address, &client)
	if err != nil {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(77))
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = 9999999
	auth.GasPrice = big.NewInt(1000000000)

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
	log.Printf("tx sent: %s", tx.Hash().Hex())
}

// Launch
//  @Description: 启动区块链连接，返回网络客户端对象
//  @param chainUrl
//  @return *ethclient.Client
//  @return error
//
func Launch(chainUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(chainUrl)
	fmt.Printf("Type %v is %v \n", reflect.TypeOf(client), client)
	if err != nil {
		log.Fatalf("Oops! There was a problem %v", err)
		return nil, err
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
		return client, nil
	}
}

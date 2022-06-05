package chainClient

import (
	"context"
	"crypto/ecdsa"
	evidencecontract "ethDemo/abi"
	"ethDemo/util"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
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

//创建钱包
func (cc *chainClient) createNewWallet() {
	fmt.Println("create new Wallate")
}

func (cc *chainClient) getBlockNumber() string {
	header, err := cc.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String()) // 5671744
	return header.Number.String()
}

func creatAccount() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
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

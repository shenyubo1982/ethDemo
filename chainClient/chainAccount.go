package chainClient

import (
	"crypto/ecdsa"
	"ethDemo/util"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
)

const accountPwd = "secret_mt"

type chainAccount struct {
	priKey  *ecdsa.PrivateKey
	address common.Address
	pubKey  *ecdsa.PublicKey
	keyDir  string
	keyFile string
}

func (c chainAccount) Address() common.Address {
	return c.address
}

func (c chainAccount) KeyDir() string {
	return c.keyDir
}

func (c chainAccount) KeyFile() string {
	return c.keyFile
}

func (c chainAccount) AddressHex() string {
	return c.address.Hex()
}

func (c chainAccount) PriKey() *ecdsa.PrivateKey {
	return c.priKey
}

// PriKeyHex
//  @Description: 获取账户私钥hash
//  @receiver c
//  @return string
//
func (c chainAccount) PriKeyHex() string {
	return common.BigToHash(c.priKey.D).Hex()
}

// PubKey
//  @Description: 获取公钥对象
//  @receiver c
//  @return ecdsa.PublicKey
//
func (c chainAccount) PubKey() *ecdsa.PublicKey {
	return c.pubKey
}

func (c chainAccount) PubKeyHex() string {
	return hexutil.Encode(crypto.FromECDSAPub(c.pubKey))
}

type chainAccounts struct {
	accounts []chainAccount
	cnt      int
}

func (c chainAccounts) Accounts() []chainAccount {
	return c.accounts
}

func (c chainAccounts) Account(index int) chainAccount {
	return c.accounts[index]
}

// Cnt
//  @Description: 获取已加载的链账号个数
//  @receiver c
//  @return int
//
func (c chainAccounts) Cnt() int {
	return c.cnt
}

// AddAccount
//  @Description: 在chainAccounts中添加一个account
//  @receiver c: chainAccounts
//  @param keyFile : UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3
//
func (c *chainAccounts) AddAccount(keyDir string, keyFile string) {
	if c.accounts == nil {
		c.accounts = make([]chainAccount, 0, 1)
	}
	c.accounts = append(c.accounts, *loadChainAccountFromKeyFile(keyDir, keyFile))
	c.cnt += 1
}

// LoadChainAccount
//  @Description:  构造器，根据要生成多少个测试账号来生成多个account
//  @param loadCnt 希望生成的account数目
//  @param keyFileDirectory 过往实验用过的账号，通过keyFile加载 ../keys/mt/
//  @return *chainAccounts
//
func LoadChainAccount(loadCnt int, keyFileDirectory string) *chainAccounts {

	instance := new(chainAccounts)
	instance.accounts = make([]chainAccount, 0, loadCnt)
	existedKeyCnt := 0
	var keyFiles []string
	if keyFileDirectory != "" {
		//需要从存量账号中获取测试账号
		keyFiles := util.GetFilesInDir(keyFileDirectory)
		existedKeyCnt = len(keyFiles)
	}
	//填补不足的账号
	for i := 0; i < loadCnt-existedKeyCnt; i++ {
		createAccountWithKs(keyFileDirectory)
	}
	//重新获取加载的文件列表
	keyFiles = util.GetFilesInDir(keyFileDirectory)
	if len(keyFiles) < loadCnt {
		panic("out of chain account instance.")
	}

	for i := 0; i < loadCnt; i++ {
		// 从keyfile导入账号
		instance.AddAccount(keyFileDirectory, keyFiles[i])
	}
	return instance
}

//
//  loadChainAccountFromKeyFile
//  @Description: 加载1个Address 从 key 文件库中
//  @param keyDir : keyStore 目录
//  @param keyFile: keyStore name
//  @return *chainAccount account
//
func loadChainAccountFromKeyFile(keyDir string, keyFile string) *chainAccount {
	instance := new(chainAccount)
	//预存测试过程中需要使用的配置
	//instance.ConfigFile = "metaTest.yaml"
	//instance.blockInfoRequestUrl = "http://testing-metain-js-official-node-service.jiansutech.com/api/mt/getInfo"
	// 配置保存 key store 的目录
	//file := "../key/mt/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
	// check file is existed and actually a file.
	storeFile := keyDir + PathSymbol + keyFile
	if !util.Exists(storeFile) || util.IsDir(storeFile) {
		panic("block chain Account Load Fail. Key File is not correct.")
	}
	//ks := keystore.NewKeyStore(cc.keyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(storeFile)
	if err != nil {
		panic("block chain Account Load Key File is Failed.")
	}
	key, err := keystore.DecryptKey(jsonBytes, accountPwd)
	if err != nil {
		panic("block chain Account Decrypt Kye Failed.")
	}
	instance.priKey = key.PrivateKey
	instance.address = key.Address
	instance.pubKey = &key.PrivateKey.PublicKey
	instance.keyDir = keyDir
	instance.keyFile = keyFile
	return instance
}

//
//  createAccountWithKs
//  @Description: 用keyStore的方式创建账号
//  @param keyDir KeyStore保存的目录路径 ./dirName
//
func createAccountWithKs(keyDir string) {
	//check
	if keyDir == "" {
		log.Fatal("No store Dir.Please set it first.")
	}
	if !util.IsDir(keyDir) {
		log.Fatal("it's Not Dir.")
	}
	ks := keystore.NewKeyStore(keyDir, keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.NewAccount(accountPwd)
	if err != nil {
		log.Fatal(err)
	}
}

// GetPrivateKey
//  @Description:  从store 文件中 提取私钥
//  @param storeFile full path file name. ex "../key/mt/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
//  @param password : store file 密码
//  @return *ecdsa.PrivateKey  私钥
//
func GetPrivateKey(storeFile string, password string) *ecdsa.PrivateKey {
	// check file is existed and actually a file.
	if !util.Exists(storeFile) || util.IsDir(storeFile) {
		panic("block chain Account Load Fail. Key File is not correct.")
	}
	//ks := keystore.NewKeyStore(cc.keyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(storeFile)
	if err != nil {
		panic("block chain Account Load Key File is Failed.")
	}
	key, err := keystore.DecryptKey(jsonBytes, accountPwd)
	if err != nil {
		panic("block chain Account Decrypt Kye Failed.")
	}
	return key.PrivateKey
}

// GetPrivateKeyHex
//  @Description: 从store 文件中，提取私钥hex
//  @param storeFile full path file name. ex "../key/mt/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
//  @param password : store file 密码
//  @return string : 私钥hex 可以直接用于导入
//
func GetPrivateKeyHex(storeFile string, password string) string {
	priKey := GetPrivateKey(storeFile, password)
	return common.BigToHash(priKey.D).Hex()
}

package chainClient

import (
	"crypto/ecdsa"
	"errors"
	"ethDemo/util"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
)

const accountPwd = "secret_mt"

type chainAccount struct {
	priKey  *ecdsa.PrivateKey //账户私钥对象
	address common.Address    //账户地址
	pubKey  *ecdsa.PublicKey  //公钥
	keyDir  string            //storeKey 目录
	keyFile string            //storeKey 文件
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
//  @Description: 获取账户私钥Hex
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

// 账号集对象
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
	account, _ := LoadChainAccountFromKeyFile(keyDir, keyFile)
	c.accounts = append(c.accounts, *account)
	c.cnt += 1
}

// LoadChainAccount
//  @Description:  构造器，加载制定数量的测试账号。不足的自动生成。
//  @param loadCnt 希望加载的测试账号数目
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
		color.Warnf("LoadChainAccount: %d accounts loaded from %s \r", len(keyFiles), keyFileDirectory)
		existedKeyCnt = len(keyFiles)
	}
	//填补不足的账号
	// add progress bar
	bar := util.ProgressBarConfig(loadCnt-existedKeyCnt, "Account creating ...", 2, 1)
	for i := 0; i < loadCnt-existedKeyCnt; i++ {
		_, err := createAccountWithKs(keyFileDirectory)
		if err != nil {
			log.Printf("LoadChainAccount Running Error: %v \n", err)
		}
		util.ShowProgressBar(bar)
	}
	//重新获取加载的文件列表
	keyFiles = util.GetFilesInDir(keyFileDirectory)
	if len(keyFiles) < loadCnt {
		panic("out of chain account instance.")
	}

	barAdd := util.ProgressBarConfig(loadCnt, "Account Loading...", 2, 2)
	for i := 0; i < loadCnt; i++ {
		// 从keyfile导入账号
		instance.AddAccount(keyFileDirectory, keyFiles[i])
		util.ShowProgressBar(barAdd)
	}
	return instance
}

// LoadChainAccountFromKeyFile
//  @Description: 加载1个Address 从 key 文件库中
//  @param keyDir : keyStore 目录
//  @param keyFile: keyStore name
//  @return *chainAccount account
//
func LoadChainAccountFromKeyFile(keyDir string, keyFile string) (*chainAccount, error) {
	instance := new(chainAccount)
	// check file is existed and actually a file.
	storeFile := keyDir + PathSymbol + keyFile
	if !util.Exists(storeFile) || util.IsDir(storeFile) {
		return nil, errors.New("block chain Account Load Fail. Key File is not correct")
	}
	//ks := keystore.NewKeyStore(cc.keyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(storeFile)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(jsonBytes, accountPwd)
	if err != nil {
		return nil, err
	}
	instance.priKey = key.PrivateKey
	instance.address = key.Address
	instance.pubKey = &key.PrivateKey.PublicKey
	instance.keyDir = keyDir
	instance.keyFile = keyFile
	return instance, nil
}

//
//  createAccountWithKs
//  @Description: 用keyStore的方式创建账号
//  @param keyDir KeyStore保存的目录路径 ./dirName
//
func createAccountWithKs(keyDir string) (bool, error) {
	//check
	if keyDir == "" {
		return false, errors.New("no store Dir.Please set it first")
	}
	if !util.IsDir(keyDir) {
		return false, errors.New("it's Not Dir")
	}
	ks := keystore.NewKeyStore(keyDir, keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.NewAccount(accountPwd)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetPrivateKey
//  @Description:  从store 文件中 提取私钥
//  @param storeFile full path file name. ex "../key/mt/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
//  @param password : store file 密码
//  @return *ecdsa.PrivateKey  私钥
//
func GetPrivateKey(storeFile string, password string) (*ecdsa.PrivateKey, error) {
	// check file is existed and actually a file.
	if !util.Exists(storeFile) || util.IsDir(storeFile) {
		return nil, errors.New("block chain Account Load Fail. directory or file name is not correct")
	}
	jsonBytes, err := ioutil.ReadFile(storeFile)
	if err != nil {
		return nil, err
	}
	key, errDecrypt := keystore.DecryptKey(jsonBytes, password)
	if errDecrypt != nil {
		return nil, errDecrypt
	}
	return key.PrivateKey, nil
}

// GetPrivateKeyHex
//  @Description: 从store 文件中，提取私钥hex
//  @param storeFile full path file name. ex "../key/mt/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
//  @param password : store file 密码
//  @return string : 私钥hex 可以直接用于导入
//
func GetPrivateKeyHex(storeFile string, password string) (string, error) {
	priKey, err := GetPrivateKey(storeFile, password)
	if err != nil {
		return "", err
	}
	return common.BigToHash(priKey.D).Hex(), nil
}

package chainClient

import (
	"crypto/ecdsa"
	"ethDemo/util"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
)

const accountPwd = "secret_mt"

type chainAccount struct {
	priKey              *ecdsa.PrivateKey
	address             common.Address
	pubKey              ecdsa.PublicKey
	ConfigFile          string
	blockInfoRequestUrl string
	keyDir              string
	keyFile             string
}

type chainAccounts struct {
	accounts []chainAccount
	cnt      int
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
		// 如果存量有数据，就先去keyfile的账号
		instance.accounts = append(instance.accounts, *loadChainAccountFromKeyFile(keyFiles[i]))
		instance.cnt++
	}
	return instance
}

// LoadChainAccountFromKeyFile
//  @Description: 加载1个Address 从 key 文件库中
//  @param file UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3
//  @return *chainAccount
//
func loadChainAccountFromKeyFile(file string) *chainAccount {
	instance := new(chainAccount)
	//预存测试过程中需要使用的配置
	instance.ConfigFile = "metaTest.yaml"
	instance.blockInfoRequestUrl = "http://testing-metain-js-official-node-service.jiansutech.com/api/mt/getInfo"
	// 配置保存 key store 的目录
	instance.keyDir = "../keys/mt"
	//file := "../key/mt/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
	storeFile := instance.keyDir + PathSymbol + file
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
	instance.priKey = key.PrivateKey
	instance.address = key.Address
	instance.pubKey = key.PrivateKey.PublicKey
	instance.keyFile = file
	return instance
}

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

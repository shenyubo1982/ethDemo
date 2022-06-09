package chainClient

//chainAccount.go 单体测试
import (
	"testing"
)

const ConfigFile = "metaTest.yaml"

//
//  TestLaunch
//  @Description:
//  @param t
//
func TestLaunch(t *testing.T) {
	t.Run("Launch", func(t *testing.T) {
		t.Helper()
		keyDir := "../keys/mt"
		iWantCnt := 15
		cas := LoadChainAccount(iWantCnt, keyDir)
		if cas.cnt != iWantCnt {
			t.Errorf("生成Account 错误！应该是%v ，实际是 %v", iWantCnt, cas.cnt)
		}
		//for _, ca := range cas.accounts {
		//	address := ca.address
		//	priKey := ca.priKey
		//	//keyFile := ca.keyFile
		//	log.Printf("Address is %v \n privateKey is %v\n", address, priKey)
		//	log.Printf("--------------------------------")
		//}
	})
}

func TestAddAccount(t *testing.T) {
	keyDir := "../keys/mt"

	t.Run("Add 1 Account", func(t *testing.T) {
		t.Helper()
		keyfileNames := []string{"UTC--2022-06-05T09-29-29.614223000Z--b319c54942ab601d89cab2654afaf9496ddaaa7d"}
		cas := new(chainAccounts)
		cas.AddAccount(keyDir, keyfileNames[0])
		if cas.Cnt() != 1 {
			t.Errorf("应该加载%v个Account,实际是%v", len(keyfileNames), cas.Cnt())
		}
		if len(cas.Accounts()) != 1 {
			t.Errorf("应该加载%v个Account,实际是%v", len(keyfileNames), cas.Cnt())
		}
		//判断加载的账号信息是否完整
		if cas.Account(0).priKey == nil {
			t.Errorf("加载的账号私钥为空")
		}
		if cas.Account(0).PriKeyHex() == "" {
			t.Errorf("加载的账号私钥Hex不正确")
		}
		if cas.Account(0).keyFile != keyfileNames[0] {
			t.Errorf("加载的账号文件名不匹配")
		}
		if cas.Account(0).keyDir != keyDir {
			t.Errorf("加载的账号目录不匹配")
		}
		if cas.Account(0).AddressHex() == "" {
			t.Errorf("加载的账号地址Hex不正确")
		}
		if cas.accounts[0].PubKey() == nil {
			t.Errorf("加载的账号PubKey不正确")
		}
		if cas.Account(0).PubKeyHex() == "" {
			t.Errorf("加载的账号共钥Hex不正确")
		}

		if cas.accounts[0].Address().Hex() == "" {
			t.Errorf("加载的账号Address不正确")
		}
		if cas.accounts[0].KeyDir() == "" {
			t.Errorf("加载的账号KeyDir不正确")
		}
		if cas.accounts[0].KeyFile() == "" {
			t.Errorf("加载的账号KeyFile不正确")
		}
		if cas.accounts[0].PriKeyHex() == "" {
			t.Errorf("加载的账号PriKeyHex不正确")
		}
		if cas.accounts[0].PriKey() == nil {
			t.Errorf("加载的账号PriKey不正确")
		}

	})

	t.Run("Add 2 Account", func(t *testing.T) {
		t.Helper()
		keyfileNames := []string{"UTC--2022-06-05T09-29-29.614223000Z--b319c54942ab601d89cab2654afaf9496ddaaa7d",
			"UTC--2022-06-05T09-29-30.994525000Z--bbf1e2529432fa7735a573ffbbe0f698dd527062"}
		cas := new(chainAccounts)
		for i := 0; i < len(keyfileNames); i++ {
			cas.AddAccount(keyDir, keyfileNames[i])
		}
		if cas.Cnt() != len(keyfileNames) {
			t.Errorf("应该加载%v个Account,实际是%v", len(keyfileNames), cas.Cnt())
		}
	})
}

//
//  TestGetPrivateKey
//  @Description: 从keystroe中获取privateKey的方法
//  @param t
//
func TestGetPrivateKey(t *testing.T) {
	keyDir := "../keys/mt"
	keyfileName := "UTC--2022-06-05T09-29-29.614223000Z--b319c54942ab601d89cab2654afaf9496ddaaa7d"
	storeFile := keyDir + PathSymbol + keyfileName
	storeFileErr := keyDir + PathSymbol + keyfileName + "_err"
	pwd := accountPwd
	pwdErr := accountPwd + "Err"

	t.Run("GetPrivateKey-Correct", func(t *testing.T) {
		t.Helper()
		_, err := GetPrivateKey(storeFile, pwd)
		if err != nil {
			t.Errorf("GetPrivateKey 报错 %v", err)
		}
	})

	t.Run("GetPrivateKey-DirError", func(t *testing.T) {
		t.Helper()
		_, err := GetPrivateKey(storeFileErr, pwd)
		if err == nil {
			t.Errorf("应该报错，没有报错")
		}
	})

	t.Run("GetPrivateKey-passwordError", func(t *testing.T) {
		t.Helper()
		_, err := GetPrivateKey(storeFile, pwdErr)
		if err == nil {
			t.Errorf("应该报错，没有报错")
		}
	})

	t.Run("GetPrivateKeyHex-Correct", func(t *testing.T) {
		t.Helper()
		_, err := GetPrivateKeyHex(storeFile, pwd)
		if err != nil {
			t.Errorf("GetPrivateKey 报错 %v", err)
		}
	})
}

func TestCreateAccountWithKs(t *testing.T) {

	t.Run("createAccountWithKs-Correct", func(t *testing.T) {
		t.Helper()
		keyDir := "../keys/mt"
		_, err := createAccountWithKs(keyDir)
		if err != nil {
			t.Errorf("创建账号异常")
		}
	})

	t.Run("createAccountWithKs-DirEmpty", func(t *testing.T) {
		t.Helper()
		keyDir := ""
		_, err := createAccountWithKs(keyDir)
		if err == nil {
			t.Errorf("创建账号异常没有显示")
		}

	})

	t.Run("createAccountWithKs-DirNotExist", func(t *testing.T) {
		t.Helper()
		keyDir := "../keysErr/mt"
		_, err := createAccountWithKs(keyDir)
		if err == nil {
			t.Errorf("创建账号异常没有显示")
		}
	})
}

func TestLoadChainAccountFromKeyFile(t *testing.T) {
	t.Run("createAccountWithKs-keyKeyNotExist", func(t *testing.T) {
		t.Helper()
		loadChainAccountFromKeyFile("errDir", "errorFileName")
	})

}

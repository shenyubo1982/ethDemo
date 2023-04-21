package ac

import (
	"bufio"
	chainClient "ethDemo/chainClient"
	"ethDemo/util"
	"github.com/gookit/color"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// 命令用途：创建以太坊账号，并保存在制定的目录下，默认是在 ../keys/eth

// Run
//  @Description: start form terminal menu , and create eth Acount(address and sn).
//
func Run() {
	// TODO: add your code here
	color.Warn.Println("How much Account you want to Create?")
	reader := bufio.NewReader(os.Stdin)
	accountNumStr, _ := reader.ReadString('\n')
	accountNumStr = strings.TrimSpace(accountNumStr)

	// todo 判断accountNum是否是数字
	if accountNumStr == "" {
		color.Warn.Println("Please input number")
		accountNumStr = "3"
	}
	accountNum, err := strconv.Atoi(accountNumStr)
	if err != nil {
		color.Warn.Println("Please input number")
		os.Exit(0)
	}

	createAccount(accountNum)

}

func createAccount(iWantCnt int) {
	// create eth account
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	color.Warn.Println(dir)
	keyDir := "../keys/eth"
	cas := chainClient.LoadChainAccount(iWantCnt, keyDir)
	if len(cas.Accounts()) != iWantCnt {
		color.Warn.Println("\ncreate account error: number of account is not equal to iWantCnt.")
		os.Exit(0)
	}
	util.Confirm()

	// save address file

	// save keystore file

	// save sn file

}

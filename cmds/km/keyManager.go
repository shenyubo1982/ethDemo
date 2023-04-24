package km

import (
	"bufio"
	"errors"
	"ethDemo/chainClient"
	"ethDemo/util"
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"os"
	"strings"
)

type KeyStoreMenu struct {
	//file path (dir)
	filePath string
	// filename k:fileName v:address
	fileName map[string]string
	//private key k:address v:privateKey
	privateKey map[string]string
}

// InitKeyManagerMenu read files from dir
func (km *KeyStoreMenu) InitKeyManagerMenu(serverPath string) {
	km.filePath = serverPath
	files, _ := ioutil.ReadDir(serverPath)

	// 根据读取文件的根数，初始化map的大小
	keyCount := len(files)
	km.privateKey = make(map[string]string, keyCount)
	km.fileName = make(map[string]string, keyCount)

	bar := util.ProgressBarConfig(keyCount, "Key Store Manager is Loading...", 1, 1)

	for _, f := range files {
		// set address and private Key from keystore file
		//fmt.Println("Debug | file name: ", f.Name())
		km.setMenuInfo(f.Name())
		util.ShowProgressBar(bar)
	}
}

// SelectKeyManagerMenu 选择菜单项（输入地址） 返回地址对应的私钥
func (km KeyStoreMenu) selectKeyManagerMenu() {
	for {
		// 选择菜单项
		selectedAddress, err := selectFunc()
		// 显示选择的地址和私钥
		if err == nil && (selectedAddress != "") {
			priKey := km.privateKey[selectedAddress]
			//show keyStore Address
			if priKey == "" {
				color.Warn.Println("Address is not exist.")
			} else {
				color.BgCyan.Print("Address: ")
				color.BgYellow.Println(selectedAddress)
				color.BgCyan.Print("PriKey:  ")
				color.BgYellow.Println(priKey)
			}
		}
		// 退出
		if err != nil && err.Error() == "quit" {
			break
		}
		// 显示所有地址
		if err != nil && err.Error() == "show" {
			km.showKeyManagerMenu()
		}
	}
}

// GetAddressInfo get address and privateKey from KeyStoreMenu convert to a new map about {address|priKey}
func (km *KeyStoreMenu) setMenuInfo(fileName string) {
	//show keyStore Address
	chainAccount, err := chainClient.LoadChainAccountFromKeyFile(km.filePath, fileName)
	if err == nil {
		address := chainAccount.AddressHex()
		priKey := chainAccount.PriKeyHex()
		//fmt.Println("Debug | address: ", address)
		//fmt.Println("Debug | priKey: ", priKey)
		// key:address value:priKey
		km.privateKey[address] = priKey
		km.fileName[fileName] = address
	}
}

// ShowKeyManagerMenu show address menu in terminal.
func (km KeyStoreMenu) showKeyManagerMenu() {
	// clean screen
	fmt.Println("")
	menuId := 0
	// filename k:fileName v:address
	for _, address := range km.fileName {
		menuId++
		listNum := fmt.Sprintf("%02d ", menuId)
		color.BgCyan.Print(listNum)
		color.BgYellow.Println(address)
	}
}

// Run
//  @Description: start form terminal menu , and show eth keystore address.
//
func Run() {
	// 1. 选择路径
	//choose keystore from your local dir
	//set default dir and show input message.
	defaultDir := "../keys/wallet"
	defaultServerPath := util.GetRunPath(defaultDir, "") //get abs path
	color.Warn.Println("Your Default keyStore folder is: " + defaultServerPath)
	color.Warn.Println("Press the keyboard Enter key. IF not in there, please input your keyStore folder.")
	reader := bufio.NewReader(os.Stdin)
	dir, _ := reader.ReadString('\n')
	dir = strings.TrimSpace(dir)
	keyMenu := new(KeyStoreMenu)
	// 2. 加载keystore文件，生成菜单
	if dir == "" {
		// 加载文件路径下的keystore文件，并解析生成菜单
		keyMenu.InitKeyManagerMenu(defaultServerPath)
	} else {
		keyMenu.InitKeyManagerMenu(dir)
	}

	// 2. 显示菜单
	keyMenu.showKeyManagerMenu()

	// 3. 选择菜单项，并且返回选择的结果
	keyMenu.selectKeyManagerMenu()

}

// 选择地址，返回地址
func selectFunc() (string, error) {
	color.Red.Print("please input Address [q:Quit |a:Show All Address]: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	selectedAddress := strings.TrimSpace(input)
	if err == nil {
		if selectedAddress == "q" {
			return "Exit", errors.New("quit")
		}
		if selectedAddress == "a" {
			return "Exit", errors.New("show")
		}
		return selectedAddress, nil
	}
	return "Exit", err
}

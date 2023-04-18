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
	"strconv"
	"strings"
)

// Run
//  @Description: start form terminal menu , and show eth keystore address.
//
func Run() {
	//choose keystore from your local dir
	color.Warn.Println("Please input your keyStore folder.")
	reader := bufio.NewReader(os.Stdin)
	dir, _ := reader.ReadString('\n')
	dir = strings.TrimSpace(dir)
	keyMenu := new(keyStoreMenu)
	if dir == "" {
		dir = "keys/mt"
		serverPath := util.GetRunPath(dir, "") //get abs path
		keyMenu = ReadFiles(serverPath)
	} else {
		keyMenu = ReadFiles(dir)
	}

	for {
		// select menuId from screen
		menuNmu, err := selectFunc()
		if err == nil && (menuNmu > 0 && menuNmu < keyMenu.menuCnt+1) {
			filename := keyMenu.fileNames[menuNmu-1]
			//show keyStore Address
			chainAccount, err := chainClient.LoadChainAccountFromKeyFile(keyMenu.filePath, filename)
			if err == nil {
				address := chainAccount.AddressHex()
				color.BgCyan.Print("Address: ")
				color.BgYellow.Println(address)
				priKey := chainAccount.PriKeyHex()
				color.BgCyan.Print("PriKey: ")
				color.BgYellow.Println(priKey)
			}
		}
		if err != nil && err.Error() == "quit" {
			break
		}

	}

}

type keyStoreMenu struct {
	//menu cnt
	menuCnt int
	//file path
	filePath string
	//file name
	fileNames []string
}

//read files from dir
func ReadFiles(serverPath string) *keyStoreMenu {
	fileMenu := new(keyStoreMenu)
	//serverPath := util.GetRunPath(dir, "") //get abs path
	fileMenu.filePath = serverPath
	files, _ := ioutil.ReadDir(serverPath)
	for fileId, f := range files {
		menuId := fileId + 1
		fileMenu.menuCnt = menuId
		fileMenu.fileNames = append(fileMenu.fileNames, f.Name())
		// Modify by bobo show address.do not show file name . start
		// UTC--2023-04-07T05-55-34.618103000Z--
		//debugFileName := fileMenu.fileNames[len("UTC--2023-04-07T05-55-34.618103000Z--"):]
		debugFileName := fileMenu.fileNames[menuId-1]
		//fmt.Println("debugFileName: ", debugFileName)
		//获取第三位到左后位置的debugFileName字符串
		debugFileOutPrintName := debugFileName[len("UTC--2023-04-07T05-55-34.618103000Z--"):]
		// Modify by bobo show address.do not show file name . end

		listNum := fmt.Sprintf("%02d ", menuId)
		color.BgCyan.Print(listNum)
		// Modify by bobo show address.do not show file name . start
		//color.BgYellow.Println(f.Name())
		color.BgYellow.Println(debugFileOutPrintName)
		// Modify by bobo show address.do not show file name . End
	}
	return fileMenu
}

func showKeyFiles() {
	fmt.Println("showKeyFiles")
}

func selectFunc() (int, error) {
	color.Red.Print("please input Number[q:Quit]: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	menuNumStr := strings.TrimSpace(input)
	if err == nil {
		menuNum, err := strconv.Atoi(menuNumStr)
		if err == nil {
			return menuNum, err
		}
		if menuNumStr == "q" {
			return -1, errors.New("quit")
		}
		if !util.IsDigit(menuNumStr) {
			return -1, err
		}
	}
	return -1, err

}

// todo 读取文件并将文件编号，和地址保持在对象中

package menu

import (
	"bufio"
	ac "ethDemo/cmds/ac"
	km "ethDemo/cmds/km"
	//l1 "demo/src/lesson01"
	//l2 "demo/src/lesson02"
	//l3 "demo/src/lesson03"
	//l13 "demo/src/lesson13"
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type MenuContent int

const EXIT = "Exit Demo"
const NA = "N/A"

const (
	exit MenuContent = iota
	//lesson01 MenuContent = iota + 1 //Say "Hello"
	keyManager
	AccountCreate
	//lesson02
	//lesson03
	//lesson04
	//lesson05
	//lesson06
	//lesson07
	//lesson08
	//lesson09
	//lesson10
	//lesson11
	//lesson12
	//lesson13
)

const demoStart = "==========Start==========\n"
const demoEnd = "===========End==========="

func (m MenuContent) String() string {
	switch m {
	case exit:
		return EXIT
	case AccountCreate:
		return "Account Create"
	case keyManager:
		return "KeyStore Manager"
		//case lesson01:
		//	return "Say hello"
		//case lesson02:
		//	return "Guess Game"
		//case lesson03:
		//	return "Function Learning"
		//case lesson04:
		//	return "ToDo 04"
		//case lesson05:
		//	return "ToDo 05"
		//case lesson06:
		//	return "ToDo 06"
		//case lesson07:
		//	return "ToDo 07"
		//case lesson08:
		//	return "ToDo 08"
		//case lesson09:
		//	return "ToDo 09"
		//case lesson10:
		//	return "ToDo 10"
		//case lesson11:
		//	return "ToDo 11"
		//case lesson12:
		//	return "ToDo 12"
		//case lesson13:
		//	return "goRoutine Learning"
	}
	return NA
}

func ShowMenu() {
	color.Red.Printf("==========【Menu】=========\n")
	hasOtherMenu := true
	for i := 0; hasOtherMenu; i++ {
		content := MenuContent(i).String()
		if content == "N/A" {
			hasOtherMenu = false
		} else {
			color.Green.Printf("[%v] %s\n", i, content)
		}
	}

}

func SelectFunc() (int, error) {
	color.Red.Print("please input Number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	menuNumStr := strings.TrimSpace(input)
	if err == nil {
		menuNum, err := strconv.Atoi(menuNumStr)

		if err == nil {
			return menuNum, err
		}
		//whoAmI = strings.Replace(input, "\n", "", -1)
	}
	return -1, err

}

func ShowSelectError() {
	color.Warn.Println("Menu Number Is not correct.")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	ClearScreen()
}

func OrderMenu(m MenuContent) {
	switch m {
	case keyManager:
		doFunc(km.Run)
		//case lesson02:
	case AccountCreate:
		doFunc(ac.Run)
		//	//doFunc(l2.Play)
		//case lesson03:
		//	//doFunc(l3.Do)
		//case lesson13:
		//	//doFunc(l13.Do)
	}

}

func doFunc(a func()) {
	ClearScreen()
	color.BgBlue.Printf(demoStart)
	a()
	color.BgBlue.Printf(demoEnd)
	time.Sleep(300 * time.Millisecond)
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	ClearScreen()
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		fmt.Println("linux is start")
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	clear["darwin"] = func() {
		fmt.Println("darwin is start")
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal(err)
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmds", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
	time.Sleep(300 * time.Millisecond)

}

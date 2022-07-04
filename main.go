package main

import (
	"ethDemo/menu"
	"github.com/gookit/color"
)

//
//  main
//  @Description:
//
func main() {

	for true {
		menu.ShowMenu()
		menuNmu, err := menu.SelectFunc()
		if err != nil {
			//menu.ClearScreen()
			//menu.ShowMenu()
			//color.Warn.Println("Menu Number Is not correct.")
			menu.ShowSelectError()
		} else {
			if menu.MenuContent(menuNmu).String() == menu.EXIT {
				color.Yellow.Println("Bye!")
				break
			}
			if menu.MenuContent(menuNmu).String() == menu.NA {
				//color.Warn.Println("Menu Number Is not correct.")
				menu.ShowSelectError()
			}
			menu.OrderMenu(menu.MenuContent(menuNmu))
		}
	}
	//lesson01.SayHello()
}

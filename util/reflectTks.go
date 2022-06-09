package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"reflect"
)

type Address struct {
	Add string
	Res int
}
type User struct {
	Id      int
	Name    string
	Address Address
}

func ReflectReceipt(u types.Receipt) {
	//u := User{Id: 1001, Name: "aaa", Address: Address{Add: "ccccccccc", Res: 12}}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段

			//判断是否是嵌套结构
			if v.Field(i).Type().Kind() == reflect.Struct {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					fmt.Printf("%s %s = %v -tag:%s \n",
						structField.Field(j).Name,
						structField.Field(j).Type,
						v.Field(i).Field(j).Interface(),
						structField.Field(j).Tag)
				}
				continue
			}

			fmt.Printf("%s %s = %v -tag:%s \n",
				t.Field(i).Name,
				t.Field(i).Type,
				v.Field(i).Interface(),
				t.Field(i).Tag)
		}

	}
}

func reflectDemo(u User) {
	//u := User{Id: 1001, Name: "aaa", Address: Address{Add: "ccccccccc", Res: 12}}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段

			//判断是否是嵌套结构
			if v.Field(i).Type().Kind() == reflect.Struct {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					fmt.Printf("%s %s = %v -tag:%s \n",
						structField.Field(j).Name,
						structField.Field(j).Type,
						v.Field(i).Field(j).Interface(),
						structField.Field(j).Tag)
				}
				continue
			}

			fmt.Printf("%s %s = %v -tag:%s \n",
				t.Field(i).Name,
				t.Field(i).Type,
				v.Field(i).Interface(),
				t.Field(i).Tag)
		}

	}
}

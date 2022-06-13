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

func ReflectReceipt(u types.Receipt, fieldName string, isDebug bool) (i interface{}) {
	//u := User{Id: 1001, Name: "aaa", Address: Address{Add: "ccccccccc", Res: 12}}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			//判断是否是嵌套结构
			if v.Field(i).Type().Kind() == reflect.Struct {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					//嵌套
					if fieldName != "" && fieldName == structField.Field(j).Name {
						//需要匹配
						return v.Field(i).Field(j).Interface()
					}
					if isDebug {
						showReceiptInfo(t, v, i, j)
					}
				}
				continue
			}
			//单层
			if fieldName != "" && fieldName == t.Field(i).Name {
				//需要匹配
				return v.Field(i).Interface()
			}
			if isDebug {
				showReceiptInfo(t, v, i, -1)
			}
		}
	}
	return nil
}

func showReceiptInfo(t reflect.Type, v reflect.Value, i int, j int) {
	if j >= 0 {
		//两层
		fmt.Printf("%s %s = %v -tag:%s \n",
			t.Field(j).Name,
			t.Field(j).Type,
			v.Field(i).Field(j).Interface(),
			t.Field(j).Tag)
	} else {
		//一层
		fmt.Printf("%s %s = %v -tag:%s \n",
			t.Field(i).Name,
			t.Field(i).Type,
			v.Field(i).Interface(),
			t.Field(i).Tag)
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

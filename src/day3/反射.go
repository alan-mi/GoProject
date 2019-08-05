package main

import (
	"fmt"
	"reflect"
)

type C int

func main() {
	//获取对象类型
	var a, b C = 100, 200
	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)
	fmt.Print(ta, tb)
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	fmt.Println(m)

}
func haha() {

}

package main

import (
	"container/list"
	"fmt"
)

func main() {
	//声明列表
	//var lst list.List
	lst := list.New()
	lst.PushFront("adfasdf")
	ele := lst.PushBack(1234)
	print(ele, "\n")
	fmt.Printf("%T =============", ele)
	lst.InsertBefore("hello", ele)
	//遍历列表
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Printf("%T %v \n", e, e)
		fmt.Println(e.Value)
	}
}

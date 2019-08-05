package main

import "fmt"

type tester interface {
	test()
	string() string
}
type data struct {
}
type node struct {
	//匿名接口
	tes interface {
		string() string
	}
}

func (*data) test()         {}
func (data) string() string { return "a" }

func main() {
	//var d data
	//var t tester = &d
	//t.string()
	var t interface{ string() string } = data{}
	n := node{tes: t}
	print(n.tes.string())
	aaa()
}

type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func aaa() {
	var t fmt.Stringer = FuncString(func() string { return "hello world" })
	fmt.Print(t)
}

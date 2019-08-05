package main

import (
	"fmt"
)

type tester interface {
	//test()
	string() string
}

type data struct {
	name string
}

func (d *data) test() {
	print(d.name)
}
func (d *data) string() string {
	d.name = "mi"
	return "string"
}

func main() {
	var d data
	//var t tester = d
	//fmt.Printf("%T ,%q \n",t,t)
	d.string()
	d.test()
	println("----------------------")
	func(t *data) {
		t.name = "asdf"
	}(&d)
	fmt.Println(d)

}

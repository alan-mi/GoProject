package main

import (
	"fmt"
)

var hre = 123

func main() {
	var (
		b string = "adfsdf"
		c float32
		d float64
		f func() bool
		l = 200
	)
	haha := 30
	var a = 10
	fmt.Println(a, b, c, d)
	fmt.Printf("%T,%v \n", a, a)
	fmt.Printf("%T,%v \n", b, b)
	fmt.Printf("%T,%q \n", c, c)
	fmt.Printf("%T,%v \n", d, d)
	fmt.Printf("%T,%v \n", f, f)
	fmt.Printf("%T,%v \n", l, l)
	fmt.Printf("%T,%v \n", haha,haha)
	fmt.Printf("%T,%v \n", hre,hre)
}

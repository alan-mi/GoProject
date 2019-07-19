package main

import "fmt"

type Man struct {
	age,height int
}

func main() {
	const Name = "SFFRER"
	fmt.Println(Name)
	const (
		a = 1<<iota
		b = 3<<iota
		c
	)
	fmt.Println(a,b,c)
}

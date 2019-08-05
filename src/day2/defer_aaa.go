package main

import "fmt"

func main() {
	x := 20

	defer func() {
		print(x)
		fmt.Println(recover())
	}()
	panic("asdf")
	x++
	print(x)
}

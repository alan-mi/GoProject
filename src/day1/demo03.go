package main

import "fmt"

type myFunc func(int) bool

func main() {
	arr := []int{2, 4, 323, 432, 23, 44}
	const (
		a = iota
		b = iota
	)
	print(a, b)
	fmt.Print(fil(arr, func(i int) bool {
		if i%2 == 0 {
			return true
		} else {
			return false
		}
	}))
}

func mytest(a int) bool {
	if a%2 == 0 {
		return false
	} else {
		return true
	}
}

func fil(arr []int, myfunc myFunc) []int {
	var b []int
	for _, v := range arr {
		if myfunc(v) {
			b = append(b, v)
		}
	}
	return b
}

package main

import "fmt"

func main() {
	//a := Sum(1, 2, 3, 4, 5, 6, 7, 9)
	//print(a)
	var b = []int{1, 2, 3, 4, 5, 6}

	print(copy(b[2:], b[3:]))
	fmt.Print(b[:len(b)-1])
	c := map[string]int{}
	sl, ok := c["hehe"]
	if !ok {
		print(ok)
	}
	//print(ok)
	print(sl)
}
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}

	return a
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

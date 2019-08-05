package main

import "fmt"

func main() {
	str := "Hello,世界"
	var r []int32

	//utf-8遍历
	for i := 0; i < len(str); i++ {
		ch := str[i]
		fmt.Println(ch)
	}
	fmt.Println("=============>Unicode遍历")
	//Unicode遍历
	for _, ch1 := range str {
		fmt.Println(ch1)
	}
}

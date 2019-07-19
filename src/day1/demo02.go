package main

import "fmt"

func main() {
	chinese := 90
	//english := 80.5
	avg := string(chinese)
	a := '一'
	var b rune ='哈'
	fmt.Printf("%T ,%v, %U",avg,avg,a)
	fmt.Printf("%T ,%v, %U",b,b,b)
	fmt.Println(string(21704))
	fmt.Printf("%T , %05d ,\n",123,123)
	fmt.Printf("%T , %q ,\n",123,"12.3294")


}

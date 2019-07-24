package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fmt.Println(time.Now())
	rand.Seed(time.Now().UnixNano())
	//10 20
	a := rand.Intn(10) + 10
	print(a)
	for {
		a = rand.Intn(10+1) + 10
		print(a, "\n")
		if a == 20 {
			break
		}
	}
	print("=================")
	var (
		name string
		age  int
	)
	mkrand()
	fmt.Println("请输入年龄")
	n, e := fmt.Scanln(&name, &age)
	print(n, e.Error())
	fmt.Printf("%c", e.Error())

}

func mkrand() {
	s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	print(s1.Intn(10))
}

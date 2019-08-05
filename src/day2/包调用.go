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

	mkrand()

}

func mkrand() {
	s := time.Now()
	s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	print(s1.Intn(10))
	b := time.Since(s)
	fmt.Print(b)
}

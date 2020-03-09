package main

import (
	"fmt"
)

type study struct {
	name string
	age  int
	sex  bool
}

func main() {
	var ming study = study{
		name: "小明",
		age:  20,
		sex:  true,
	}
	zhizhen(ming)
	sums(2, 4, 2, 34, 3)
}

func sums(sum ...int) {
	fmt.Print(sum)
	println(len(sum))
	fmt.Println(sum[0:2])
	for _, v := range sum {
		fmt.Println(v)
	}
}

func digui(a int64) int64 {
	if a == 0 {
		return 1
	}
	a = a * digui(a-1)
	return a

}

func zhizhen(s study) {
	a := &s
	print("年龄:", s.age, " 姓名:", s.name, " 性别:", s.sex, "\n")
	print("年龄:", a.age, " 姓名:", a.name, " 性别:", a.sex, "\n")
	print("年龄:", *&a.name, " 姓名:", a, " 性别:", a.sex, "\n")

}

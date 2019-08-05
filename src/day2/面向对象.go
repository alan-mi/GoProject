package main

import "fmt"

type pp interface {
	call()
}

type Person struct {
	name string
	age  int8
	sex  bool
}
type Student struct {
	Person
	class string
}

func main() {

	stu := pp(Student{Person{name: "haha", age: 20, sex: true}, "五年一般"})
	//stu.name = "haha"
	//stu.age = 20
	//stu.sex = true
	//stu.class = "五年一班"
	if instence, ok := stu.(Student); ok {
		fmt.Print(instence.name)
	}
	stu.call()
	fmt.Print(stu)
}

func (s Student) String() string {
	return fmt.Sprintf("name: %v", s.name)
}

func (s Student) call() {
	fmt.Print(s.name, "打电话了")
}

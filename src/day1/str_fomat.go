package main

import (
	"fmt"
)

func main() {
	var a []int
	fmt.Print(a)
	a = []int{1, 2, 4, 5, 6, 7}
	TestCont(a)
	fmt.Print(a, "\n")
	fmt.Printf("%T,%x \n", a, *&a)
	a = append(a, 33)
	//a = a[3:]
	a = append(a[3:], a[:]...)
	fmt.Println(a)
	b := make([]int, 5, 20)
	b = append(b, 2)
	b = append(b, 2)
	b = append(b, 2)
	b = append(b, 2)
	b = append(b, 2)
	fmt.Println(b)
	mapTest()
}
func mapTest() {
	a := make(map[string]string)
	a["adf"] = "adf"
	delete(a, "adf")
	fmt.Println(a)

}
func TestCont(s []int) {
	s[0] = s[2]

}

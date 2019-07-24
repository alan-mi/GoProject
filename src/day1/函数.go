package main

import "strings"

var a int = 7
var b int = 9

//定义接口type
type caseFunc func(string) string

func main() {
	a, b := 33, 22
	c := 0
	println(a, b, c)
	c = sum(a, b)
	println(c)
	st := "asdfdsfasdfasdf"
	print(bianliString(st))
	print(testhe(st, bianliString))
	print(test4(st, bianliString))

}

func sum(a, b int) int {

	c := a + b
	return c

}

func bianliString(str string) string {
	res := ""
	for i, v := range str {
		if i%2 == 0 {
			res += strings.ToUpper(string(v))
		} else {
			res += string(v)
		}

	}
	return res
}

func testhe(str string, myfunc func(string) string) string {
	return myfunc(str)
}

func test4(str string, myfunc caseFunc) string {
	return myfunc(str)
}

package main

import (
	"errors"
	"fmt"
	"log"
)

var errMizero = errors.New("错误的零输入")

type DivError struct {
	x, y int
}

func (DivError) Error() string {
	return "输入的不对"
}

func main() {
	//neizhi()
	//tool(div)
	test(20, 0)
	print("adsfads" + "asdfasdf")
	var a []interface{}
	a = append(a, "adfadf")
	a = append(a, 123)
	fmt.Println(a[0])

}

func test(x, y int) {
	z := 0
	func() {
		defer func() {
			fmt.Println(recover())
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
		println(x / y)
		println("aaa")
	}()
	println(z)
}

func neizhi() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()
	panic("i am deead")
	println("exit.")
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, DivError{x, y}
	}
	return x / y, nil

}

func tool(some func(int, int) (int, error)) {
	a, ok := some(20, 0)
	fmt.Println(a, ok)
	if ok == errMizero {
		log.Fatalln(ok)
	}
	if ok != nil {
		switch e := ok.(type) {
		case DivError:
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println(e)
		}
		log.Fatalln(ok)

	}

}
func die(x, y int) (int, error) {
	if y == 0 {
		return 0, errMizero
	}
	return x / y, nil
}

package main

import (
	"runtime"
	"sync"
	"time"
)

var c int

func counter() int {
	c++
	return c
}
func main() {
	ex := make(chan struct{})
	//go println("hello world!")
	//go func(s string) {print(s)}("hellll sdjfals")
	a := 100
	go func(x, y int) {
		time.Sleep(time.Second)
		println("go:", x, y)
		close(ex)
	}(a, counter())
	a += 100
	println("main:", a, counter())
	//time.Sleep(time.Second * 3)
	<-ex
	print("exit")
	dijianjishu()
}

func dijianjishu() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)
			println("推出", id)
		}(i)
	}
	wg.Wait()
	runtime.Goexit()

}

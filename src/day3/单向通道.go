package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main() {
	siglechanl()
}
func siglechanl() {
	c := make(chan int, 0)
	c = nil
	c <- 10
	var wg sync.WaitGroup
	wg.Add(2)
	var send chan<- int = c
	var recv <-chan int = c
	go func() {
		defer wg.Done()

		for x := range recv {
			fmt.Printf("%T,%v,%p,%d,%d %p %p\n", c, c, c, len(c), cap(c), unsafe.Pointer(&c), &c)
			println(x)
		}
		a := <-recv
		println(a)
	}()
	go func() {
		defer wg.Done()
		defer close(c)
		for i := 0; i < 5; i++ {

			send <- i
			send <- i
			send <- i

		}
	}()
	wg.Wait()
}

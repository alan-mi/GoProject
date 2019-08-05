package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a, b chan int = make(chan int, 3), make(chan int)
	var c chan bool
	a <- 3
	fmt.Println(a, b, c)
	fmt.Printf("%p,%d\n", a, unsafe.Sizeof(a))
	tongdaozhuce()
}

func shoufa() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for {
			x, ok := <-c
			if !ok {
				return
			}
			fmt.Println(x)
		}
	}()
	c <- 1
	c <- 1
	c <- 1
	c <- 1
	c <- 1
	close(c)
	fmt.Print(done, "---------")
	<-done
}

func xunhuan() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for x := range c {
			println(x)
		}
	}()
	c <- 2
	c <- 2
	c <- 2
	c <- 2
	close(c)
	<-done
}

func tongdaozhuce() {
	c := make(chan int, 5)
	c <- 10
	c <- 20
	c <- 20
	c <- 20
	//close(c)
	for i := 0; i < cap(c)+4; i++ {
		c <- 60
		x, ok := <-c
		println(i, ":", ok, x)

	}

}

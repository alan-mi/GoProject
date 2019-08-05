package main

import (
	"sync"
)

type receiver struct {
	sync.WaitGroup
	data chan int
}

func main() {
	r := newReceiver()
	r.data <- 12
	r.data <- 20
	close(r.data)
	r.Wait()
}
func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		for x := range r.data {
			println(x)
		}
		defer r.Done()
	}()
	return r
}

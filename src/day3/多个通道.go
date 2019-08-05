package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	a, b := make(chan int), make(chan int)
	go func() {
		defer wg.Done()
		for {
			var (
				name string
				x    int
				ok   bool
			)
			select {
			case x, ok = <-a:
				name = "a"
				a = nil
			case x, ok = <-b:
				name = "b"

			}
			if !ok {
				return
			}
			println(name, x)
		}
	}()
	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)
		var c int
		var d int
		for i := 0; i < 10; i++ {
			select {
			case a <- i:
				c++
				println("a==>", c)

			case b <- i * 10:
				d++
				println("b==>", d)

			}
		}
	}()
	wg.Wait()
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := make([]int, 0, 5)
	push := func(x int) error {
		n := len(stack)
		if n == cap(stack) {
			return errors.New("沾满了")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}

	pop := func() (int, error) {
		n := len(stack)
		if n == 0 {
			return -1, errors.New("空了")

		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x, nil
	}
	for i := 0; i < 7; i++ {
		fmt.Printf("push %d: %v,%v\n", i, push(i), stack)
	}
	for i := 0; i < 7; i++ {
		x, err := pop()
		fmt.Printf("pop: %d, %v ,%v\n", x, err, stack)
	}
}

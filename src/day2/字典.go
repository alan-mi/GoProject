package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type te struct {
	name string
	age  string
}

func main() {
	//var m = map[int]string{}
	//m[1] = "a"
	//fmt.Print(m)
	//zidian()
	//test_map()
	a := te{name: "haha", age: "hxix"}
	fmt.Println(&a)

}

func zidian() {
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i + 10
	}
	fmt.Println(m)
	for i := range m {
		if i == 5 {
			m[100] = 1000
		}
		delete(m, i)
		fmt.Println(i, m)
	}
}

func test_map() {
	var lock sync.RWMutex //读写锁
	m := make(map[string]int)
	go func() {
		for {
			lock.Lock() //使用锁
			m["a"] += 1 //写操作
			lock.Unlock()
			time.Sleep(time.Microsecond)
			println(m["a"])
		}
	}()
	go func() {
		for {
			lock.RLock()
			_ = m["b"] //读操作
			lock.RUnlock()
			time.Sleep(time.Microsecond)
		}
	}()
	select { //阻止结束

	}
}

func test_sudo() {
	a := make(map[int]int)
	for i := 0; i < 1000; i++ {
		a[i] = i
	}

}

func test_cap() {
	a := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i
	}
}

func Btest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_cap()
	}
}
func Ctest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_sudo()
	}
}

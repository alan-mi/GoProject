package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	pl_game()
}

func pl_game() {
	target := randnum()
	var shuru int
	count := 0
	fmt.Println("猜数字游戏开始...")
	for {

		fmt.Scan(&shuru)
		count++
		if shuru > target {
			fmt.Println("大了!")
		} else if shuru < target {
			fmt.Println("小了!")
		} else {
			fmt.Print("猜对了！！！")
			fmt.Println("猜了", count, "次！")
			pl_game()
		}
		if count == 6 {
			gameover(target, count)
			pl_game()
		}

	}
}

func gameover(target int, count int) {
	fmt.Println("太笨了都猜了", count, "次了")
	fmt.Println("正确答案是", target)

}

func randnum() int {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r1.Intn(100)
}

package main

import "time"

func main() {
	//test()
	//test2()
	cutBamboo()
}
func test(){
	i := 20
	for ; i >= 0; i-- {
		println(i)
	}
	for ; i<=20	;i++{
		println(i)
	}
	for ;	;i--{
		println(i)
		time.Sleep(1)
		if i ==-10000{
			break
		}
	}
}


func test2(){
	sum :=0
	for i := 1; i<=100;i++{
		if i%3 ==0{
			sum += i
		}

	}
	println(sum)
}

func cutBamboo()  {
	str := "asdfasd"
	for i ,value := range str{
		println(i,value)
	}
}
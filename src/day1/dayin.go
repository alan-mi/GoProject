package main

func main() {
	//juxing()
	zhishu2()
}

//打印循环

func juxing() {
	for i := 0; i <= 10; i++ {
		for j := 10 - i; j > 0; j-- {
			print(" ")
		}
		//for k := 11-i; k > 0;k--{
		//	print("0")
		//}
		for k := i; k > 0; k-- {
			print("0")
		}
		println()
	}

}
func zhishu() {
	for i := 1; i > 0; i++ {
		for k := i - 1; k > 1; k-- {
			if i%k == 0 {
				break
			} else if k == 2 {
				println(i)
			}
		}
	}

}
func zhishu2() {
	i := 1;
LOOP:
	for i > 0 {
		i++
		for k := i - 1; k > 1; k-- {
			if i%k == 0 {
				goto LOOP
			}

		}
		println(i)
	}

}

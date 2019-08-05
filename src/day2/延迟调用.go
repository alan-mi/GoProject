package main

func main() {
	x, y := 1, 2
	defer func(a int) {
		println("defer x,y= ", y, a)
	}(x)
	x += 100
	y += 100
	println(x, y)
}

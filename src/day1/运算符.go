package main

import "fmt"

var (
	a = 21.0
	b = 5.0
	c float64
	d = true
	e = false
)
func main() {
	arithmetic()
	fmt.Println(a)
	relational()
}

func arithmetic() {
	c = a + b
	a = 33
	fmt.Printf("%.3F ,\n",c)
	fmt.Printf("%v \n",int(c)%int(b))
	fmt.Println(a)
	b = 33
}

func relational()  {
	if (a == b){
		fmt.Print("duiadsf")
	}
}

func logicl()  {
	d = true
	e = false
	if (d && e){
		fmt.Print("aaaaaaaaa")
	} else {
	fmt.Print("badsfasdf")
	}

}

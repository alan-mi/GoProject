package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func changeBook(book *Books) { // 这个方法传入的参数一个Books类型的指针
	book.title = "book1_change" //直接用指针.属性的方式，修改原地址的值。
}

func main() {
	var book1 Books
	book1.title = "book1"
	book1.author = "zuozhe"
	book1.book_id = 1
	changeBook(&book1) //将book1这个对象的内存地址传进去，
	fmt.Println(book1)
}

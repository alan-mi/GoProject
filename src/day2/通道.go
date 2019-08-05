package main

func main() {
	done := make(chan struct{})
	c := make(chan string)
	go func() {
		s := <-c
		println(s)
		close(done)
	}()
	c <- "hai!"
	<-done //一直阻塞直到管道关闭
	tongdao()
}

func tongdao() {
	c := make(chan int, 1)
	c <- 1

	c <- 2
	close(c)
	println(<-c)
	println(<-c)

}

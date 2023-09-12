package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {

	ch := make(chan int, 1) //创建一个容量为1的有缓冲区通道
	ch <- 10
	x := <-ch // 从ch中接收值并赋值给变量x
	fmt.Print(x)
	close(ch) //关闭通道

	ch2 := make(chan int) // 使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值
	go recv(ch2)          // 启用goroutine从通道接收值
	ch <- 10              // 把10发送到通道
	fmt.Println("发送成功")
}

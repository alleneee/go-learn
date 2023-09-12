package main

import (
	"fmt"
	"time"
)

func hello1(i int) {
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go hello1(i) //每个循环进来起一个线程去执行
	}
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second * 2)
}

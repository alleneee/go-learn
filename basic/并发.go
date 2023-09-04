package main

import (
	"fmt"
	"time"
)

func hellos() {
	fmt.Println("Hello Goroutine!")
}

func hellon(i int) {
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	//go hellos()
	//fmt.Println("main goroutine done!") //只打印了main goroutine done!   goroutine协程
	////在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
	////当main()函数返回的时候该goroutine就结束了
	//time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		go hellon(i)
	}
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second * 2)

}

package main

import (
	"fmt"
)

// 一个函数，用于模拟可能引发 panic 的情况
func divide(a, b int) int {
	if b == 0 {
		panic("除数不能为零")
	}
	return a / b
}

func main() {
	// 使用 defer 和 recover 来捕获和处理 panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生了 panic：", err)
		}
	}()

	fmt.Println("开始运行程序")

	// 调用可能引发 panic 的函数
	result := divide(10, 0)
	fmt.Println("结果：", result) // 这行不会执行，因为前一行会引发 panic

	fmt.Println("程序继续执行") // 这行也不会执行，因为已经发生了 panic
}

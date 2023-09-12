package main

import "fmt"

// 一个接受两个整数和一个函数作为参数的函数
func calculate(a, b int, operation func(int, int) int) int {
	result := operation(a, b)
	return result
}

func main() {
	// 定义一个匿名函数，实现加法操作
	addition := func(x, y int) int {
		return x + y
	}

	// 调用 calculate 函数，并将 addition 函数作为参数传递
	sum := calculate(10, 20, addition)
	fmt.Println("加法结果:", sum)

	// 定义一个不同的匿名函数，实现乘法操作
	multiplication := func(x, y int) int {
		return x * y
	}

	// 调用 calculate 函数，并将 multiplication 函数作为参数传递
	product := calculate(5, 4, multiplication)
	fmt.Println("乘法结果:", product)
}

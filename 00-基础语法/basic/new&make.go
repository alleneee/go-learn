package main

import "fmt"

/*
*
new函数：用于创建值类型的变量，并返回其指针
*/
func main() {
	//make函数：用于创建切片、映射和通道
	slice := make([]int, 5, 10) // 创建一个长度为5，容量为10的整数切片
	m := make(map[string]int)   // 创建一个字符串键和整数值的映射
	ch := make(chan int)        // 创建一个整数类型的通道
	ptr := new(int)             // 创建一个整数类型的变量，并返回其指针
	fmt.Print(slice)
	fmt.Print(m)
	fmt.Print(ch)
	fmt.Print(ptr)

}

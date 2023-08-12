package main

import "fmt"

// 遍历切片的每个元素, 通过给定函数进行元素访问

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func FGen(x, y int) (func() int, func(int) int) {

	//求和的匿名函数
	sum := func() int {
		return x + y
	}

	// (x+y) *z 的匿名函数
	avg := func(z int) int {
		return (x + y) * z
	}
	return sum, avg
}

func main() {

	func(data int) {
		fmt.Println("hello", data)
	}(100) //(100)，表示对匿名函数进行调用，传递参数为 100。
	// 遍历切片的每个元素, 通过给定函数进行元素访问

	// 使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
	//传入x=1,y=2 返回时候调用对应的匿名函数 x+y  和 (x+y) *z 的匿名函数
	f1, f2 := FGen(1, 2)
	fmt.Println("第一个匿名函数:", f1())
	fmt.Println("第二个匿名函数:", f2(3))
}

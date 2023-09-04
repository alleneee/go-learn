package main

import "fmt"

/**
	func(参数列表)(返回参数列表){
    	函数体
	}
*/
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

func test1(x, y string) (z string) {
	return x + y
}
func test2(fn func() int) int {
	return fn()
}
func fn2() int {
	return 200
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
	//传入x=1,y=2
	f1, f2 := FGen(1, 2)
	fmt.Println("第一个匿名函数:", f1())  // x+y的匿名函数
	fmt.Println("第二个匿名函数:", f2(3)) //  (x+y) *z  的匿名函数

	fmt.Println(test1("h", "x"))

	//匿名函数
	// 定义一个匿名函数，并将其赋值给变量 add
	add := func(a, b int) int {
		return a + b
	}

	// 调用匿名函数
	result := add(10, 20)
	fmt.Println("结果:", result)

}

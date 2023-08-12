package main

import "fmt"

var grade string = "B"
var score int = 90

func main() {
	x := 0
	// if 语句
	if x <= 0 {
		fmt.Println("为真进入这里")
	} else if x > 1 {
		fmt.Println("为真进入这里2")
	} else {
		fmt.Println("为真进入这里3")
	}
	// for
	sum := 0
	//i := 0; 赋初值，i<10 循环条件 如果为真就继续执行 ；i++ 后置执行 执行后继续循环
	for i := 0; i < 10; i++ {
		sum += i
		//return //执行一次就结束了
		//break //跳出循环,还会继续执行循环外的语句
		//panic("出错了")		//报错了，直接结束
		goto breakHere // 跳转到标签
	breakHere:
		fmt.Println("进入标签了!!", i)
	}
	fmt.Println(sum)

	//for range 结构是Go语言特有的一种的迭代结构
	//for range 可以遍历数组、切片、字符串、map 及管道（channel）

	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println("遍历map", key, value)
	}
	//Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case
	fmt.Println("before", score, grade)
	switch score {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println("after", score, grade)

}

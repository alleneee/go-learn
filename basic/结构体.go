package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	//使用.来访问结构体的成员变量,结构体成员变量的赋值方法与普通变量一致。
	var p Point //实例化
	p.X = 1
	p.Y = 2
	fmt.Printf("%v,x=%d,y=%d", p, p.X, p.Y)
}

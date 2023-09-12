package main

import "fmt"

/**
在Go中，接口是隐式实现的，这意味着你无需明确声明一个类型实现了某个接口。只要一个类型拥有接口中定义的所有方法，它就被视为实现了该接口，不需要显式地声明它实现了哪个接口
*/
// 定义一个接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义一个矩形结构体，实现了 Shape 接口
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 结构体实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 结构体实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func main() {
	// 创建一个 Rectangle 类型的实例
	rectangle := Rectangle{Width: 5, Height: 3}

	// 使用 Shape 接口来调用 Rectangle 的方法
	var shape Shape
	shape = rectangle

	fmt.Println("面积:", shape.Area())
	fmt.Println("周长:", shape.Perimeter())
}

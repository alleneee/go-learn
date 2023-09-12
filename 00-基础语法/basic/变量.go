package main

import (
	"fmt"
)

// 全局变量
// 声明类型  var 名字 类型
var age int

// 自动推导
var age1 = 12

// boolean类型
var boy bool

// 常量 const
const pi = 3.14159

// 批量申明
const (
	e   = 2.7182818
	pii = 3.1415926
)

var hello string = "hello"

func main() {
	// 局部变量
	age = 25

	//类型转换
	//只有相同底层类型的变量之间可以进行相互转换（如将 int16 类型转换成 int32 类型）
	//不同底层类型的变量相互转换时会引发编译错误（如将 bool 类型转换为 int 类型）
	a := 5.0
	b := int(a)
	fmt.Println("类型转换:", b)

	//修改字符串
	s1 := "localhost:8080"
	// 强制类型转换 string to byte
	strByte := []byte(s1)
	fmt.Println("修改前:", strByte)

	// 下标修改
	strByte[len(s1)-1] = '1'
	fmt.Println("修改后:", strByte)

	// 强制类型转换 []byte to string
	s2 := string(strByte)
	fmt.Println("强制类型转换 []byte to string:", s2)

	fmt.Println("常量:", pi, pii, e)

}

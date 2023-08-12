package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// 声明类型
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

// 指针 &变量：取出变量的地址 *指针= 取出指针对应变量的值
// 当一个指针被定义后没有分配到任何变量时，它的默认值为 nil
var ptr = &age1 //取出age1的内存地址

var intPtr *int       // int类型的指针
var strPtr *string    // string类型的指针
var floatPtr *float64 // float64类型的指针
var boolPtr *bool     // bool类型的指针

/*
*
字符串中可以使用转义字符来实现换行、缩进等效果，常用的转义字符包括:
 1. `\n：`换行符
 2. `\r：`回车符
 3. `\t：`tab 键
 4. `\u 或 \U：`Unicode 字符
 5. \：反斜杠自身
*/
var hello string = "hello"

func main() {
	age = 25
	// 定义 float32 float64类型
	floatStr1 := 3.2
	floatStr2 := 3.22

	/**
	- `print ：` 结果写到标准输出
	- `Sprint：`结果会以字符串形式返回
		%c  单一字符
		%T  动态类型
		%v  本来值的输出
		%+v 字段名+值打印
		%d  十进制打印数字
		%p  指针，十六进制
		%f  浮点数
		%b 二进制
		%s string
	*/
	fmt.Print(age)
	fmt.Println() // 换行

	fmt.Printf(strconv.Itoa(age1))
	fmt.Println() // 换行

	fmt.Printf("调用sum函数求和", strconv.Itoa(sum(1, 2)))
	fmt.Println() // 换行

	fmt.Printf("打印floast32:%.2f\n", floatStr1)
	fmt.Println() // 换行

	fmt.Printf("打印floast64:%.2f\n", floatStr2)
	fmt.Println() // 换行

	fmt.Println(boy)
	fmt.Println() // 换行

	str1 := "你好,"
	str2 := "ms的go教程"
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)
	// Sprint 以字符串形式返回
	result := fmt.Sprintf(stringBuilder.String())
	fmt.Println(result)

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

	fmt.Println("age1指针地址:", ptr)
	fmt.Println("age1指针地址指向变量的值:", *ptr)

	*ptr = 1111
	fmt.Println("使用指针修改age1值:", *ptr)

}

// 第一个 int 标识 形参的类型
// 第二个 int 标识 返回值类型
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}

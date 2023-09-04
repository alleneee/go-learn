package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/*
*
字符串中可以使用转义字符来实现换行、缩进等效果，常用的转义字符包括:
 1. `\n：`换行符
 2. `\r：`回车符
 3. `\t：`tab 键
 4. `\u 或 \U：`Unicode 字符
 5. \：反斜杠自身
*/

func main() {
	str1 := "你好,"
	str2 := "ms的go教程"
	var stringBuilder bytes.Buffer
	//除了使用+进行拼接，我们也可以使用WriteString()
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)
	// Sprint 以字符串形式返回
	result := fmt.Sprintf(stringBuilder.String())
	fmt.Println(result)

	// 定义 float32 float64类型
	//短变量声明
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

	age := 11
	fmt.Print(age)
	fmt.Println() // 换行
	age1 := 22
	fmt.Printf(strconv.Itoa(age1))
	fmt.Println() // 换行

	fmt.Printf("调用sum函数求和", strconv.Itoa(sum(1, 2)))
	fmt.Println() // 换行

	fmt.Printf("打印floast32:%.2f\n", floatStr1)
	fmt.Println() // 换行

	fmt.Printf("打印floast64:%.2f\n", floatStr2)
	fmt.Println() // 换行

	boy := "boy"
	fmt.Println(boy)
	fmt.Println() // 换行

}

// 第一个 int 标识 形参的类型
// 第二个 int 标识 返回值类型
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}

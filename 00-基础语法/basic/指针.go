package main

import "fmt"

func main() {

	age1 := 12
	/*
		指针 &变量：取出变量的地址 *指针= 取出指针对应变量的值
		当一个指针被定义后没有分配到任何变量时，它的默认值为 nil
	*/
	var ptr = &age1 //取出age1的内存地址

	/*
		var intPtr *int       // int类型的指针
		var strPtr *string    // string类型的指针
		var floatPtr *float64 // float64类型的指针
		var boolPtr *bool     // bool类型的指针

	*/

	fmt.Println("age1指针地址:", ptr)
	fmt.Println("age1指针地址指向变量的值:", *ptr)

	*ptr = 1111
	fmt.Println("使用指针修改age1值:", *ptr)
}

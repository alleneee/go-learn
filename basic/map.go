package main

import "fmt"

func main() {
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println("遍历map", key, value)
	}
	//Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case
	fmt.Println("before", score, grade)
	
}

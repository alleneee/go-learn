package main

import (
	"fmt"
	"strconv"
	"sync"
)

// var 数组变量名 [元素数量]Type
// 数组容量必须提前确定
var arr = [3]int{1, 3, 2}

// 多位数组
// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
var array [4][2]int

//// 使用数组字面量来声明并初始化一个二维整型数组
//array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
//// 声明并初始化数组中索引为 1 和 3 的元素
//array = [4][2]int{1: {20, 21}, 3: {40, 41}}
//// 声明并初始化数组中指定的元素
//array = [4][2]int{1: {0: 20}, 3: {1: 41}}

func main() {
	//遍历
	for index, value := range arr {
		fmt.Printf("索引:%d,值：%d \n", index, value)
	}

	for v := range arr {
		fmt.Printf(strconv.Itoa(v))
	}

	//并发map
	//sync.Map 不能使用 make 创建
	var scene sync.Map
	// 将键值对保存到sync.Map
	//sync.Map 将键和值以 interface{} 类型进行保存。
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	//遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("遍历所有sync.Map中的键值对:", k, v)
		return true
	})
}

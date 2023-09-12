package main

import (
	"fmt"
	"log"
	_ "log"
	"time"
	_ "time"
)

func main() {
	//functionName(2, 1)

	//这是直接使用匿名函数
	s1 := test(func() int { return 100 })
	//这是传入一个函数
	s1 = test(fn)
	//传递函数
	fmt.Println(s1)

	// 创建一个玩家生成器
	generator := playerGen("ms")
	// 返回玩家的名字和血量
	name, hp := generator()
	// 打印值
	fmt.Println(name, hp)

	generator1 := playerGen1()
	name1, hp1 := generator1("ms")
	// 打印值
	fmt.Println(name1, hp1)

	start := time.Now()
	log.Printf("开始时间为：%v", start)
	//调用 defer 关键字会立刻拷贝函数中引用的外部参数
	defer log.Printf("时间差：%v", time.Since(start)) // Now()此时已经copy进去了
	//不受这3秒睡眠的影响
	time.Sleep(3 * time.Second)

	log.Printf("函数结束")
}

func functionName(param1, param2 int) int {
	if param1 > param2 {
		print(param1)
		return param1
	}
	print(param2)
	return param2
}

func test(fn func() int) int {
	return fn()
}
func fn() int {
	return 200
}

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen1() func(string) (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func(name string) (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

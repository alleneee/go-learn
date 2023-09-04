package main

import (
	"fmt"
	"sync"
)

var x int64

/*
*
sync.WaitGroup 是Go语言标准库中的一个同步原语，它用于等待一组 goroutine 完成它们的任务。
WaitGroup 通常用于确保在主 goroutine 中等待所有其他 goroutines 完成工作，然后再继续执行主 goroutine 的后续操作。

sync.WaitGroup 提供了三个主要方法：

Add(delta int)：用于增加等待的任务计数器。你可以在主 goroutine 中调用 Add 来设置等待的任务数量。

Done()：用于标记一个任务已完成。在每个 goroutine 中，在完成任务后，应该调用 Done 来减少任务计数器。

Wait()：用于阻塞主 goroutine，直到任务计数器减少到零。一旦任务计数器为零，Wait 方法将解除阻塞，允许主 goroutine 继续执行。
*/
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}
func main() {
	wg.Add(2) //设置等待的任务数量
	go add()
	go add()
	wg.Wait() //用于阻塞主 goroutine，直到任务计数器减少到零
	fmt.Println(x)
}

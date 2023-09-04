package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x1  int64
	wg1 sync.WaitGroup
	// 互斥锁
	lock1 sync.Mutex
	// 读写锁
	rwlock1 sync.RWMutex
)

func write2() {
	// lock.Lock()   // 加互斥锁
	rwlock1.Lock() // 加写锁
	x1 = x1 + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock1.Unlock()                  // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg1.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock1.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock1.RUnlock()            // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write2()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

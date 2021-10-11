package main

import (
	"fmt"
	"sync"
)

var (
	sum   int
	mutex sync.RWMutex
)

func main() {
	run()
}

func run() {
	var wg sync.WaitGroup
	// 因为要监控110个协程，所以设置计数器为110
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			// 计数器值减1
			defer wg.Done()
			go add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			// 计数器值减1
			defer wg.Done()
			fmt.Println("和为 ", readSum())
		}()
	}
	// 一直等待，直到计时器值为0
	wg.Wait()
}

func add(i int) {
	mutex.Lock()
	sum += i
	mutex.Unlock()
}
func readSum() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

package main

import (
	"fmt"
	"time"
)

var sum = 0

func main() {
	// 共享的资源
	for i := 0; i < 100; i++ {
		go add(10)
	}
	// 防止提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("和为: ", sum)
}

func add(i int) {
	sum += i
}

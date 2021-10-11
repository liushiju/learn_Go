package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)
	// 同时开启3个goroutine下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()
	go func() {
		secondCh <- downloadFile("secondCh")
	}()
	go func() {
		threeCh <- downloadFile("threeCh")
	}()
	// 开始select多路复用，哪个channel能获取到值
	// 就说明哪个先下载好，就用哪个
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filepath := <-secondCh:
		fmt.Println(filepath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
	}
}

func downloadFile(chanName string) string {
	// 模拟下文文件，可以自己谁家time.Sleep时间点
	time.Sleep(time.Second)
	return chanName + " : filePath"
}

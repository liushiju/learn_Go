package main

import (
	"fmt"
	"time")


func main() {
	result := make(chan string)
	go func() {
		// 模拟网络访问
		time.Sleep(8 * time.Second)
		result <- "服务端结果"
	}()
	select {
	case v := <-result:
		fmt.Println(v)
	case <- time.After(5*time.Second):
		fmt.Println("网络访问超时了")
	}
}

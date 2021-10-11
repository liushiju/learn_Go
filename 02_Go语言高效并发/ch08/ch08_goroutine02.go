package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		fmt.Println("大锤")
		ch <- "goroutine 完成"
	}()
	fmt.Println("我是main goroutine")
	v := <-ch
	fmt.Println("接收到的chan中的值为: ", v)
}

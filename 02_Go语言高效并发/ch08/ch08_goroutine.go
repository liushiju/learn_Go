package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("大锤")
	fmt.Println("我是main goroutine")
	time.Sleep(time.Second)
}

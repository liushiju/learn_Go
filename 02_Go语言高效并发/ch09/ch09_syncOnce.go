package main

import (
	"fmt"
	"sync"
)

func main()  {
	doOnce()
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	// 用于等待协程执行完毕
	done := make(chan bool)
	// 启动10个协程once.Do(onceBody)
	for i:=0;i<10;i++{
		go func()  {
			// 把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done<- true
		}()
	}
	for i:=0;i<10;i++{
		<-done
	}
}

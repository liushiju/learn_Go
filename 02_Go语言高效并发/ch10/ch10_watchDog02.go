package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// 用来停止监控狗
	stopCh := make(chan bool)
	go func() {
		defer wg.Done()
		watchDog(stopCh, "【监控狗1】")
	}()
	// 想让监控狗监控5秒
	time.Sleep(5 * time.Second)
	// 发出停止指令
	stopCh <- true
	wg.Wait()
}

func watchDog(stopCh chan bool,name string) {
	// 开启for select循环，一直后台监控
	for {
		select {
		case <- stopCh:
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控... ...")
		}
		time.Sleep(1 * time.Second)
	}
}

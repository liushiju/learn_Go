package main

import "fmt"

func moreDefer() {
	defer fmt.Println("first defer")
	defer  fmt.Println("Second defer")
   	defer  fmt.Println("Three defer")
   	fmt.Println("函数自身代码")
}

func main()  {
	moreDefer()
}
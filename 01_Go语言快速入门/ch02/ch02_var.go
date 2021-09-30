package main

import "fmt"

func main() {
	var i int = 10
	/*
	Go具有类型推到功能
	var i = 10
	一次声明多个变量
	var (
		j int = 0
		k int = 1
	)
	也可以这么写
	var (
		j = 0
		k = 1
	)
	*/
	fmt.Println(i)
}
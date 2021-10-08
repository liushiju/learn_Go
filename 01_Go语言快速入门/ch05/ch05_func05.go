package main

import "fmt"

func main() {
	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
}

func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
// Go语言中，函数也是一种类型，它可以被用来声明
// 函数类型的变量、参数或者作为另一个函数的返回值类型
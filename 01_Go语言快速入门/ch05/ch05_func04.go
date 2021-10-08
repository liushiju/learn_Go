package main

import "fmt"


// 匿名函数
func main() {
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(1,2))
}

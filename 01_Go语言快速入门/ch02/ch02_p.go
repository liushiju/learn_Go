package main

import "fmt"

func main() {
	i := 10
	pi := &i
	fmt.Println(*pi)
}

// 指针的值就是变量的内存地址

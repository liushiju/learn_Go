package main

import (
	"fmt"	
)


func main() {
	// 比较繁琐
	// const (
	// 	one   = 1
	// 	two   = 2
	// 	three = 3
	// 	four  = 4
	// )
	const(
		one=iota + 1
		two
		three
		four
	)
	fmt.Println(one,two,three,four)
}

/*
iota 的初始值为0，它的作用就是每一个有变量声明的后面+1
*/
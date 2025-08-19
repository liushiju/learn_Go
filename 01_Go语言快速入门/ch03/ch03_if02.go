package main

import "fmt"

func main() {
	if i := 6; i > 10 {
		fmt.Println("i>10")
	} else if i > 5 && i <= 10 {
		fmt.Println("i<5<100")
	} else {
		fmt.Println("i<5")
	}
}

// 在if关键字之后，i>10条件之前，通过;号分隔被初始化的i:=6。这个简单语句主要用来在if条件判断之前做一些初始化工作

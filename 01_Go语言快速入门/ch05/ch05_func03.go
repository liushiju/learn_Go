package main

import "fmt"

// 定义可变参数，只要在参数类型前加三个点...
func sum(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}
func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
}

// 可变参数的类型就是切片，案例中params参数的类型就是[]int
// 所以可以使用for range进行循环
// 定义的函数中，有普通参数，又有可变参数，可变参数一定
// 要放在参数列表最后一个

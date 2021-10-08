package main

import "fmt"

/*
方法和函数是两个概念，但是很相似
不同点：方法必须要有一个接受者，接受者是一个类型
方法和类型绑定在一起，成为这个类型的方法
*/

// 定义新类型Age，相当于uint的重命名
type Age uint

// 方法String()就是类型Age的方法，类型Age是方法String()的接收者
func (age Age) Str() {
	fmt.Println("this age is ", age)
}

func main() {
	age := Age(25)
	age.Str()
}

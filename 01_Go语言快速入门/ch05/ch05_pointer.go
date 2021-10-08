package main

import "fmt"

// 方法的接收者可以是值类型也可以是指针类型

type Age uint

func (age Age) Str() {
	fmt.Println("this age is ", age)
}

func (age *Age) Modify() {
	*age = Age(30)
}

func main() {
	age := Age(25)
	age.Str()
	age.Modify()
	// 或者
	(&age).Modify()
	age.Str()
}

// 方法的调用者，既可以是值，也可以是指针

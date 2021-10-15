package main

import "fmt"

func main() {
	p:=person{name: "张三",age: 30}
	fmt.Printf("main函数：p的内存地址为%p\n",&p)
	modifyPerson(p)
	fmt.Println("person name:",p.name,",age:",p.age)
}

func modifyPerson(p person) {
	fmt.Printf("modifyPerson函数：p的内存地址为%p\n",&p)
	p.name = "李四"
	p.age = 20
}

type person struct {
	name string
	age  uint
}
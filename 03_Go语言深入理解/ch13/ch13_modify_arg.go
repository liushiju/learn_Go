package main

import "fmt"

func main() {
	p := person{name: "张三", age: 30}
	modifyPerson(p)
	fmt.Println("person name:", p.name, ",age:", p.age)
}

func modifyPerson(p person) {
	p.name = "李四"
	p.age = 20
}

type person struct {
	name string
	age  uint
}

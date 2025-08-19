package main

import "fmt"

type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s*%s", addr.province, addr.city)
}

func NewPerson(name string,age uint) *person {
	return &person{name: name,age:age}
}


func main()  {
	p1 := NewPerson("张三",25)
	fmt.Println(p1)
}
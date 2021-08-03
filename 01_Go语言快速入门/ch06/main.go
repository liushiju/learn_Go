package main

import "fmt"

type person struct {
	name string
	age  uint
	address
}

type address struct {
	province string
	city     string
}

func NewPerson(name string) *person {
	return &person{name: name}
}

func (p *person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func main() {
	p1 := NewPerson("张三")
	var s fmt.Stringer
	s = p1
	p2 := s.(*person)
	fmt.Println(p2)
	//a := s.(address)
	//fmt.Println(a)
	a, ok := s.(address)
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("s不是一个address")
	}

}

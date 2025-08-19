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

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	p := person{
		name: "hanson",
		age:  26,
		addr: address{
			province: "广东",
			city:     "深圳",
		},
	}
	fmt.Println(p.addr.province)
	printString(p)
	printString(p.addr)
}

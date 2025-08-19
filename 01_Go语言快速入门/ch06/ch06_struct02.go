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

func main() {
	p := person{
		name: "Hanson",
		age:  26,
		addr: address{
			province: "广东",
			city:     "深圳",
		},
	}
	fmt.Println(p.addr)
}

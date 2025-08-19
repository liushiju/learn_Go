package main

import "fmt"

type person struct {
	name string
	age  uint
}

func main() {
	p := person{"hanson", 25}
	fmt.Println(p.name, p.age)
}
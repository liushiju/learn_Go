package main

import "fmt"

type person struct {
	name string
	age  uint
}

type WalkRun interface {
	Walk()
	Run()
}

func (p person) Walk() {
	fmt.Printf("%s 能走\n", p.name)
}
func (p person) Run() {
	fmt.Printf("%s 能跑, %d 岁\n", p.name, p.age)
}

func da(arg WalkRun)  {
	arg.Walk()
	arg.Run()
}

func main() {
	p := person{"lsj", 26}
	da(p)
}

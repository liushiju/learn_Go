package main

import "fmt"

type sayer interface{
	say()
}

type dog struct{}

type cat struct{}

func (d dog) say() {
	fmt.Println("汪汪汪~")
}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

func da(arg sayer)  {
	arg.say()
}


func main()  {
	var d dog
	d.say()
	c := cat{}
	c.say()
	c2 := cat{}
	da(c2)
}

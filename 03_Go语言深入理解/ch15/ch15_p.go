package main

import (
	"fmt"
	"reflect")


func main() {
	p := person{Name: "寒梅清风", Age: 20}
	ppv := reflect.ValueOf(&p)
	ppv.Elem().Field(0).SetString("张三")
	fmt.Println(p)
}

type person struct {
	Name string
	Age  uint
}